//   Onix Config Manager - Dbman
//   Copyright (c) 2018-2020 by www.gatblau.org
//   Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
//   Contributors to this project, hereby assign copyright in this code to the project,
//   to be licensed under the same terms as the rest of the code.
package cmd

import (
	"github.com/spf13/cobra"
)

type DbCmd struct {
	cmd *cobra.Command
}

func NewDbCmd() *DbCmd {
	c := &DbCmd{
		&cobra.Command{
			Use:   "db",
			Short: "Manage a database",
			Long:  `Access database maintenance commands.`,
		},
	}
	return c
}
