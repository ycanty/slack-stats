# Slack Statistics

Generate statistics reports from Slack conversation history.

## TODO

* Add `db update-info --users|--channels|--all`:
  for all users/channels in db, fetch missing info
* Don't print a warning if config file doesn't exist;
  if it does exist and there's a problem reading it, print an error and exit
* Find a way to get a user's team information (slack doesn't have it) 

## References

* [JSONPath](https://godoc.org/k8s.io/client-go/util/jsonpath)
* [GORM](https://gorm.io/)
* [Slack API](https://api.slack.com)
