package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/itsritiksingh/inMemoryStore/pkg/store"
)

// keysCmd represents the add command
var getKey = &cobra.Command{
    Use:   "get",
    Short: "get key",

    Run: func(cmd *cobra.Command, args []string) {
		keys := cmd.Context().Value("store")
		store := keys.(*store.Store)
        
        for _,key := range args{   
            value , err := store.Get(key)
            if err != nil {
                fmt.Println(key , " not found")	
            }else{
                fmt.Printf("%s %s \n" , key,value)	
            }
            _ = value
        }
    },
}

func init() {
    rootCmd.AddCommand(getKey)
}