package main

import (
	"errors"
	"log"
	"os"

	"github.com/denysovkos/isup/cmd/isup"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			domain := cCtx.Args().Get(0)
			if len(domain) == 0 {
				domainMissing := errors.New("domain not passed")
				log.Fatal(domainMissing)
			}

			isup.IsUp(domain)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
