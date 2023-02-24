// Copyright © 2023 The Interplanetary Network Authors
package cmd

import (
	"context"
	"os"

	"github.com/interplanetary-network/proxyd/internal/auth"
	"github.com/interplanetary-network/proxyd/internal/proxy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	poolID              string
	noAuth              bool
	bypass              []string
	logLevel            string
	contractAddress     string = ""
	mainNetGatewayURL   string = "https://arb1.arbitrum.io/rpc"
	goerliNetGatewayURL string = "https://goerli-rollup.arbitrum.io/rpc"
	useGoerliNet        bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "proxyd",
	Args:  cobra.MinimumNArgs(1),
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		ll, err := log.ParseLevel(logLevel)
		if err != nil {
			panic(err)
		}
		log.SetLevel(ll)

		var authenticator auth.Authenticator

		if !noAuth {
			ethereumGatewayURL := mainNetGatewayURL
			if useGoerliNet {
				ethereumGatewayURL = goerliNetGatewayURL
			}

			if contractAddress == "" {
				log.Error("contract address is required")
				return
			}

			if poolID == "" {
				log.Error("pool ID is required")
				return
			}

			authenticator = auth.NewPoolAuthenticator(poolID, contractAddress, ethereumGatewayURL, auth.NewPoolOrderCache())
			log.Info("using pool authenticator with pool ID: ", poolID)
		} else {
			authenticator = auth.NewDummyAuthenticator()
			log.Info("using dummy authenticator")
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for _, v := range args {
			px, err := proxy.ParseFromURL(v, proxy.WithAuthenticator(authenticator), proxy.WithBypass(bypass...))
			if err != nil {
				log.Error(err)
				return
			}

			go px.Run(ctx)
		}

		select {}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&poolID, "pool", "", "registered pool ID from exchange platform")
	rootCmd.Flags().StringSliceVar(&bypass, "bypass", nil, "network bypass list, separated by comma")
	rootCmd.Flags().BoolVar(&noAuth, "noauth", false, "no authentication required")
	rootCmd.Flags().StringVar(&logLevel, "log-level", "info", "log level, one of: trace, debug, info, warn/warning, error, fatal, panic")
	rootCmd.Flags().BoolVar(&useGoerliNet, "goerli", false, "to use goerli testnet")
	rootCmd.Flags().StringVar(&contractAddress, "contract", "", "contract address")
}
