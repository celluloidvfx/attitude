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

	//log "github.com/Sirupsen/logrus"
	"encoding/json"
	"github.com/urfave/cli"
)

func setUserAttribute(c *cli.Context) error {

	req, err := buildAttributeRequest(c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if DebugFlag {
		d, _ := json.MarshalIndent(&req, "", "  ")
		fmt.Println(string(d))
	}

	err = setAttributesViaLdap(req)
	if err != nil {
		fmt.Println("request failed.")
		return err
	}
	return nil

}

func getUserAttribute(c *cli.Context) error {

	req, err := buildAttributeRequest(c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = getAttributesViaLdap(req)
	if err != nil {
		fmt.Println("request failed.")
		return err
	}
	return nil
}
