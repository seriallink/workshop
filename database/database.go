package database

import (
    "database/sql"
   _"github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {

    var err error

    if db, err = sql.Open("sqlite3", "workshop.sqlite"); err == nil {
        if _, err = db.Exec("create table if not exists users(id text, name text, email text, password text)"); err == nil {
            if _, err = db.Exec("create table if not exists books(id text, title text, author text)"); err == nil {
                _, err = CreateUser()
            }
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