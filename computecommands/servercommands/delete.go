package servercommands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jrperritt/rackcli/clients"
	"github.com/jrperritt/rackcli/util"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

// delete is a reserved word in Go.
var remove = cli.Command{
	Name:        "delete",
	Usage:       fmt.Sprintf("%s %s delete <serverID> [flags]", util.Name, commandPrefix),
	Description: "Deletes an existing server",
	Action:      commandDelete,
	Flags:       flagsDelete(),
}

func flagsDelete() []cli.Flag {
	return []cli.Flag{}
}

func commandDelete(c *cli.Context) {
	util.CheckArgNum(c, 1)
	serverID := c.Args()[0]
	client := clients.Create("compute")
	err := servers.Delete(client, serverID).ExtractErr()
	if err != nil {
		fmt.Printf("Error deleting server (%s): %s\n", serverID, err)
		os.Exit(1)
	}
}