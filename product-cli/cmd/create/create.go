package create

import (
	"github.com/openshift/hypershift/product-cli/cmd/cluster"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "create",
		Short:        "Commands for creating HostedClusters",
		SilenceUsage: true,
	}

	cmd.AddCommand(cluster.NewCreateCommands())

	return cmd
}