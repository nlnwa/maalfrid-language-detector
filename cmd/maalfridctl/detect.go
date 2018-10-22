// Copyright © 2017 Marius Elsfjordstrand Beck <marius.beck@nb.no>
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
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nlnwa/maalfrid-language-detector/pkg/maalfrid"
)

var detectCmd = &cobra.Command{
	Use:   "detect [text]",
	Short: "Detect the language of the text",
	Long:  `Detect the language of the text`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := detect(cmd, args); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
}

func detect(cmd *cobra.Command, args []string) error {
	port := viper.GetInt("port")
	host := viper.GetString("host")

	text := strings.Join(args, " ")

	client := maalfrid.NewApiClient(maalfrid.WithAddress(host, port))
	client.Dial()
	defer client.Hangup()

	if languages, err := client.DetectLanguage(text); err != nil {
		return err
	} else {
		for _, language := range languages {
			fmt.Printf("%s %.3f\n", language.Code, language.Count)
		}
	}

	return nil
}
