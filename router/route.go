package router

import (
	"github.com/henrylee2cn/faygo"
	"github.com/my_project/myfaygo/handler"
	"github.com/my_project/myfaygo/middleware"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {
	ApiHandle := new(handler.Api)
	HtmlHandle := new(handler.Html)

	frame.Route(
		frame.NewNamedAPI("Index", "GET", "/", handler.Index),
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),

		frame.NewNamedGroup("use for api service", "/api",
			frame.NewGET("/countrys", ApiHandle.CountryAll()),
		),

		frame.NewNamedGroup("use for html service", "/html",
			frame.NewGET("/country_list", HtmlHandle.CountryListAll()),
		),
	)
}
