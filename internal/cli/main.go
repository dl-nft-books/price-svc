package cli

import (
	"context"
	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/dl-nft-books/price-svc/internal/config"
	"github.com/dl-nft-books/price-svc/internal/service"
	updatePricer "github.com/dl-nft-books/price-svc/internal/service/runners"
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	app := kingpin.New("price-svc", "")

	runCmd := app.Command("run", "run command")
	serviceCmd := runCmd.Command("service", "run service")                  // you can insert custom help
	updatePricerCmd := runCmd.Command("update-pricer", "run update-pricer") // you can insert custom help

	// custom commands go here...

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case serviceCmd.FullCommand():
		service.Run(cfg)
	case updatePricerCmd.FullCommand():
		updatePricer.RunUpdatePricer(context.Background(), cfg)
	// handle any custom commands here in the same way
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}
	return true
}
