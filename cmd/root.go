/*
Copyright © 2024 Prithvijit Dasgupta <prithvid@umich.edu>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var filePath string
var outPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wikiextract",
	Short: "Application to extract MediaWiki compressed XML dumps",
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wikiextract.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(
		&filePath,
		"file",
		"f",
		"",
		"The file name of the XML dump file",
	)
	rootCmd.MarkFlagRequired("file")
	rootCmd.PersistentFlags().StringVarP(
		&outPath,
		"out",
		"o",
		"temp.csv",
		"The output file name",
	)
}
