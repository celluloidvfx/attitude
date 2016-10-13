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
	"strconv"
	"strings"

	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/nmcclain/ldap"
)

type request struct {
	CN        string
	Attribute []attribute
}

// Host is global var
var Host string

// Port is a global var
var Port string

// LdapDomain
var appLdapDomain string

// AdminUser
var appAdminUser string

// AdminPass
var appAdminPass string

func getAttributesViaLdap(a *request) error {

	l, err := connectToLdapServer(a)

	// Set filter to user
	filter := "(cn=" + a.CN + ")"

	s := buildSearch(filter)

	// Do search
	sr, err := l.Search(s)

	if err != nil {
		log.Warn("ERROR: %s\n", err.Error())
		return err
	}

	if len(sr.Entries) > 0 {
		a := sr.Entries[0].Attributes

		if !raw {
			d, _ := json.MarshalIndent(&a, "", "  ")
			fmt.Println(string(d))
		} else {

			for _, v := range a {
				//b := sr.Entries[0].GetAttributeValue(a.Attribute)
				fmt.Println(v)
			}
		}
	}
	return err
}

func setAttributesViaLdap(a *request) error {

	l, err := connectToLdapServer(a)

	// Set filter to user
	filter := "(cn=" + a.CN + ")"

	s := buildSearch(filter)

	// Do search
	sr, err := l.Search(s)

	if err != nil {
		log.Warn("ERROR: %s\n", err.Error())
		return err
	}

	if len(sr.Entries) > 0 {
		modify := ldap.NewModifyRequest(sr.Entries[0].DN)
		//value := fmt.Sprint(a.Attribute[0].Value)
		for _, v := range a.Attribute {

			modify.Replace(v.Key, []string{fmt.Sprint(v.Value)})

			err = l.Modify(modify)

			if err != nil {
				log.Warn("ERROR: %s\n", err.Error())
				return err
			}
		}
	}
	return err
}

func connectToLdapServer(a *request) (*ldap.Conn, error) {
	p, err := strconv.Atoi(Port)
	if err != nil {

	}

	// Dial
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", Host, p))

	if err != nil {
		log.Warn("WARNING: Host ", Host, ":", Port, "\" is not answering: %s\n", err.Error())
		return l, err
	}

	if DebugFlag {
		l.Debug = true
	}
	// Bind
	domain := strings.Split(string(decodeCompileFlags(appLdapDomain)), ".")
	err = l.Bind("cn="+string(decodeCompileFlags(appAdminUser))+",cn=Users,dc="+domain[0]+",dc="+domain[1], string(decodeCompileFlags(appAdminPass)))
	if err != nil {
		log.Warn("ERROR: Cannot bind: ", err.Error())
		return nil, err
	}
	return l, err
}

func buildSearch(filter string) *ldap.SearchRequest {
	domain := strings.Split(string(decodeCompileFlags(appLdapDomain)), ".")

	return ldap.NewSearchRequest(
		"cn=Users,dc="+domain[0]+",dc="+domain[1],
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{},
		nil)
}
