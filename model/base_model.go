package model

import (
	"apiproject/entity"
	"github.com/bwmarrin/snowflake"
)

type BaseModel struct {
	ID        snowflake.ID     `form:"id" binding:"-" json:"id"`
	CreatedAt *entity.JsonTime `json:"createTime"`
	UpdatedAt *entity.JsonTime `json:"updateTime"`
	DeletedAt *entity.JsonTime `json:"deleteTime"`
}
