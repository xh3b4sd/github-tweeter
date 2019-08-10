package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
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

		_, _, err := twClient.Statuses.Update("TODO", nil)
		if err != nil {
			return microerror.Mask(err)
		}

		newLogger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("tweeted on behalf of user %#q", userName))
	}

	return nil
}
