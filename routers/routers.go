package routers

import (
    "github.com/julienschmidt/httprouter"
    "github.com/seriallink/workshop/handlers"
)

func GetRouters() (router *httprouter.Router) {

    router = httprouter.New()

    router.GET("/login", handlers.LoginHandler)
    router.POST("/login", handlers.LoginAction)
    router.GET("/logout", handlers.LogoutAction)

    router.GET("/", handlers.HomeHandler)
    router.GET("/home", handlers.HomeHandler)

    router.POST("/user", handlers.UserSaveHandler)
    router.POST("/pass", handlers.UserPassHandler)

    return
}
