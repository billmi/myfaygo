package db

import (
	_ "github.com/go-sql-driver/mysql"
	native "github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo/ext/db/xorm"
	"github.com/my_project/myfaygo/global"
)

type Engine struct {
	Native *native.Engine
}

var (
	DefaultEngine *Engine
)

func init() {
	DefaultEngine = &Engine{Native: GetEngine(global.DB_DEFAULT)}
}

func GetEngine(db_name string) *native.Engine {
	engine := xorm.MustDB(db_name)
	engine.ShowSQL(true)
	return engine
}

//更新一条记录
func (engine *Engine) Update(id int64, i interface{}) error {
	_, err := engine.Native.Id(id).Update(i)
	return err
}

//修改指定的列
func (engine *Engine) UpdateCols(id int64, cols string, i interface{}) error {
	_, err := engine.Native.Id(id).Cols(cols).Update(i)
	return err
}
