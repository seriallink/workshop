package controllers

import (
    "net/http"
    "strings"
    "github.com/goincremental/negroni-sessions"
    "github.com/jinzhu/gorm"
    "github.com/julienschmidt/httprouter"
    "github.com/seriallink/workshop/models"
)

type LoginController struct {
    MainController
}

type LoginMessage struct {
    Message string
}

func (c *LoginController) Form(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // mensagem padrao
    message := LoginMessage{"Sign In"}

    // se o user voltou da tela de login
    if strings.Contains(request.Referer(),"/login") {
        message.Message = "Login invalido"
    }

    // renderiza o form de login
    c.Get().Render.HTML(response, http.StatusOK, "login", nil)

}

func (c *LoginController) Login(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // inicializa o model com os parametros de busca
    user := &models.User{
        Email: request.FormValue("email"),
        Password: request.FormValue("password"),
    }

    // procura o login no banco de dados
    err := c.Get().ORM.Where(user).Find(user).Error

    if err == nil {

        // adiciona o user na session
        session := sessions.GetSession(request)
        session.Set("user",user)

        // redirect para home
        http.Redirect(response, request, "/home", http.StatusFound)
        return

    }

    // usuario/senha nao encontrado
    if err == gorm.RecordNotFound {
        http.Redirect(response, request, "/login", http.StatusFound)
        return
    }

    if err != nil {
        c.Get().Render.HTML(response, http.StatusInternalServerError, "500", nil)
    }

}

func (c *LoginController) Logout(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // remove o user na session
    session := sessions.GetSession(request)
    session.Delete("user")

    // redireciona para o login
    http.Redirect(response, request, "/login", http.StatusFound)
    return

}