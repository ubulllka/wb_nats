package main

import (
	"github.com/ubulllka/wbL0/internal/logger"
	"github.com/ubulllka/wbL0/internal/sub"
)

func main() {
	logg := logger.GetLogger()
	if err := sub.Run(); err != nil {
		logg.Fatal(err)
	}
}
