package main

import (
	//"database/sql"

	"log"
	"os"

	//"github.com/Siddhant6674/ECOM/cmd/api"
	"github.com/Siddhant6674/ECOM/config"
	db "github.com/Siddhant6674/ECOM/dataBase"
	mysqlcfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewMySQLStorage(mysqlcfg.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassward,
		Addr:   config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net:    "tcp",

		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
