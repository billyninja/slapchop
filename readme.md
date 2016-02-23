# SLAPCHOP [![Build Status](https://travis-ci.org/billyninja/slapchop.svg?branch=master)](https://travis-ci.org/billyninja/slapchop) [![GoReport](https://goreportcard.com/badge/billyninja/slapchop)](http://goreportcard.com/report/billyninja/slapchop) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
---

[![Slapchop video!](http://img.youtube.com/vi/rUbWjIKxrrs/0.jpg)](http://www.youtube.com/watch?v=rUbWjIKxrrs)

----
## Basic Idea

Post images to this services, and it will slice and dice it into serveral image tiles.


## Options/Flags

- -port Port number for the process **default:** "3001"
- -puzzler Host/Url for the puzzle service **default:** "localhost:8000"

`cd $GOPATH/src/github.com/billyninja/slapchop`

`go build .`

`./slapchop -port=9000 -puzzler=localhosthost:8001`
>Notice that I'm changing the default commandline args/flags on purpose


----
## Requirements for dev and building
`go get github.com/julienschmidt/httprouter`

`go get github.com/go-resty/resty`

----
## Requirements for production
`Just the properly configured binary ;)`

----
## Running the tests
`go test -v -cover ./...`


---
## Actions

- **GET** `/` -> Retrieves a list of the uploaded slapchops.


- **POST** `/chopit` -> Upload image to the service storage, split into serveral tiles, and retrieves the hrefs to the user.

- **GET** `/chopit/$CHOP_ID` -> Retrieves info and href on the given $CHOP_ID

- **DELETE** `/chopit/$CHOP_ID` -> Deletes this entry and its files

- **GET** `/random/$CHOP_ID` -> Retrieves a randomized list of the slapchop tiles

- **GET** `/random/$CHOP_ID?preview=1` -> Retrieves info and href on the given $CHOP_ID
