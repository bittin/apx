package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2023
	Description: Apx is a wrapper around multiple package managers to install packages and run commands inside a managed container.
*/

import (
	"github.com/spf13/cobra"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewRunCommand() *cmdr.Command {
	cmd := cmdr.NewCommand("run", apx.Trans("run.long"), apx.Trans("run.short"), run)
	/*
			Example: "apx run htop",
			Use:     "run <program>",
			Short:   "Run a program inside a managed container.",
			RunE:    run,
		}
		cmd.Flags().SetInterspersed(false)
	*/
	return cmd
}

func run(cmd *cobra.Command, args []string) error {

	container.Run(args...)

	return nil
}
