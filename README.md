# Slack CLI

Interact with slack from the command line.

## TODO

* Refactor: Add subcommands: db, report (for the stats), slack
* `slack get-conversation-history`: add parameter --after <msg_id>
* Add `db get-last-message`: gets the msg_id of last imported message
* Add `db update-info --users|--channels|--all`:
  for all users/channels in db, fetch missing info
* Refactor: Move argparse_utils to its own argparse package
* Find a way to get a user's team information (slack doesn't have it) 

## References

* [GORM](https://gorm.io/)
* [Slack API](https://api.slack.com)
