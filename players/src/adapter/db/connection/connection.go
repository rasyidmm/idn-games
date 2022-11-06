package connection

import (
	"fmt"
	"gorm.io/gorm"
	"players/internal/config"
	dbConf "players/internal/config/db"
)

var GamesDB *gorm.DB

func init() {
	var err error
	cfg := config.GetConfig()
	fmt.Println("runnn")

	GamesDB, err = NewsConnection(cfg.Database.Games.Mysql)
	if err != nil {
		fmt.Println("Error in connection to database: ", err)
	}

}

func NewsConnection(db dbConf.Database) (*gorm.DB, error) {
	driver, err := NewInstanceDb(db)
	if err != nil {
		return nil, err
	}
	return driver.Db().(*gorm.DB), nil
}
