package db

import (
	"clk/db/models"
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func init() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	Client, err = gorm.Open(
		sqlite.Open(fmt.Sprint(home, "/.clk_db.db")),
		&gorm.Config{
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	var version int
	Client.Raw("PRAGMA user_version").Scan(&version)

	if version != 1 {
		fmt.Println("DB version is wrong, performing migration")
		Client.AutoMigrate(&models.ToDo{})
		Client.Raw("PRAGMA user_version = 1")
	}
}
