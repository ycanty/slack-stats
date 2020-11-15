package report

import (
	"github.com/ycanty/slack-stats/db"
)

type Api struct {
	db *db.Api
}

func NewApi(db *db.Api) (*Api, error) {
	return &Api{
		db: db,
	}, nil
}
