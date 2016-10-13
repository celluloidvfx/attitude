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
	"errors"

	//log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func registerUserSubCommands() cli.Commands {

	return cli.Commands{
		cli.Command{
			Name:  "set",
			Usage: "set user attribute",
			Flags: registerUserSetFlags(),
			Before: func(c *cli.Context) error {
				if len(c.Args().Get(0)) == 0 {
					err := errors.New("Missing username. Please provide me with a cn.\n")
					return err
				}

				if c.NumFlags() == 0 {
					err := errors.New("Missing attribute flags.\n")
					return err
				}

				return nil
			},
			Action: setUserAttribute,
		},
		cli.Command{
			Name:  "get",
			Usage: "get user attribute",
			Flags: registerUserGetFlags(),
			Before: func(c *cli.Context) error {
				if len(c.Args().Get(0)) == 0 {
					err := errors.New("Missing username. Please provide me with a cn.\n")
					return err
				}
				return nil
			},
			Action: getUserAttribute,
		},
	}
}
