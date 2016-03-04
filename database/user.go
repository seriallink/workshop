package database

import (
    "github.com/seriallink/workshop/models"
)

func CreateUser(user *models.User) (err error) {

    // checa se o usuario existe
    count := -1

    if count, err = UserExists(user.Id); count == 0 {
        stm := "insert into users (id, name, email, password) values (?, ?, ?, ?)"
        _, err = db.Exec(stm, user.Id, user.Name, user.Email, user.Password)
    }

    return
}

func SaveUser(user *models.User) (err error) {

    stm := `update users
               set name = ?
                 , email = ?
              where id = ?`

    _, err = db.Exec(stm, user.Name, user.Email, user.Id)

    return
}

func SavePass(id, pass string) (err error) {

    stm := `update users
               set password = ?
              where id = ?`

    _, err = db.Exec(stm, pass, id)

    return
}

func UserExists(id string) (count int, err error) {

    stm := `select count(0) cnt
              from users
             where id = ?`

    err = db.QueryRow(stm,id).Scan(&count)

    return
}

func GetUser(email, pass string) (user models.User, err error) {

    stm := `select id, name, email
              from users
             where email = ?
               and password = ?`

    err = db.QueryRow(stm, email, pass).Scan(&user.Id, &user.Name, &user.Email)

    return
}