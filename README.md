ABOUT
=====

Attitude is a command line application to manipulate user attributes in an Active Directory Setup. We use it internally at [Celluloid VFX](http://celluloid-vfx.com/) to batch edit attributes like login scripts or VPN Flags. This is used against a Samba 4.2 AD setup and not tested on something else.

BIG FAT WARNING !!!1!
======

We don't have a clue about anything! Neither LDAP, AD, LDIF or programming. If you checkout this app and *instadestroy™* your company don't blame us. Seriously! read the source code and decide if this is something you would allow to manipulate your directory service.

This code is barely tested and considered a hack. You have been warned!

INSTALLATION
============

Things you probably want to change for your environment before build:

1. Edit cell-ci.yaml with your values
2. Install Dependencies: [musl-gcc](http://www.musl-libc.org/), [upx](https://upx.github.io/), [ciparser](https://github.com/celluloidvfx/ciparser), [gometalinter](https://github.com/alecthomas/gometalinter)

Build Attitude from source
-----

This program is written in [go](https://golang.org/). With Debian based Linux systems it's the easiest way to install go via [godeb](https://github.com/niemeyer/godeb).

Beware the make file depends on our internal continuous integration tool (ciparser) and will fail if you don't have it installed. Ciparser generates values in compile time from the cell-ci.yaml. Please read and install ciparser from here.

    git clone https://github.com/celluloidvfx/attitude.git
    cd attitude
    vi cell-ci.yaml #Edit the build config
    make

Beware, you can compile the app without *make* using go's standard tools, but the app will not work, because it needs on compile-time generated variables. This behavior eventually will be dropped.

We use [musl-gcc](http://www.musl-libc.org/), [upx](https://upx.github.io/) and [gometalinter](https://github.com/alecthomas/gometalinter) as default. We statically compile libraries into the executable. You can switch this behavior off in the cell-ci.yaml. See ciparser for more information on compile options.

Tested on
-----

Attitude client on:

- Windows 7 --> Check
- Ubuntu 16.04 --> Check

Server Running

- [Sernet Samba 4.2](https://portal.enterprisesamba.com/) --> Check


RUNNING THE APP
==================

Attitude is a commandline app and is started in your terminal/cmd.

for more information see

    ./attitude --help

TODO
===

1. Right now the admin password is compiled into the app and is displayed in debug mode. We kind of like this *configless* non-interactive execution. But this is very specific to our internal needs and needs external security measures. In future an non-/interactive User/Password request should be done similar to i.e. mysqlclient.
2. Right now you can easily break your directory. This app should prevent you from doing this.
3. Ldaps. TLS encrypted communication
4. Add more attributes
5. Enhance documentation

If you see places to enhance this app (and there are some!), please file a Pull Request

Thanks to
=========

ldap library
----

Copyright (C) 2012 The Go Authors. See https://github.com/nmcclain/ldap/blob/master/LICENSE
This seems to be a fork from https://github.com/mmitton/ldap/blob/master/LICENSE

logrus
------

Copyright (c) 2014 Simon Eskildsen https://github.com/Sirupsen/logrus/blob/master/LICENSE

cli
---

Copyright (c) 2016 Jeremy Saenz & Contributors https://github.com/urfave/cli/blob/master/LICENSE


COPYRIGHT
========
// Copyright (C) 2016 Celluloid VFX and Johannes Amorosa


LICENSE
=======

```
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
```
