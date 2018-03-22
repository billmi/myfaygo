package handler

import (
	"github.com/henrylee2cn/faygo"
	"github.com/my_project/myfaygo/db"
	"github.com/my_project/myfaygo/model"
)

func (html *Html) CountryListAll() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		countrys := new([]*model.Country)
		err := db.DefaultEngine.Native.Select("*").Find(countrys)
		if err != nil {
			return err
		}

		return ctx.Render(200, faygo.JoinStatic("view/country/index.html"), faygo.Map{
			"country": countrys,
		})
	})
}
