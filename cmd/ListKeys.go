package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/itsritiksingh/inMemoryStore/pkg/store"
)

// keysCmd represents the add command
var keysCmd = &cobra.Command{
    Use:   "keys",
    Short: "list all keys",

    Run: func(cmd *cobra.Command, args []string) {
		keys := cmd.Context().Value("store")
		store := keys.(*store.Store)

        fmt.Println("result of addition is", store.GetAllKeys())	
    },
}

func init() {
    rootCmd.AddCommand(keysCmd)
}