# Slack Statistics

Generate statistics reports from Slack conversation history.

## TODO

* slack get-conversation-history: Filter out channel join/leave messages
* Don't print a warning if config file doesn't exist;
  if it does exist and there's a problem reading it, print an error and exit
* Find a way to get a user's team information (slack doesn't have it) 

Reports:

* Total number of messages
* Average number of messages per day (total)
* Average number of messages each day of the week
* Total number of users
* Number of messages per user
* Top 10 messages with the highest thread count

## Usage

```bash
go run main.go slack get-conversation-history -s 2020-01-01 > tmp/data.json
go run main.go db import -f tmp/data.json
go run main.go db update-names
```
### JSONPath support

JSON output can be filtered using `--jsonpath`.  The syntax is the same
as supported by kubernetes `kubectl`.  See [JSONPath Support](https://kubernetes.io/docs/reference/kubectl/jsonpath/)

For example, the following will output in markdown syntax a link to each message using the message's timestamp
and its permalink as follows: `[timestamp](permalink)`

```bash
go run main.go slack get-conversation-history --jsonpath="{range .messages[*]}{'['}{.Timestamp}{']('}{.Permalink}{')\n'}{end}"
```

## References

* [JSONPath](https://godoc.org/k8s.io/client-go/util/jsonpath)
* [GORM](https://gorm.io/)
* [Slack API](https://api.slack.com)
