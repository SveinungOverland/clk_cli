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

	const NEWEST_VERSION int = 2
	if version != NEWEST_VERSION {
		fmt.Println("DB version is wrong, performing migration")
		Client.AutoMigrate(&models.ToDo{})
		Client.Raw(fmt.Sprint("PRAGMA user_version = ", NEWEST_VERSION)).Scan(nil)
		Client.Raw("PRAGMA user_version").Scan(&version)
		fmt.Println("New version is", version)
	}
}
