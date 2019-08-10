# github-tweeter



```
export TWITTER_CONSUMER_KEY=$(cat ~/.credential/twitter-consumer-key)
export TWITTER_CONSUMER_SECRET=$(cat ~/.credential/twitter-consumer-secret)
export TWITTER_ACCESS_TOKEN=$(cat ~/.credential/twitter-access-token)
export TWITTER_ACCESS_SECRET=$(cat ~/.credential/twitter-access-secret)
```



```
$ go run main.go
{"caller":"github-tweeter/main.go:36","level":"debug","message":"initializing twitter client","time":"2019-08-10T13:53:01.152344+00:00"}
{"caller":"github-tweeter/main.go:43","level":"debug","message":"initialized twitter client","time":"2019-08-10T13:53:01.15255+00:00"}
{"caller":"github-tweeter/main.go:48","level":"debug","message":"verifying twitter credentials","time":"2019-08-10T13:53:01.152581+00:00"}
{"caller":"github-tweeter/main.go:61","level":"debug","message":"verified twitter credentials","time":"2019-08-10T13:53:01.749231+00:00"}
{"caller":"github-tweeter/main.go:69","level":"debug","message":"tweeting on behalf of user `xh3b4sd`","time":"2019-08-10T13:53:01.749264+00:00"}
{"caller":"github-tweeter/main.go:76","level":"debug","message":"tweeted on behalf of user `xh3b4sd`","time":"2019-08-10T13:53:02.434425+00:00"}
```
