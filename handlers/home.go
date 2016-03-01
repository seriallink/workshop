package handlers

import (
    "html/template"
    "net/http"
    "path"
    "github.com/julienschmidt/httprouter"
    "github.com/goincremental/negroni-sessions"
    "github.com/seriallink/workshop/models"
)
func HomeHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // recupera o user na session
    session := sessions.GetSession(request)
    user := session.Get("user").(models.User)

    // parse do template da home
    tmpl, err := template.ParseFiles(path.Join("templates","home.html"))

    if err == nil {
        err = tmpl.Execute(response,user)
    }

    if err != nil {
        http.Error(response, err.Error(), http.StatusInternalServerError)
        return
    }

}