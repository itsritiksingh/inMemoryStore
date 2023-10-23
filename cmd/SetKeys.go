package cmd

import (
	"log"

	"github.com/itsritiksingh/inMemoryStore/pkg/store"
	"github.com/spf13/cobra"
)

// keysCmd represents the add command
var setKey = &cobra.Command{
	Use:   "set",
	Short: "set key",

	Run: func(cmd *cobra.Command, args []string) {
		keys := cmd.Context().Value("store")
		store := keys.(*store.Store)

		if len(args) > 2 {
			log.Fatalf("space separated key value is required provided args of len %d",len(args))
		}
		//caution args need to be filtered here
		store.Upsert(args[0],args[1])
	},
}

func init() {
	rootCmd.AddCommand(setKey)
}
