# SLAPCHOP [![Build Status](https://travis-ci.org/billyninja/slapchop.svg?branch=master)](https://travis-ci.org/billyninja/slapchop) [![GoReport](https://goreportcard.com/badge/billyninja/slapchop)](http://goreportcard.com/report/billyninja/slapchop) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
---

[![Slapchop video!](http://img.youtube.com/vi/rUbWjIKxrrs/0.jpg)](http://www.youtube.com/watch?v=rUbWjIKxrrs)

----
## Basic Idea

Post images to this services, and it will slice and dice it into serveral image tiles.


## Options/Flags
- -host Default: "localhost, The host address which it will be visible
- -puzzler Default: "" Puzzler Service remote url
- -port Default: "3001 HTTP port number
- -tile Default: 64Tile Size in pixels
- -size Default: int64(1024*1024*5) Max upload file size in BYTES
- -dir Default: "/tmp/slapchop/upload"
- -template Default: "notsetted.html" Absolute path to the preview.html template file

`cd $GOPATH/src/github.com/billyninja/slapchop`

`go build .`

`./slapchop -port=9000 -puzzler=username:pwd@localhosthost:8001`
>Notice that I'm changing the default commandline args/flags on purpose


----
## Requirements for dev and building
`go get github.com/julienschmidt/httprouter`

`go get github.com/go-resty/resty`

`go get github.com/hoisie/mustache`

----
## Requirements for production
`Just the properly configured binary ;)`

----
## Running the tests
`go test -v -cover ./...`


---
## Actions

- **GET** `/` -> Retrieves a list of the uploaded slapchops.

- **POST** `/chopit/$USERNAME` -> Upload image to the service storage, split into serveral tiles, and retrieves the hrefs to the user.

- **GET** `/chopit/$USERNAME/$CHOP_ID` -> Retrieves info and href on the given $CHOP_ID

- **DELETE** `/chopit/$USERNAME/$CHOP_ID` -> Deletes this entry and its files

- **GET** `/tiled/$USERNAME/$CHOP_ID` -> Retrieves a html preview of the tiled version of the chopped image

- **GET** `/random/$USERNAME/$CHOP_ID` -> Retrieves a html preview of the shuffled version of the chopped image
