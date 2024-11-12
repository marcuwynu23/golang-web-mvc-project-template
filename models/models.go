package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`

	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Age   int    `json:"age" bson:"age"`
}
