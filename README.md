# polls-api

[![Build Status](https://travis-ci.org/drklee3/polls-api.svg?branch=master)](https://travis-ci.org/drklee3/polls-api)

Simple REST API to make polls.

# Running

```bash
go get github.com/drklee3/polls-api
go build
./polls-api
```

# Configuration

Configuration options are read from environment variables, optionally loaded from an `.env` file.  The variables used are given below with example values. If `INTERFACE`, `PORT`, or `DB_LOG` are not provided, they will use `127.0.0.1`, `3000`, `false` respecitvely.

```shell
DB_USERNAME=admin
DB_PASSWORD=Hunter2
DB_NAME=polls
DB_HOST=127.0.0.1

# optional variables
INTERFACE=127.0.0.1  # web server interface
PORT=3000            # web server port
DB_LOG=1             # gorm logging
```

# Endpoints

## /polls

* `GET` Get all polls
* `POST` Create a new poll

## /polls/{id:[0-9]+}

* `GET` Get a single poll
* `PUT` Update a poll
* `DELETE` Delete a poll

## /polls/{id:[0-9]+}/vote

* `POST` Vote on a poll

## /polls/{id:[0-9]+}/archive

* `PUT` Archive a poll (disables votes)
* `DELETE` Restores a poll (re-enables votes)
