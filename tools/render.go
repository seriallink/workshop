package tools

import (
    "gopkg.in/unrolled/render.v1"
)

var r *render.Render

func InitRender() *render.Render {

    // inicializa o render
    r = render.New(render.Options{
        Directory: "templates",
        Extensions: []string{".html",".tmpl"},
    })

    return r
}

func GetRender() *render.Render {

    if r == nil {
        return InitRender()
    }

    return r

}