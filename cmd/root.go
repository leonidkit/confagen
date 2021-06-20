/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/leonidkit/confagen/pkg/confagen"

	"github.com/spf13/cobra"
)

var (
	srcFilePath string
	dstFilePath string
	confag      *confagen.Confagen
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "confagen",
	Short: "Generating a configuration file for a golang application from a yaml file",
	Long: `Confagen is a CLI utility for Go that generates a golang configuration file from a yaml file.
It is enough to specify the name of the source file and the file to be generated.`,
	Example: `	confagen --src=/path/to/yaml/config/config.yaml --dst=/path/to/destination/file/config.go`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		confag.Generate(cmd.Flag("src").Value.String(), cmd.Flag("dst").Value.String())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(func() {
		confag = confagen.New()
	})

	rootCmd.PersistentFlags().StringVar(&srcFilePath, "src", "", "source file")
	rootCmd.PersistentFlags().StringVar(&dstFilePath, "dst", "", "destination file")

	rootCmd.MarkPersistentFlagRequired("src")
	rootCmd.MarkPersistentFlagRequired("dst")
}
