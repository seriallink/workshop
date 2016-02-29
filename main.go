package main

import (
    "encoding/gob"
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/goincremental/negroni-sessions"
    "github.com/goincremental/negroni-sessions/cookiestore"
    "github.com/satori/go.uuid"

    "github.com/seriallink/workshop/database"
    "github.com/seriallink/workshop/routers"
    "github.com/seriallink/workshop/models"
)

func main() {

    // inicia o banco de dados
    database.Init()

    // fecha conexao com o banco
    defer database.Close()

    // registra o model user
    gob.Register(models.User{})

    // middleware stack
    middleware := negroni.Classic()
    middleware.Use(negroni.NewStatic(http.Dir("templates")))
    middleware.Use(sessions.Sessions("workshop",cookiestore.New([]byte(uuid.NewV4().String()))))
    middleware.UseFunc(Authenticate)
    middleware.UseHandler(routers.GetRouters())
    middleware.Run(":8080")

}

func Authenticate(response http.ResponseWriter, request *http.Request, handler http.HandlerFunc) {

    // verifica se o usuario existe na session
    user := sessions.GetSession(request).Get("user")

    if request.RequestURI == "/login" || user != nil {
        handler(response,request)
    } else {
        http.Redirect(response, request, "/login", 302)
    }

}