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
	"github.com/urfave/cli"
)

var raw bool

func registerUserSetFlags() []cli.Flag {

	return []cli.Flag{
		// https://msdn.microsoft.com/en-us/library/ms679366(v=vs.85).aspx
		cli.StringFlag{
			Name:  "postalCode",
			Usage: "String. The postal or zip code for mail delivery. https://goo.gl/Md4vme",
		},
		// https://msdn.microsoft.com/en-us/library/ms679656(v=vs.85).aspx
		cli.StringFlag{
			Name:  "scriptPath",
			Usage: "String. This attribute specifies the path for the user's logon script. The string can be null. https://goo.gl/K3aK48",
		},
		// https://msdn.microsoft.com/en-us/library/ms678093(v=vs.85).aspx
		cli.BoolFlag{
			Name:  "msNPAllowDialin",
			Usage: "Bool. Indicates whether the account has permission to dial in to the RAS server. https://goo.gl/GytCHw",
		},
		// https://msdn.microsoft.com/en-us/library/ms680832(v=vs.85).aspx
		cli.IntFlag{
			Name:  "userAccountControl",
			Usage: "Integer. Flags that control the behavior of the user account. https://goo.gl/S6nWQr",
		},
	}
}

func registerUserGetFlags() []cli.Flag {

	return []cli.Flag{
		cli.BoolFlag{
			Name:        "raw",
			Usage:       "dump output",
			Destination: &raw,
		},
	}
}
