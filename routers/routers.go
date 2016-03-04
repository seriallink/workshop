package routers

import (
    "github.com/julienschmidt/httprouter"
    "github.com/seriallink/workshop/handlers"
    "github.com/seriallink/workshop/controllers"
)

func GetHandlers() (router *httprouter.Router) {

    router = httprouter.New()

    router.GET("/", handlers.HomeHandler)
    router.GET("/home", handlers.HomeHandler)

    router.GET("/login", handlers.LoginHandler)
    router.POST("/login", handlers.LoginAction)
    router.GET("/logout", handlers.LogoutAction)

    router.POST("/user", handlers.UserSaveHandler)
    router.POST("/pass", handlers.UserPassHandler)

    return
}

func GetControllers() (router *httprouter.Router) {

    router = httprouter.New()

    // controllers
    home := &controllers.HomeController{}
    login := &controllers.LoginController{}
    user := &controllers.UserController{}

    router.GET("/", home.Home)
    router.GET("/home", home.Home)

    router.GET("/login", login.Form)
    router.POST("/login", login.Login)
    router.GET("/logout", login.Logout)

    router.POST("/user", user.SaveUser)
    router.POST("/pass", user.SavePass)

    return
}