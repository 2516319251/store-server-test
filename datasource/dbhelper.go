package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"stroe-server/config"
	"sync"
	"xorm.io/core"
)

var (
	masterEngine *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := config.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", c.User, c.Password, c.Host, c.Port, c.DbName, c.Charset)
	engine, err := xorm.NewEngine(config.DriverName, driveSource)
	if err != nil {
		log.Fatal("mysql.InstanceMaster err: ", err.Error())
	} else {
		masterEngine = engine
		masterEngine.ShowSQL(c.ShowLog)
		if c.DebugLog {
			masterEngine.Logger().SetLevel(core.LOG_DEBUG)
		}
	}

	return masterEngine
}
