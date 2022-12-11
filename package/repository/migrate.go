package repository

import (
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/mrboburs/Norbekov/config"
	"github.com/mrboburs/Norbekov/util/logrus"
	//file is needed for migration url
)

func migrateInit() string {
	config := config.Config()

	// URL for Migration
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDatabase,
	)
	return url
}

// MigrateUP ...
func MigrateUP(logrus *logrus.Logger) {
	url := migrateInit()
	m, err := migrate.New("file://schema", url)
	if err != nil {
		logrus.Fatal("error in creating migrations: ", err.Error())
	}
	fmt.Printf("")
	if err := m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			logrus.Info("Error in migrating ", err.Error())
			panic(err)
		}
	}
}
