# Slack CLI

Interact with slack from the command line.

## TODO

* `slack get-conversation-history`: add parameter --after <msg_id>
* Add `db get-last-message`: gets the msg_id of last imported message
* Add `db update-info --users|--channels|--all`:
  for all users/channels in db, fetch missing info
* Don't print a warning if config file doesn't exist;
  if it does exist and there's a problem reading it, print an error and exit
* Find a way to get a user's team information (slack doesn't have it) 

## References

* [GORM](https://gorm.io/)
* [Slack API](https://api.slack.com)
