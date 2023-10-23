package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo app to demonstrate cobra",
	Long:  `demo app to demonstrate cobra by addition`,
}

func Execute(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		programName := os.Args[0]
		os.Args = append([]string{programName},strings.Split(scanner.Text()," ")...)
		if err := rootCmd.ExecuteContext(ctx); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
