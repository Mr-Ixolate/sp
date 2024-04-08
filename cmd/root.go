/*
Copyright © 2024 Ixo
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/Mr-Ixolate/sp/internal/raise"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sp",
	Short: "return given text with characters replaced with unicode superscript where characters match",
	Long: `Replaces numerical and some mathmetical symbols in a string with their superscript equivalents as in some units of measurement. 
For example:

sp m2
>>> m²

Note, not suitable for entire expressions as it is not smart about what is superscripted, best used for short sections.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: code

		// Shows help if user does not provide at least one arg
		if len(args) == 0 {
			fmt.Println("Include some text with the command")
			cmd.Help()
			os.Exit(1)
		}

		math_flag, _ := cmd.Flags().GetBool("math")
		ext_flag, _ := cmd.Flags().GetBool("extra")

		find_up := regexp.MustCompile(`\^`)

		for _, value := range args {
			// decide which function to use, if there are no "^" then it can yse the simple one
			if find_up.MatchString(value) {
				raise.SuperUpReg(value, math_flag, ext_flag)
			} else {
				raise.SuperSimple(value, math_flag, ext_flag)
			}
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("math", "m", false, "add mathmatical symbols note: '-' already included")
	rootCmd.Flags().BoolP("extra", "x", false, "i and n")
}
