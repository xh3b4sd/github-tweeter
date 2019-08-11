package main

import (
	"context"
	"crypto/rand"
	"fmt"
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

func main() {
	err := mainE(context.Background())
	if err != nil {
		panic(microerror.Stack(err))
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
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TWEETER_GITHUB_TOKEN")},
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

		config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
		token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET"))

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

	// Fetch the latest commit made in the configured folder. The sha is used to
	// get its associated file which changed with the commit. The name of this
	// file then by convention indicates the highest number of files in the
	// sequence of existing file names. The example file names below describe the
	// convention of numbers defining the file sequence. The method below fetches
	// the sha that assumedly added file055 which tells us that there are 55 files
	// to chose from randomly.
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
			Path: "philosophy",
			ListOptions: github.ListOptions{
				PerPage: 1,
			},
		}

		out, _, err := ghClient.Repositories.ListCommits(ctx, "xh3b4sd", "content", in)
		if err != nil {
			return microerror.Mask(err)
		}

		sha = out[0].GetSHA()

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

		out, _, err := ghClient.Repositories.GetCommit(ctx, "xh3b4sd", "content", sha)
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
	var number int
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "computing total number of files")

		m := regexp.MustCompile(`([0-9]+)$`).FindString(file)

		number, err = strconv.Atoi(m)
		if err != nil {
			return microerror.Mask(err)
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("computed total number of files %d", number))
	}

	var path string
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "chosing random file")

		i, err := newRandom.CreateMax(number + 1)
		if err != nil {
			return microerror.Mask(err)
		}
		n := strconv.Itoa(i)
		p := len(file) - len(regexp.MustCompile(`([0-9]+)$`).ReplaceAllString(file, "")) - len(n)

		path += "philosophy/"
		for i := 0; i < p; i++ {
			path += "0"
		}
		path += n

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("chose random file %#q", path))
	}

	var content string
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "finding content")

		in := &github.RepositoryContentGetOptions{}

		out, _, _, err := ghClient.Repositories.GetContents(ctx, "xh3b4sd", "content", path, in)
		if err != nil {
			return microerror.Mask(err)
		}

		c, err := out.GetContent()
		if err != nil {
			return microerror.Mask(err)
		}

		content = strings.TrimSpace(regexp.MustCompile(`[\n\r\t]`).ReplaceAllString(c, " "))

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("found content %#q", content))
	}

	// We just make sure that we deal with valid credentials and collect
	// information about the user on which behalf we are going to tweet below.
	//
	//     https://developer.twitter.com/en/docs/accounts-and-users/manage-account-settings/api-reference/get-account-verify_credentials
	//
	var userName string
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", "verifying twitter credentials")

		p := &twitter.AccountVerifyParams{
			SkipStatus: twitter.Bool(true),
		}
		user, _, err := twClient.Accounts.VerifyCredentials(p)
		if err != nil {
			return microerror.Mask(err)
		}

		userName = user.ScreenName

		newLogger.LogCtx(ctx, "level", "debug", "message", "verified twitter credentials")
	}

	// Once the necessary content is gathered it can be tweeted using the Twitter
	// client initialized above. For more information about the API client and API
	// specs see the following resources.
	//
	//     https://godoc.org/github.com/dghubble/go-twitter/twitter#StatusService.Update
	//     https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/post-statuses-update
	//
	{
		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("tweeting on behalf of user %#q", userName))

		_, _, err := twClient.Statuses.Update(content, nil)
		if err != nil {
			return microerror.Mask(err)
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("tweeted on behalf of user %#q", userName))
	}

	return nil
}
