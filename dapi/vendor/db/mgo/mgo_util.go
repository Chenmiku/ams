package mgo

import (
	"ams/dapi/x/mlog"
)

var mongoDBLog = mlog.NewTagLog("MongoDB")

type M map[string]interface{}
