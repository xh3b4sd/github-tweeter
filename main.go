package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/google/go-github/github"
	"github.com/the-anna-project/random"
	"golang.org/x/oauth2"
)

const (
	org  = "xh3b4sd"
	repo = "content"
	dir  = "philosophy"
)

var (
	githubToken = os.Getenv("GITHUB_TWEETER_GITHUB_TOKEN")

	twitterConsumerKey    = os.Getenv("TWITTER_CONSUMER_KEY")
	twitterConsumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	twitterAccessToken    = os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessSecret   = os.Getenv("TWITTER_ACCESS_SECRET")
)

var (
	// seqExp is the regular expression for the conventional sequence number
	// encoded in the content file names. File names must follow a pattern like
	// shown below.
	//
	//     philosophy/0000001
	//     philosophy/0000002
	//     ...
	//     philosophy/0000054
	//     philosophy/0000055
	//
	seqExp = regexp.MustCompile(`([0-9]+)$`)
	// wspExp is the regular expression used to trim space and line control
	// characters from the gathered content. Content might be wrapped or formatted
	// in a certain way which makes it necessary to trim spaces in order to cary
	// out sanitized content for a Tweet on Twitter.
	wspExp = regexp.MustCompile(`[\n\r\t]`)
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		panic(microerror.JSON(err))
	}
}

