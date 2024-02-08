package pub

import (
	"github.com/ubulllka/wbL0/internal/config"
	"github.com/ubulllka/wbL0/internal/logger"
)

func Run() error {
	logg := logger.GetLogger()
	if err := config.InitConfig(); err != nil {
		logg.Panic(err)
		return err
	}
	if _, err := Connect(); err != nil {
		logg.Panic(err)
		return err
	}
	return nil
}
