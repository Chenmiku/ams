package shared

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

func StrToInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
