/*
	Get all facts for a specific node.
*/

package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func nodeFacts(c *cli.Context) {
	if c.Args().First() == "" {
		fmt.Println("Please provide the FQDN of a node.")
		return
	}

	// Initialize helpers.
	factacularInit(c)

	resp, err := pdbClient.NodeFacts(c.Args().First())
	if err != nil {
		fmt.Println(err)
	}
	for _, element := range resp {
		fmt.Printf("%30v %25v %v\n", c.Args().First(), element.Name, element.Value)
	}
}
