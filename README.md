
# Attitude [![Go Report Card](https://goreportcard.com/badge/github.com/celluloidvfx/attitude)](https://goreportcard.com/report/github.com/celluloidvfx/attitude)

Attitude is a command line application written in Go (golang) to manipulate user attributes in an [Active Directory Setup](https://wiki.samba.org/index.php/Setup_a_Samba_Active_Directory_Domain_Controller). We use it at [Celluloid VFX](http://celluloid-vfx.com/) to automatically edit attributes like login scripts paths (scriptPath).

If you find this useful please file bugs and pull requests.



**BIG FAT WARNING !!!1!**

This is used against a Samba 4.2 Active Directory Setup and not tested on something else.

> We don't have a clue about anything! Neither LDAP, AD, LDIF or programming. If you checkout this app and *instadestroy™* your company don't blame us. Seriously! read the source code and decide if this is something you would allow to manipulate your directory service.

This code is barely tested and considered a hack. You have been warned!

## Installation

### Dependencies
 - Linux
 - Go (golang) > 1.7.1 With Debian based Linux systems it's the easiest way to install go via [godeb](https://github.com/niemeyer/godeb).
 - [ciparser](https://github.com/celluloidvfx/ciparser) - a celluloid vfx build tool

### Build Attitude from source

```
git clone https://github.com/celluloidvfx/attitude.git
cd attitude
```
Beware the make file depends on our internal continuous integration tool called ciparser and will fail if you don't have it installed. Ciparser generates values in compile time from the cell-ci.yaml.

[Read more here](https://github.com/celluloidvfx/ciparser#custom-variables)


> Things you should change for your environment before build. Edit following values in cell-ci.yaml

```
    customvars:
      - name: LdapServer
        value: "ldap://yourserver.com:389"
      - name: LdapDomain
        value: "celluloidvfx.inc"
      - name: AdminUser
        value: admin
      - name: AdminPass
        value: "yourpass"
```

Once that is done simply type:

`make`

The app will not work using go's standard tools without ciparser and make. This behavior eventually will be dropped.

Ciparser uses [musl-gcc](http://www.musl-libc.org/), [upx](https://upx.github.io/) and [gometalinter](https://github.com/alecthomas/gometalinter) as default. We statically compile libraries into the executable. You can switch this behavior off in the cell-ci.yaml. See ciparser for more information on compile options.

### Tested on

- Windows 7 --> Check
- Ubuntu 16.04 --> Check

Server Running

- [Sernet Samba 4.2](https://portal.enterprisesamba.com/) --> Check


## Using attitude

Attitude is a command line app and is started in your terminal/cmd.

for more information see

    ./attitude --help

## Todo

1. Right now the admin password is compiled into the app and is displayed in debug mode. We kind of like this *configless* non-interactive execution. But this is very specific to our internal needs and needs external security measures. In future an non-/interactive User/Password request should be done similar to i.e. mysqlclient.
2. Right now you can easily break your directory. This app should prevent you from doing this.
3. Ldaps. TLS encrypted communication
4. Add more attributes
5. Enhance documentation. List attributes with M$ reference

## Thanks to
#### ldap library
Copyright (C) 2012 The Go Authors. See https://github.com/nmcclain/ldap/blob/master/LICENSE
This seems to be a fork from https://github.com/mmitton/ldap/blob/master/LICENSE

#### logrus
Copyright (c) 2014 Simon Eskildsen https://github.com/Sirupsen/logrus/blob/master/LICENSE

#### cli
Copyright (c) 2016 Jeremy Saenz & Contributors https://github.com/urfave/cli/blob/master/LICENSE


## Copyright
Copyright (C) 2016 Celluloid VFX and Johannes Amorosa

##License
GPL Version 2 see attached [License file](./LICENSE)

