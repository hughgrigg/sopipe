package cmd

import (
	"fmt"
	"github.com/hughgrigg/sopipe/fetch"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var location string = os.Args[2]

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch social media posts from a source",
	Long:  `Fetch a stream of social media posts from various source types`,
	Run: func(cmd *cobra.Command, args []string) {
		posts, err := fetch.PostsFromLocation(location)
		if err != nil {
			log.Fatal(err)
			return
		}
		for _, post := range posts {
			fmt.Println(fetch.PostToTabbed(post))
		}
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}
