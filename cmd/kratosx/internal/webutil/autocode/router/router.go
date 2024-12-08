package router

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/kratosx/cmd/kratosx/internal/webutil/autocode/handler"
)

func Register(router *http.Router) {
	router.GET("/", handler.ListModel)
}
