package handlers

import (
    "net/http"

    "github.com/goincremental/negroni-sessions"
    "github.com/julienschmidt/httprouter"

    "github.com/seriallink/workshop/database"
    "github.com/seriallink/workshop/models"
)

func UserSaveHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // recupera o user na session
    session := sessions.GetSession(request)
    user := session.Get("user").(models.User)

    // atualiza model
    user.Name = request.FormValue("name")
    user.Email = request.FormValue("email")

    err := database.SaveUser(&user)

    if err == nil {

        // atualiza session
        session := sessions.GetSession(request)
        session.Delete("user")
        session.Set("user",user)

        // redirect para home
        http.Redirect(response, request, "/home", 302)
        return

    }

    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }

}

func UserPassHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // recupera o user na session
    session := sessions.GetSession(request)
    user := session.Get("user").(models.User)

    // atualiza model
    password := request.FormValue("password")
    confirm := request.FormValue("confirm")

    // confere as senhas
    if password == confirm {

        /// salva nova senha do user
        err := database.SavePass(user.Id, password)

        if err != nil {
            http.Error(response, err.Error(), http.StatusInternalServerError)
            return
        }

    }

    // redirect para home
    http.Redirect(response, request, "/home", 302)
    return

}