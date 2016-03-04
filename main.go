package main

import (
    "encoding/gob"
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/goincremental/negroni-sessions"
    "github.com/goincremental/negroni-sessions/cookiestore"
    "github.com/satori/go.uuid"

    "github.com/seriallink/workshop/database"
    "github.com/seriallink/workshop/models"
    "github.com/seriallink/workshop/orm"
    "github.com/seriallink/workshop/routers"
)

func main() {

    // inicia o banco de dados
    orm.Init()
    database.Init()

    // fecha conexao com o banco
    defer orm.Close()
    defer database.Close()

    // registra o model user
    gob.Register(models.User{})

    // middleware stack
    middleware := negroni.Classic()
    middleware.Use(negroni.NewStatic(http.Dir("templates")))
    middleware.Use(sessions.Sessions("workshop",cookiestore.New([]byte(uuid.NewV4().String()))))
    middleware.UseFunc(Authenticate)
    //middleware.UseHandler(routers.GetHandlers())
    middleware.UseHandler(routers.GetControllers())
    middleware.Run(":8081")

}

func Authenticate(response http.ResponseWriter, request *http.Request, handler http.HandlerFunc) {

    // verifica se o usuario existe na session
    user := sessions.GetSession(request).Get("user")

    if request.RequestURI == "/login" || user != nil {
        handler(response,request)
    } else {
        http.Redirect(response, request, "/login", http.StatusFound)
    }

}