package sqlitedb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Connection struct {
	Db *gorm.DB
}

func NewConnection() *Connection {
	//mysql
	// connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=5s&readTimeout=10s&writeTimeout=10s&charset=utf8&parseTime=True&loc=Local", config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Endpoint, config.Config.Database.Port, config.Config.Database.DbName)

	//sqlite
	db, err := gorm.Open(sqlite.Open("./data/boilerplate.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Connection{
		Db: db,
	}
}
