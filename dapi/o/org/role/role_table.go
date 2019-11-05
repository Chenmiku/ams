package role

import (
	"ams/dapi/o/model"
)

//TableRole : Table in DB
func (b *Role) CreateTable() {
	model.NewTable(b)
}

