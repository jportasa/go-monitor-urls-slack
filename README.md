# Monitoring URL endpoints

Defining a set of URL endpoints in conf.yaml, the script will GET these endpoints and post a slack message in a channel if the reply toke longer than the Timeout defined for each enpoint.

## Env vars

````
export SLACKWEBHOOK=https://hooks.slack.com/services/...
````
