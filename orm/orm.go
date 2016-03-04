package orm

import (
    "github.com/jinzhu/gorm"
   _"github.com/mattn/go-sqlite3"
    "github.com/seriallink/workshop/models"
)

var db gorm.DB

func Init() {

    var err error

    // abre conexao de forma muito similar
    if db, err = gorm.Open("sqlite3","workshop.sqlite"); err == nil {

        // podemos invocar metodos do sql.DB
        if err = db.DB().Ping(); err == nil {

            // habilita o log (nossas queries serao exibidas na console)
            db.LogMode(true)

            // verifica se a tabela existe
            if !db.HasTable(&models.User{}) {

                // cria tabela user baseada no nosso model
                if err = db.CreateTable(&models.User{}).Error; err == nil {

                    // cria o user default caso nao exista
                    err = db.FirstOrCreate(models.UserDefault).Error

                }

            }

        }

    }

    if err != nil {
        panic(err)
    }

}

func Get() *gorm.DB {
    return &db
}

func Close() error {
    return db.DB().Close()
}