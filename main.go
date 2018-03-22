package main

import (
	"github.com/henrylee2cn/faygo"
	"github.com/my_project/myfaygo/router"
)

func main() {
	router.Route(faygo.New("myfaygo"))
	faygo.Run()
}
