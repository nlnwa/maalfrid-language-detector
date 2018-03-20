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
	"os"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nlnwa/maalfrid-language-detector/pkg/version"
	"github.com/nlnwa/pkg/log"
	"github.com/nlnwa/pkg/logfmt"
)

var debug bool
var logger = getLogger()

var rootCmd = &cobra.Command{
	Use:   "maalfrid",
	Short: "Twitter API client",
	Long:  `Twitter API client`,
}

func init() {
	cobra.OnInitialize(func() {
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viper.AutomaticEnv()
	})
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug")
}

func getLogger() log.Logger {
	l := log15.New()
	logHandler := log15.CallerFuncHandler(log15.StreamHandler(os.Stdout, logfmt.LogbackFormat()))
	if debug {
		l.SetHandler(log15.CallerStackHandler("%+v", logHandler))
	} else {
		l.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, logHandler))
	}
	return l
}

func main() {
	logger.Info(version.String(), "app", "maalfrid")

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
