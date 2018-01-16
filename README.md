[![Build Status](https://travis-ci.org/harikiriboy/nifcloud-ssh-config.svg?branch=master)](https://travis-ci.org/harikiriboy/nifcloud-ssh-config)
[![Go Doc](https://godoc.org/github.com/harikiriboy/nifcloud-ssh-config?status.svg)](http://godoc.org/github.com/harikiriboy/nifcloud-ssh-config)
[![Go Report](https://goreportcard.com/badge/github.com/harikiriboy/nifcloud-ssh-config)](https://goreportcard.com/report/github.com/harikiriboy/nifcloud-ssh-config)
[![Coverage Status](https://coveralls.io/repos/github/harikiriboy/nifcloud-ssh-config/badge.svg?branch=master)](https://coveralls.io/github/harikiriboy/nifcloud-ssh-config?branch=master)

## About

A very simple tool that generates SSH config file using NIFCLOUD Computing API.

## Install

```
$ go get github.com/harikiriboy/nifcloud-ssh-config
```

or download binary [here](https://github.com/harikiriboy/nifcloud-ssh-config/releases/download/v1.0.0/nifcloud-ssh-config)

## Usage

```
Usage:
   [flags]

Flags:
  -h, --help                              help for this command
      --access-key string                 NIFCLOUD API ACCESS KEY (default NIFCLOUD_ACCESS_KEY_ID environment variable
      --exclude-instance-id stringSlice   List of Exclude instanceID
      --keydir string                     Location of private keys (default "~/.ssh")
      --no-identities-only                Do not include IdentitiesOnly=yes in ssh config; may cause connection refused if using ssh-agent
      --port string                       SSH port (default "22")
      --prefix string                     Specify a prefix to prepend to all host names
      --private                           Use private IP addresses (public are used by default)
      --region stringSlice                List of NIFCLOUD Regions (default all region)
      --secret-key string                 NIFCLOUD API SECRET KEY (default NIFCLOUD_SECRET_ACCESS_KEY environment variable
      --ssh-key-name string               SSH key name (default use keyName get from API response)
      --strict-hostkey-checking           Do not include StrictHostKeyChecking=no in ssh config
      --user string                       SSH username (default "root")
```


### Example

```
$ nifcloud-ssh-config > ~/.ssh/config
$ cat ~/.ssh/config
Host api
  HostName 0.0.0.0
  User root
  Port 22
  IdentityFile ~/.ssh/nifcloud.pem
  IdentitiesOnly yes
  StrictHostKeyChecking no

Host web
  HostName 0.0.0.0
  User root
  Port 22
  IdentityFile ~/.ssh/nifcloud.pem
  IdentitiesOnly yes
  StrictHostKeyChecking no

Host db
  HostName 0.0.0.0
  User root
  Port 22
  IdentityFile ~/.ssh/nifcloud.pem
  IdentitiesOnly yes
  StrictHostKeyChecking no
```