func mainE(ctx context.Context) error {
	var err error

	var newLogger micrologger.Logger
	{
		c := micrologger.Config{}

		newLogger, err = micrologger.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var ghClient *github.Client
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "initializing github client")

		c := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		))

		ghClient = github.NewClient(c)

		newLogger.LogCtx(ctx, "level", "debug", "message", "initialized github client")
	}

	// Create a Twitter API client for further use below. The required credentials
	// are generated via a Twitter app that has to be set up properly. Therefore
	// you need to register an application account and create an app in the apps
	// dashboard.
	//
	//     https://developer.twitter.com/en/apps
	//
	var twClient *twitter.Client
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "initializing twitter client")

		config := oauth1.NewConfig(twitterConsumerKey, twitterConsumerSecret)
		token := oauth1.NewToken(twitterAccessToken, twitterAccessSecret)

		twClient = twitter.NewClient(config.Client(oauth1.NoContext, token))

		newLogger.LogCtx(ctx, "level", "debug", "message", "initialized twitter client")
	}

	var newRandom random.Service
	{
		c := random.ServiceConfig{
			BackoffFactory: func() random.Backoff {
				return backoff.NewMaxRetries(3, time.Second)
			},
			RandFactory: rand.Int,

			RandReader: rand.Reader,
			Timeout:    1 * time.Second,
		}

		newRandom, err = random.NewService(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	// Fetch the commit of the latest content made in the configured folder. It
	// might happen that commits have been made to fix typos which means that
	// these commits are made on older content files. We want the latest content
	// and therefore check the last 10 commits which should give us enough
	// buffer to fix and improve existing content before we get the chance to
	// add actually new content which is numbered accordingly. The sha we are
	// looking for is then used to get its associated file which changed with
	// the commit. The name of this file then by convention indicates the
	// highest number of files in the sequence of existing file names. The
	// example file names below describe the convention of numbers defining the
	// file sequence. The method below fetches the sha that assumedly added
	// file055 which tells us that there are 55 files to chose from randomly.
	//
	//     path/file001
	//     path/file002
	//     ...
	//     path/file054
	//     path/file055
	//
	var sha string
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "finding latest commit")

		in := &github.CommitsListOptions{
			Path: dir,
			ListOptions: github.ListOptions{
				PerPage: 10,
			},
		}

		out, _, err := ghClient.Repositories.ListCommits(ctx, org, repo, in)
		if err != nil {
			return microerror.Mask(err)
		}

		var n int
		for _, o := range out {
			a := strings.Split(o.GetCommit().GetMessage(), "/")
			if len(a) != 2 {
				continue
			}

			b := strings.Split(a[1], " ")
			if len(b) != 2 {
				continue
			}

			c, err := strconv.Atoi(b[0])
			if err != nil {
				return microerror.Mask(err)
			}

			if c > n {
				n = c
				sha = o.GetSHA()
			}
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("found latest commit %#q", sha))
	}

	// We fetch the file name using the commit hash found above. The commit is
	// expected to have changed exactly one file.
	//
	//     path/file055
	//
	var file string
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "finding latest file")

		out, _, err := ghClient.Repositories.GetCommit(ctx, org, repo, sha)
		if err != nil {
			return microerror.Mask(err)
		}

		file = out.Files[0].GetFilename()

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("found latest file %#q", file))
	}

	// We lookup the number of files in the traversed folder. The name of the
	// changed file is expected to comply with a format as such that the end of
	// the file name is a number of a sequence of files. When we extract the
	// number of the example file name above it should result in the following.
	//
	//     55
	//
	var total int
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "computing total number of files")

		m := seqExp.FindString(file)

		total, err = strconv.Atoi(m)
		if err != nil {
			return microerror.Mask(err)
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("computed total number of files %d", total))
	}

	var number int
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "choosing random number")

		number, err = newRandom.CreateMax(total + 1)
		if err != nil {
			return microerror.Mask(err)
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("chose random number %d", number))
	}

	// The content repo structure is as follows.
	//
	//     philosophy/2020/0000568
	//
	// We need to identify the years, that is sub folders in case we chose a
	// random number that refers to content written in another year.
	//
	var years []string
	{
		in := &github.RepositoryContentGetOptions{}

		_, out, _, err := ghClient.Repositories.GetContents(ctx, org, repo, dir, in)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, d := range out {
			years = append(years, d.GetName())
		}
	}

	// We compute a random file of which we take the content to tweet from. Since
	// we have the number of the upper end of the file sequence we generate a
	// random number and put together the new file name including eventual
	// padding. Padding is necessary because the upper number might have 4 digits
	// while the chosen number might have 2, which implies to add a padding of 2
	// padding characters assuming a consistent conventional file name format.
	var paths []string
	{
		n := strconv.Itoa(number)
		p := len(file) - len(seqExp.ReplaceAllString(file, "")) - len(n)

		for _, y := range years {
			path := dir + "/" + y + "/"
			for i := 0; i < p; i++ {
				path += "0"
			}
			path += n

			paths = append(paths, path)
		}
	}

	var content string
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "finding content")

		for _, p := range paths {
			in := &github.RepositoryContentGetOptions{}

			out, _, r, err := ghClient.Repositories.GetContents(ctx, org, repo, p, in)
			if r.StatusCode == http.StatusNotFound {
				continue
			} else if err != nil {
				return microerror.Mask(err)
			}

			c, err := out.GetContent()
			if err != nil {
				return microerror.Mask(err)
			}

			content = strings.TrimSpace(wspExp.ReplaceAllString(c, " "))
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("found content %#q", content))
	}

	// We just make sure that we deal with valid credentials and collect
	// information about the user on which behalf we are going to tweet below.
	//
	//     https://developer.twitter.com/en/docs/accounts-and-users/manage-account-settings/api-reference/get-account-verify_credentials
	//
	var userName string
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "verifying twitter credentials for user")

		p := &twitter.AccountVerifyParams{
			SkipStatus: twitter.Bool(true),
		}
		user, _, err := twClient.Accounts.VerifyCredentials(p)
		if err != nil {
			return microerror.Mask(err)
		}

		userName = user.ScreenName

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("verified twitter credentials for user %#q", userName))
	}

	// Once the necessary content is gathered it can be tweeted using the Twitter
	// client initialized above. For more information about the API client and API
	// specs see the following resources.
	//
	//     https://godoc.org/github.com/dghubble/go-twitter/twitter#StatusService.Update
	//     https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/post-statuses-update
	//
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "tweeting content")

		_, _, err := twClient.Statuses.Update(content, nil)
		if err != nil {
			return microerror.Mask(err)
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", "tweeted content")
	}

	return nil
}
