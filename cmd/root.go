package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var contextAdder ctxAdder

type ctxAdder struct {
	ctx context.Context
}


var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo app to demonstrate cobra",
	Long:  `demo app to demonstrate cobra by addition`,
}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
