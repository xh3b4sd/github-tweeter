# github-tweeter



### Setup

The Github Token needs the privileges `repo:status` and `public_repo`.



### Example

```
$ go run main.go
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"initializing github client", "cal":"github-tweeter/main.go:92" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"initialized github client", "cal":"github-tweeter/main.go:100" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"initializing twitter client", "cal":"github-tweeter/main.go:112" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"initialized twitter client", "cal":"github-tweeter/main.go:119" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"finding latest commit", "cal":"github-tweeter/main.go:143" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"found latest commit `6eb2279e2a3e3733b5ae5ae6a3d251e50f8a44ff`", "cal":"github-tweeter/main.go:180" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"finding latest file", "cal":"github-tweeter/main.go:190" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"found latest file `philosophy/2022/0001207`", "cal":"github-tweeter/main.go:199" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"computing total number of files", "cal":"github-tweeter/main.go:211" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"computed total number of files 1207", "cal":"github-tweeter/main.go:220" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"choosing random number", "cal":"github-tweeter/main.go:225" }
{ "tim":"2022-12-08 15:12:34", "lev":"inf", "mes":"chose random number 700", "cal":"github-tweeter/main.go:229" }
{ "tim":"2022-12-08 15:12:35", "lev":"inf", "mes":"finding content", "cal":"github-tweeter/main.go:277" }
{ "tim":"2022-12-08 15:12:36", "lev":"inf", "mes":"found content `One important lesson to learn is that people you do not like may also have good ideas. You cannot just write everyone off you do not have sympathy for.`", "cal":"github-tweeter/main.go:297" }
{ "tim":"2022-12-08 15:12:36", "lev":"inf", "mes":"verifying twitter credentials for user", "cal":"github-tweeter/main.go:307" }
{ "tim":"2022-12-08 15:12:36", "lev":"inf", "mes":"verified twitter credentials for user `xh3b4sd`", "cal":"github-tweeter/main.go:319" }
{ "tim":"2022-12-08 15:12:36", "lev":"inf", "mes":"generating image via prompt", "cal":"github-tweeter/main.go:324" }
{ "tim":"2022-12-08 15:12:44", "lev":"inf", "mes":"generated image via prompt", "cal":"github-tweeter/main.go:340" }
{ "tim":"2022-12-08 15:12:44", "lev":"inf", "mes":"uploading media", "cal":"github-tweeter/main.go:345" }
{ "tim":"2022-12-08 15:12:47", "lev":"inf", "mes":"uploaded media 1600871014170255368", "cal":"github-tweeter/main.go:358" }
{ "tim":"2022-12-08 15:12:47", "lev":"inf", "mes":"tweeting content", "cal":"github-tweeter/main.go:369" }
{ "tim":"2022-12-08 15:12:48", "lev":"inf", "mes":"tweeted content at https://twitter.com/xh3b4sd/status/1600871016846311425", "cal":"github-tweeter/main.go:382" }
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

export OPENAI_API_KEY=$(cat ~/.credential/openai-api-key)
export GITHUB_TWEETER_GITHUB_TOKEN=$(cat ~/.credential/github-tweeter-github-token)
export GITHUB_TWEETER_TWITTER_ACCESS_SECRET=$(cat ~/.credential/github-tweeter-twitter-access-secret)
export GITHUB_TWEETER_TWITTER_ACCESS_TOKEN=$(cat ~/.credential/github-tweeter-twitter-access-token)
export GITHUB_TWEETER_TWITTER_CONSUMER_KEY=$(cat ~/.credential/github-tweeter-twitter-consumer-key)
export GITHUB_TWEETER_TWITTER_CONSUMER_SECRET=$(cat ~/.credential/github-tweeter-twitter-consumer-secret)

/Users/xh3b4sd/project/xh3b4sd/github-tweeter/github-tweeter
```

```
export VISUAL=vim
crontab -e
```

```
$ crontab -l
0 */4 * * * /Users/xh3b4sd/.script/github-tweeter.sh
```
