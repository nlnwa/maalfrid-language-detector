// Copyright 2018 National Library of Norway
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"net"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/nlnwa/maalfrid/api"
	"github.com/nlnwa/maalfrid/pkg/maalfrid"
	"github.com/nlnwa/pkg/log"
)

type serveConfig struct {
	port  int
	count int
	log.Logger
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Maalfrid language detector service",
	Long:  `Maalfrid language detector service`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := serveConfig{
			port:  viper.GetInt("port"),
			count: viper.GetInt("count"),
		}
		if err := serve(cfg); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	serveCmd.Flags().Int("port", 8672, "server listening port")
	serveCmd.Flags().Int("count", 5, "number of suggested languages in replies")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("count", serveCmd.Flags().Lookup("count"))

	rootCmd.AddCommand(serveCmd)
}

func serve(cfg serveConfig) error {
	port := cfg.port
	logger := cfg.Logger
	count := cfg.count

	var grpcOpts []grpc.ServerOption

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrapf(err, "listening on %d failed", port)
	} else {
		logger.Info("API server listening", "port", port)
	}
	srv := grpc.NewServer(grpcOpts...)
	api.RegisterMaalfridServer(srv, maalfrid.NewApiServer(maalfrid.WithLimit(count)))

	return srv.Serve(listener)
}
