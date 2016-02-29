package handlers

import (
    "html/template"
    "net/http"
    "path"
    "strings"

    "github.com/goincremental/negroni-sessions"
    "github.com/julienschmidt/httprouter"
    "github.com/seriallink/workshop/database"
)

type LoginMessage struct {
    Message string
}

func LoginHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // mensagem padrao
    message := LoginMessage{"Sign In"}

    // se o user voltou da tela de login
    if strings.Contains(request.Referer(),"/login") {
        message.Message = "Login invalido"
    }

    // parse do template de login
    tmpl, err := template.ParseFiles(path.Join("templates","login.html"))

    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(response,message); err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }

}

func LoginAction(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // valida o login no banco de dados
    email := request.FormValue("email")
    password := request.FormValue("password")
    user, err := database.GetUser(email, password)

    if err == nil {

        // adiciona o user na session
        session := sessions.GetSession(request)
        session.Set("user",user)

        // redirect para home
        http.Redirect(response, request, "/home", 302)
        return

    }


    // usuario/senha nao encontrado
    if err.Error() == "sql: no rows in result set" {
        http.Redirect(response, request, "/login", 302)
        return
    }

    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }

}

func LogoutAction(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // remove o user na session
    session := sessions.GetSession(request)
    session.Delete("user")

    // redireciona para o login
    http.Redirect(response, request, "/login", 302)
    return

}