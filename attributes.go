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
	"fmt"
	"strconv"

	//log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

type attribute struct {
	Key   string
	Value string
}

// buildAttributeRequest builds the request to the AD server. Flags are typesafe. But for
// several reasons we need to transform the values from flags to the appropriate values
// to be digestable by the AD. This list will substantially grow for a lot of  attribute that are
// exposed on the commandline. Strings are converted 1:1 if they don't nee a transformation
// see switch case below: default.
func buildAttributeRequest(c *cli.Context) (*request, error) {

	var r request
	var a attribute
	var err error

	r.CN = c.Args().Get(0)

	for _, flag := range c.App.VisibleFlags() {
		if c.IsSet(flag.GetName()) {
			a.Key = flag.GetName()

			switch {
			// Custom Transformation
			case a.Key == "userAccountControl":
				a.Value, err = userAccountControlTransformer(c.Int(flag.GetName()))

			case a.Key == "msNPAllowDialin":
				a.Value = genericBooleanTransformer(c.Bool(flag.GetName()))

			// Unspecified Transformation (StringFlags)
			default:
				a.Value = fmt.Sprint(c.Generic(flag.GetName()))
			}

			r.Attribute = append(r.Attribute, a)
		}
	}

	return &r, err
}

// userAccountControl just takes certain integers (or hex) see https://support.microsoft.com/en-us/kb/305144
func userAccountControlTransformer(i int) (string, error) {

	a := map[int]string{
		512: "NORMAL_ACCOUNT",
		514: "ACCOUNTDISABLE",
	}
	if _, ok := a[i]; ok {
		return strconv.Itoa(i), nil
	}
	return "", errors.New("Not Valid Integer, currently supported: 512, 514. Please read https://support.microsoft.com/en-us/kb/305144 for help")
}

// The Bools have to be transformed to uppercase strings
func genericBooleanTransformer(b bool) string {

	if b {
		return "TRUE"
	}
	return "FALSE"
}
