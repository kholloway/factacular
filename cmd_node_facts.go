/*
	Get all facts for a specific node.
*/

package main

import (
	"fmt"
	"sort"

	"github.com/codegangsta/cli"
	"github.com/temal-/go-puppetdb"
)

type ByName []puppetdb.FactJson

func (slice ByName) Len() int {
	return len(slice)
}

func (slice ByName) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice ByName) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

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

	sort.Sort(ByName(resp))
	for _, element := range resp {
		fmt.Printf("%-30v %-25v %v\n", c.Args().First(), element.Name, element.Value)
	}
}
