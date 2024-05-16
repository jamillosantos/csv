package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jamillosantos/csv/internal/service"
)

var (
	skipHeaders bool
	separator   string
	columns     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csv",
	Short: "Parse CSV file and output data according to input.",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := service.NewService()

		input := os.Stdin
		if len(args) > 0 {
			f, err := os.Open(args[0])
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer f.Close()
			input = f
		}

		err := s.Run(context.Background(), service.RunRequest{
			Output:      os.Stdout,
			Input:       input,
			SkipHeaders: skipHeaders,
			Separator:   separator,
			Columns:     columns,
		})
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
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
	rootCmd.Flags().StringVar(&separator, "separator", ",", "CSV separator")
	rootCmd.Flags().BoolVar(&skipHeaders, "skip-headers", false, "Skip headers")
	rootCmd.Flags().StringVar(&columns, "columns", "", "Output format (example: 1,2,3)")
}
