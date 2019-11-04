package role

import (
	"ams/dapi/o/model"
)

//TableRole : Table in DB
func (b *Role) CreateTable() {
	model.NewTable(b)
}

//Create :
func (b *Role) Create() error {
	model.Create(b)
	return nil
}

//MarkDelete :
func (b *Role) MarkDelete() error {
	model.MarkDelete(b)
	return nil
}

//Update :
func (v *Role) Update(newValue *Role) error {
	model.UpdateByID(v, newValue)
	return nil
}
