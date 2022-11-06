package connection

import (
	"errors"
	dbConf "games/internal/config/db"
	"games/src/shared/enum"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	errorInvalidDbInstance                     = errors.New("Invalid db instance")
	instanceDb             map[string]DbDriver = make(map[string]DbDriver)
)

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
}

// NewInstanceDb is used to create a new instance DB
func NewInstanceDb(config dbConf.Database) (DbDriver, error) {
	var err error
	var dbName = config.Dbname

	switch config.Adapter {
	case enum.MySql:
		dbConn, sqlErr := NewMySQLDriver(config)
		if sqlErr != nil {
			err = sqlErr
			log.Fatal("Database connection failed.")
		}
		instanceDb[dbName] = dbConn
	default:
		err = errorInvalidDbInstance
	}

	return instanceDb[dbName], err
}
