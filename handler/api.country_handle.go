package handler

import (
	"net/http"

	"github.com/henrylee2cn/faygo"
	"github.com/my_project/myfaygo/db"
	"github.com/my_project/myfaygo/model"
)

func (api *Api) CountryAll() faygo.HandlerFunc {
	return faygo.HandlerFunc(func(ctx *faygo.Context) error {
		countrys := new([]*model.Country)
		err := db.DefaultEngine.Native.Select("*").Find(countrys)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, model.DataInfo{"", countrys}, true)
	})
}
