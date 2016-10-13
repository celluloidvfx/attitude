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
	//log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// DebugFlag is a global variable to set the Debug Flag
var DebugFlag bool

// Serverflags
var overrideServerFlag string

func registerGlobalFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "overrideserver",
			Usage:       "Specify an alternativ Server URL like this: ldap://yourserver.com:387",
			Destination: &overrideServerFlag,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "More Output to StdOut",
			Destination: &DebugFlag,
		},
	}
}
