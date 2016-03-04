package controllers

import (
    "net/http"
    "github.com/goincremental/negroni-sessions"
    "github.com/julienschmidt/httprouter"
    "github.com/seriallink/workshop/models"
)

type UserController struct{
    MainController
}

func (c UserController) SaveUser(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // recupera o user na session
    session := sessions.GetSession(request)
    user := session.Get("user").(models.User)

    // atualiza model
    user.Name = request.FormValue("name")
    user.Email = request.FormValue("email")

    err := c.Get().ORM.Save(&user).Error

    if err == nil {

        // atualiza session
        session := sessions.GetSession(request)
        session.Delete("user")
        session.Set("user",user)

        // redirect para home
        http.Redirect(response, request, "/home", http.StatusFound)
        return

    }

    if err != nil {
        c.Get().Render.HTML(response, http.StatusInternalServerError, "500", nil)
    }

}

func (c UserController) SavePass(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // recupera o user na session
    session := sessions.GetSession(request)
    user := session.Get("user").(models.User)

    // atualiza model
    password := request.FormValue("password")
    confirm := request.FormValue("confirm")

    // confere as senhas
    if password == confirm {

        // atualiza o model
        user.Password = password

        // salva nova senha no banco
        err := c.Get().ORM.Save(&user).Error

        if err != nil {
            c.Get().Render.HTML(response, http.StatusInternalServerError, "500", nil)
            return
        }

    }

    // redirect para home
    http.Redirect(response, request, "/home", http.StatusFound)
    return

}