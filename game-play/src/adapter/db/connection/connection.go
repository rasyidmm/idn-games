package connection

import (
	"fmt"
	"game-play/internal/config"
	dbConf "game-play/internal/config/db"
	"gorm.io/gorm"
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
