package database

import (
    "database/sql"
   _"github.com/mattn/go-sqlite3"
    "github.com/seriallink/workshop/models"
)

var db *sql.DB

func Init() {

    var err error

    if db, err = sql.Open("sqlite3", "workshop.sqlite"); err == nil {
        if _, err = db.Exec("create table if not exists users(id text, name text, email text, password text)"); err == nil {
            err = CreateUser(models.UserDefault)
        }
    }

    if err != nil {
        panic(err)
    }

}

func Get() *sql.DB {
    return db
}

func Close() error {
    return db.Close()
}