package sub

import (
	"github.com/ubulllka/wbL0/internal/cache"
	"github.com/ubulllka/wbL0/internal/config"
	"github.com/ubulllka/wbL0/internal/db"
	"github.com/ubulllka/wbL0/internal/logger"
	"github.com/ubulllka/wbL0/internal/transport"
)

func Run() error {
	logg := logger.GetLogger()

	if err := config.InitConfig(); err != nil {
		logg.Panic(err)
		return err
	}
	
	if err := db.InitializeDB(); err != nil {
		logg.Panic(err)
		return err
	}

	if err := cache.Restore(); err != nil {
		logg.Panic(err)
		return err
	}

	go Connect()

	if err := transport.InitRouter(); err != nil {
		logg.Panic(err)
		return err
	}

	return nil
}

