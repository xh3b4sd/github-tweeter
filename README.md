# github-tweeter



### Setup

The Github Token needs the privileges `repo:status` and `public_repo`.



### Example

```
$ go run main.go
{ "caller":"github-tweeter/main.go:79", "level":"info", "message":"initializing github client", "time":"2021-01-09 16:39:54" }
{ "caller":"github-tweeter/main.go:87", "level":"info", "message":"initialized github client", "time":"2021-01-09 16:39:54" }
{ "caller":"github-tweeter/main.go:99", "level":"info", "message":"initializing twitter client", "time":"2021-01-09 16:39:54" }
{ "caller":"github-tweeter/main.go:106", "level":"info", "message":"initialized twitter client", "time":"2021-01-09 16:39:54" }
{ "caller":"github-tweeter/main.go:159", "level":"info", "message":"finding latest commit", "time":"2021-01-09 16:39:54" }
{ "caller":"github-tweeter/main.go:196", "level":"info", "message":"found latest commit `2864a540d55df17a61f771d6a3ca42a01e8eab6c`", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:206", "level":"info", "message":"finding latest file", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:215", "level":"info", "message":"found latest file `philosophy/2020/0001004`", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:227", "level":"info", "message":"computing total number of files", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:236", "level":"info", "message":"computed total number of files 1004", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:241", "level":"info", "message":"choosing random number", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:248", "level":"info", "message":"chose random number 55", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:296", "level":"info", "message":"finding content", "time":"2021-01-09 16:39:55" }
{ "caller":"github-tweeter/main.go:316", "level":"info", "message":"found content `You can operate within the range of your #abilities and #opportunities. You can #live up to your #potential so to speak. This is what #people mean when they say you can be and do anything. Your potential is the #limit. You are your single worst #enemy in each and every #fight.`", "time":"2021-01-09 16:39:56" }
{ "caller":"github-tweeter/main.go:326", "level":"info", "message":"verifying twitter credentials for user", "time":"2021-01-09 16:39:56" }
{ "caller":"github-tweeter/main.go:338", "level":"info", "message":"verified twitter credentials for user `xh3b4sd`", "time":"2021-01-09 16:39:56" }
{ "caller":"github-tweeter/main.go:349", "level":"info", "message":"tweeting content", "time":"2021-01-09 16:39:56" }
{ "caller":"github-tweeter/main.go:356", "level":"info", "message":"tweeted content", "time":"2021-01-09 16:39:56" }
```


### Automation

```
$ cat ~/.script/github-tweeter.sh
#!/bin/bash

# This script is executed by a crontab every 4 hours in order to automatically
# tweet philosophical lines from my content repository.
#
#     https://github.com/xh3b4sd/content
#

export GITHUB_TWEETER_GITHUB_TOKEN=$(cat ~/.credential/github-tweeter-github-token)
export GITHUB_TWEETER_TWITTER_ACCESS_SECRET=$(cat ~/.credential/github-tweeter-twitter-access-secret)
export GITHUB_TWEETER_TWITTER_ACCESS_TOKEN=$(cat ~/.credential/github-tweeter-twitter-access-token)
export GITHUB_TWEETER_TWITTER_CONSUMER_KEY=$(cat ~/.credential/github-tweeter-twitter-consumer-key)
export GITHUB_TWEETER_TWITTER_CONSUMER_SECRET=$(cat ~/.credential/github-tweeter-twitter-consumer-secret)

/Users/xh3b4sd/project/xh3b4sd/github-tweeter/github-tweeter
```

```
$ crontab -l
0 */4 * * * /Users/xh3b4sd/.script/github-tweeter.sh
```
