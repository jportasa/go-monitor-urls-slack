# Monitoring URL endpoints

Defining a set of URL endpoints in conf.yaml, the script will GET these endpoints and post a slack message in a channel if the reply toke longer than the Timeout defined for each enpoint.

## Env vars

export SLACKWEBHOOK=https://hooks.slack.com/services/...

## Log

````
2023/09/04 12:24:38 http://www.equisens.es 200
2023/09/04 12:24:38 1.5367649669999999
2023/09/04 12:24:38 http://www.cisco.com 200
2023/09/04 12:24:38 0.210603165
```
