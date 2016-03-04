package controllers

import (
    "database/sql"
    "github.com/jinzhu/gorm"
    "github.com/seriallink/workshop/database"
    "github.com/seriallink/workshop/orm"
    "github.com/seriallink/workshop/tools"
    "gopkg.in/unrolled/render.v1"
)

type MainController struct{
    DB      *sql.DB
    ORM     *gorm.DB
    Render  *render.Render
}

func (c *MainController) Get() *MainController {
    c.DB = database.Get()
    c.ORM = orm.Get()
    c.Render = tools.GetRender()
    return c
}