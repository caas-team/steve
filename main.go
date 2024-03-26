package main

import (
	"os"

	"github.com/caas-team/steve/pkg/debug"
	stevecli "github.com/caas-team/steve/pkg/server/cli"
	"github.com/caas-team/steve/pkg/version"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	config      stevecli.Config
	debugconfig debug.Config
)

func main() {
	app := cli.NewApp()
	app.Name = "steve"
	app.Version = version.FriendlyVersion()
	app.Usage = ""
	app.Flags = append(
		stevecli.Flags(&config),
		debug.Flags(&debugconfig)...)
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(_ *cli.Context) error {
	ctx := signals.SetupSignalContext()
	debugconfig.MustSetupDebug()
	s, err := config.ToServer(ctx)
	if err != nil {
		return err
	}
	return s.ListenAndServe(ctx, config.HTTPSListenPort, config.HTTPListenPort, nil)
}
