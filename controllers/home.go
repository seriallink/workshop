package controllers

import (
    "net/http"
    "github.com/goincremental/negroni-sessions"
    "github.com/julienschmidt/httprouter"
    "github.com/seriallink/workshop/models"
)

type HomeController struct {
    MainController
}

func (c *HomeController) Home(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // recupera o user na session
    session := sessions.GetSession(request)
    user := session.Get("user").(models.User)

    // renderiza nosso html
    c.Get().Render.HTML(response, http.StatusOK, "home", user)

}