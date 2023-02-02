package cmd

import (
	"os"

	"errors"

	"github.com/nathan-boyd/multisearch/search"
	"github.com/spf13/cobra"
)

// the name of the collection chosen by the user for search, if no collection is selected use all
var userSelectedCollection string = "ALL"
var verbose bool
var cfgFile string

var error_no_args = errors.New("not enough arguments")
var error_to_many_args = errors.New("too many arguments, multi word queries should be quoted")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "multisearch",
	Short: "search a configured site collection",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return error_no_args
		}
		if len(args) > 1 {
			return error_to_many_args
		}
		search_query := args[0]
		return search.SearchCollection(search_query, userSelectedCollection, cfgFile)
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.multisearch.json)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&userSelectedCollection, "collection", "c", "", "")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
