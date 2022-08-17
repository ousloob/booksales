package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/ardanlabs/service/foundation/logger"
	"go.uber.org/zap"
)

var build = "develop"

func main() {
	// Construct the application logger.
	log, err := logger.New("SHOP-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer log.Sync()
}

func run(log *zap.SugaredLogger) error {

	// =========================================================================
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeOut    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:12s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:3000"`
			DebugHost       string        `conf:default:0.0.0.0:5000"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "© 1996 2022 Oussama Moulana",
		},
	}

	return nil
}
