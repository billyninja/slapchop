# SLAPCHOP

![Slapchop video!](http://img.youtube.com/vi/rUbWjIKxrrs/0.jpg)](http://www.youtube.com/watch?v=rUbWjIKxrrs)

----
## Basic Idea

Post images to this services, and it will slice and dice it into serveral image tiles.

----
## Requirements for dev and building
`go get github.com/julienschmidt/httprouter`

----
## Requirements for production
`Just the properly configured binary ;)`


---
## Actions

- **GET** `/` -> Retrieves a list of the uploaded slapchops.


- **POST** `/chopit` -> Upload image to the service storage, split into serveral tiles, and retrieves the hrefs to the user.

- **GET** `/chopit/$CHOP_ID` -> Retrieves info and href on the given $CHOP_ID

- **DELETE** `/chopit/$CHOP_ID` -> Deletes this entry and its files

- **GET** `/random/$CHOP_ID` -> Retrieves a randomized list of the slapchop tiles

- **GET** `/random/$CHOP_ID?preview=1` -> Retrieves info and href on the given $CHOP_ID