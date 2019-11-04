package business

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type BusinessConfig struct {
	gorm.Model
	General GeneralConfig `bson:"general" json:"general"`
}

func (c BusinessConfig) String() string {
	return fmt.Sprintf("business=[%s]", c.General)
}

func (c *BusinessConfig) Check() {
	c.General.Check()
}
