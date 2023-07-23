/*
Copyright © 2023 Till Hoffmann <till@thetillhoff.de>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/spell/pkg/spell"
)

var (
	nato   bool
	german bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spell",
	Short: "Spell things",
	Args:  cobra.ArbitraryArgs, // Default anyway, but make it explicit for clarity
	Run: func(cmd *cobra.Command, args []string) {
		var (
			input    string
			spelling []string
		)

		if nato && german { // Validate flags
			log.Fatalln("Illegal flag combination")
		}

		input = strings.Join(args, " ") // Merge all args into one input string, separated by spaces

		if german { // If german is selected
			spelling = spell.SpellGerman(input) // Retrieve German spelling
		} else { // If nato is selected (default)
			spelling = spell.SpellNato(input) // Retrieve NATO spelling
		}

		for _, line := range spelling { // For each spelled character
			fmt.Println(line) // Print the spelling
		}
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spell.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVar(&nato, "nato", false, "Enable NATO spell mode")
	rootCmd.Flags().BoolVar(&german, "german", false, "Enable German spell mode")
}
