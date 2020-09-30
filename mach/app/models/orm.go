package models

import (
	"github.com/go-pg/pg/v10"
	"os"
)

var DB *pg.DB

func ORMInit() {
	opt, err := pg.ParseURL(os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(opt)
}