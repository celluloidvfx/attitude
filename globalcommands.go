// This Application delivers information on Celluloid VFX AD settings
// Copyright (C) 2016 Celluloid VFX and Johannes Amorosa

// This file is part of attitude.

// attitude is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Attitude is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Attitude.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func registerCommands() cli.Commands {
	return cli.Commands{

		cli.Command{
			Name:        "user",
			Usage:       "get/set user attribute",
			Subcommands: registerUserSubCommands(),
			Action: func(c *cli.Context) error {
				log.Debug("Subcommand User")
				_ = cli.ShowSubcommandHelp(c)

				return nil

			}, //app.Action
		},

		cli.Command{
			Name:  "server",
			Usage: "show compiled-in server settings",
			Action: func(c *cli.Context) error {
				fmt.Printf("Server: %s\n", Host)
				fmt.Printf("Port: %s\n", Port)
				return nil

			}, //app.Action
		},
	}
}
