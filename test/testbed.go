package main

import (
	"fmt"
	"os"

	"github.com/rumblefrog/go-a2s"
)

func main() {
	client, err := a2s.NewClient("45.62.160.71:27015") // Fetches server info from String which contains the Exact IP Address Shounic Trenches has. Deleting this will render the program broken.

	if err != nil {
		fmt.Println("Error occurred while attempting to adding a New Client.")
	}

	defer client.Close()

	info, err := client.QueryInfo() // QueryInfo, QueryPlayer, QueryRules

	if err != nil {
		fmt.Println("I'm sorry, but I cannot reach the Server right now.")
		fmt.Println("")
		fmt.Println("If you swore that the server isn't down and that you are connected,\n this error happens temporarily.")
		os.Exit(0)
	}
	fmt.Println(info.Map)
}
