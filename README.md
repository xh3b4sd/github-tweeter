# github-tweeter



### Setup

The Github Token needs the privileges `repo:status` and `public_repo`.



### Environment

```
export GITHUB_TWEETER_GITHUB_TOKEN=$(cat ~/.credential/github-tweeter-github-token)
export TWITTER_CONSUMER_KEY=$(cat ~/.credential/twitter-consumer-key)
export TWITTER_CONSUMER_SECRET=$(cat ~/.credential/twitter-consumer-secret)
export TWITTER_ACCESS_TOKEN=$(cat ~/.credential/twitter-access-token)
export TWITTER_ACCESS_SECRET=$(cat ~/.credential/twitter-access-secret)
```



### Example

```
$ go run main.go
{"caller":"github-tweeter/main.go:80","level":"debug","message":"initializing github client","time":"2020-08-15T12:12:46.311952+00:00"}
{"caller":"github-tweeter/main.go:88","level":"debug","message":"initialized github client","time":"2020-08-15T12:12:46.31205+00:00"}
{"caller":"github-tweeter/main.go:100","level":"debug","message":"initializing twitter client","time":"2020-08-15T12:12:46.31206+00:00"}
{"caller":"github-tweeter/main.go:107","level":"debug","message":"initialized twitter client","time":"2020-08-15T12:12:46.312135+00:00"}
{"caller":"github-tweeter/main.go:144","level":"debug","message":"finding latest commit","time":"2020-08-15T12:12:46.312146+00:00"}
{"caller":"github-tweeter/main.go:160","level":"debug","message":"found latest commit `dfabb772ad0ad89c0a794cbcefa35f6a80883694`","time":"2020-08-15T12:12:46.68702+00:00"}
{"caller":"github-tweeter/main.go:170","level":"debug","message":"finding latest file","time":"2020-08-15T12:12:46.687057+00:00"}
{"caller":"github-tweeter/main.go:179","level":"debug","message":"found latest file `philosophy/2020/0000828`","time":"2020-08-15T12:12:46.869399+00:00"}
{"caller":"github-tweeter/main.go:191","level":"debug","message":"computing total number of files","time":"2020-08-15T12:12:46.869471+00:00"}
{"caller":"github-tweeter/main.go:200","level":"debug","message":"computed total number of files 828","time":"2020-08-15T12:12:46.869553+00:00"}
{"caller":"github-tweeter/main.go:205","level":"debug","message":"choosing random number","time":"2020-08-15T12:12:46.869589+00:00"}
{"caller":"github-tweeter/main.go:212","level":"debug","message":"chose random number 476","time":"2020-08-15T12:12:46.869897+00:00"}
{"caller":"github-tweeter/main.go:260","level":"debug","message":"finding content","time":"2020-08-15T12:12:47.038595+00:00"}
{"caller":"github-tweeter/main.go:280","level":"debug","message":"found content `#sky is the #limit. What you can do and what not is only a #matter of #imagination. You can #explore the unexplored #territory. You can #go #big.`","time":"2020-08-15T12:12:47.809698+00:00"}
{"caller":"github-tweeter/main.go:290","level":"debug","message":"verifying twitter credentials for user","time":"2020-08-15T12:12:47.809785+00:00"}
{"caller":"github-tweeter/main.go:302","level":"debug","message":"verified twitter credentials for user `xh3b4sd`","time":"2020-08-15T12:12:48.061411+00:00"}
{"caller":"github-tweeter/main.go:313","level":"debug","message":"tweeting content","time":"2020-08-15T12:12:48.061456+00:00"}
{"caller":"github-tweeter/main.go:320","level":"debug","message":"tweeted content","time":"2020-08-15T12:12:48.269885+00:00"}
```
