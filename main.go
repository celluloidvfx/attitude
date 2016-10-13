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
	"net"
	"net/url"
	"os"

	b64 "encoding/base64"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var appVersion, appReleaseTag, appShortCommitID, appLdapServer string

func main() {
	app := cli.NewApp()
	app.Name = "Celluloid Active Directory Attribute Getter/Setter\n"
	app.Usage = "This Application delivers information on Celluloid VFX AD settings"
	app.Author = "Johannes Amorosa"
	app.Email = "johannesa@celluloid-vfx.com"
	app.EnableBashCompletion = true
	app.Version = mainVersion()

	registerGlobalFlags(app)

	app.Commands = registerCommands()

	app.Before = func(c *cli.Context) error {

		if DebugFlag {
			log.SetLevel(log.DebugLevel)
			log.Debug("Debug Mode")
		}

		var u string

		if len(overrideServerFlag) > 0 {
			u = overrideServerFlag
		} else {
			u = string(decodeCompileFlags(appLdapServer))
		}

		Host, Port = splitServerpath(u)
		return nil
	}

	app.Action = func(c *cli.Context) bool {

		cli.ShowAppHelp(c)
		return true
	}

	_ = app.Run(os.Args)

} //main

func mainVersion() string {
	s := ""
	s = s + appVersion + "\n"
	s = s + " release-Tag: " + appReleaseTag + "\n"
	s = s + " commit-ID: " + appShortCommitID + "\n"
	return s
}

func splitServerpath(i string) (string, string) {
	u, err := url.Parse(i)
	if err != nil {
		log.Fatal(err)
	}
	h, p, _ := net.SplitHostPort(u.Host)
	return h, p
}

// func (a serverpath) port() string {
// 	b := decodeCompileFlags([]byte(a))
// 	u, err := url.Parse(string(b))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, port, _ := net.SplitHostPort(u.Host)
// 	return port
// }

// func (a serverpath) host() string {
// 	b := decodeCompileFlags(a)
// 	u, err := url.Parse(string(b))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	host, _, _ := net.SplitHostPort(u.Host)
// 	return host
// }

func decodeCompileFlags(enc string) []byte {
	d, err := b64.StdEncoding.DecodeString(enc)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
