# github-tweeter



```
export GITHUB_TWEETER_GITHUB_TOKEN=$(cat ~/.credential/github-tweeter-github-token)
export TWITTER_CONSUMER_KEY=$(cat ~/.credential/twitter-consumer-key)
export TWITTER_CONSUMER_SECRET=$(cat ~/.credential/twitter-consumer-secret)
export TWITTER_ACCESS_TOKEN=$(cat ~/.credential/twitter-access-token)
export TWITTER_ACCESS_SECRET=$(cat ~/.credential/twitter-access-secret)
```



```
$ go run main.go
{"caller":"github-tweeter/main.go:79","level":"debug","message":"initializing github client","time":"2019-08-11T12:19:09.423082+00:00"}
{"caller":"github-tweeter/main.go:87","level":"debug","message":"initialized github client","time":"2019-08-11T12:19:09.423237+00:00"}
{"caller":"github-tweeter/main.go:99","level":"debug","message":"initializing twitter client","time":"2019-08-11T12:19:09.423274+00:00"}
{"caller":"github-tweeter/main.go:106","level":"debug","message":"initialized twitter client","time":"2019-08-11T12:19:09.423342+00:00"}
{"caller":"github-tweeter/main.go:143","level":"debug","message":"finding latest commit","time":"2019-08-11T12:19:09.423354+00:00"}
{"caller":"github-tweeter/main.go:159","level":"debug","message":"found latest commit `6857e548e89f4f91c51473645295d2e1f7b92f2f`","time":"2019-08-11T12:19:09.855181+00:00"}
{"caller":"github-tweeter/main.go:169","level":"debug","message":"finding latest file","time":"2019-08-11T12:19:09.855234+00:00"}
{"caller":"github-tweeter/main.go:178","level":"debug","message":"found latest file `philosophy/0000422`","time":"2019-08-11T12:19:10.166257+00:00"}
{"caller":"github-tweeter/main.go:190","level":"debug","message":"computing total number of files","time":"2019-08-11T12:19:10.166299+00:00"}
{"caller":"github-tweeter/main.go:199","level":"debug","message":"computed total number of files 422","time":"2019-08-11T12:19:10.166346+00:00"}
{"caller":"github-tweeter/main.go:210","level":"debug","message":"choosing random file","time":"2019-08-11T12:19:10.166369+00:00"}
{"caller":"github-tweeter/main.go:226","level":"debug","message":"chose random file `philosophy/0000146`","time":"2019-08-11T12:19:10.166453+00:00"}
{"caller":"github-tweeter/main.go:231","level":"debug","message":"finding content","time":"2019-08-11T12:19:10.166486+00:00"}
{"caller":"github-tweeter/main.go:247","level":"debug","message":"found content `The #nice shit is just #around #the #corner. You just need to #pick #it #up and #make #candy out of it.`","time":"2019-08-11T12:19:10.376245+00:00"}
{"caller":"github-tweeter/main.go:257","level":"debug","message":"verifying twitter credentials for user","time":"2019-08-11T12:19:10.376297+00:00"}
{"caller":"github-tweeter/main.go:269","level":"debug","message":"verified twitter credentials for user `xh3b4sd`","time":"2019-08-11T12:19:10.626737+00:00"}
{"caller":"github-tweeter/main.go:280","level":"debug","message":"tweeting content","time":"2019-08-11T12:19:10.6268+00:00"}
{"caller":"github-tweeter/main.go:287","level":"debug","message":"tweeted content","time":"2019-08-11T12:19:10.826546+00:00"}
```
