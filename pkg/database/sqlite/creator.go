package sqlite

import (
	"sync"

	"terralist/pkg/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Creator struct{}

var (
	lock = &sync.Mutex{}
)

func (t *Creator) New(config database.Configurator, migrator database.Migrator) (database.Engine, error) {
	lock.Lock()
	defer lock.Unlock()

	cfg := config.(*Config)

	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &database.DefaultEngine{
		Handle:   db,
		Migrator: migrator,
	}, nil
}