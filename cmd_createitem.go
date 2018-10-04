package main

import (
	"strings"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

func cmdAdd2netlist(c *cli.Context) error {
	return add2netlist(c)
}

func add2netlist(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")

	// Modify existing network list
	// Since CLI tooling does not split our slice flag we will just split it on our own
	editListOpts := edgegrid.CreateNetworkListOptions{
		List: strings.Split(c.StringSlice("items")[0], ","),
	}

	netLists, _, err := apiClient.NetworkLists.AddNetworkListItems(listID, editListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil

}