# go-rate-limiter
## General info
HI there 👋, this is a simple rate limiter, implemented using sliding window algorithm.

## Technologies
Project is created with:
* Go
* Redis
* Docker

## Setup
To run this project, make sure you have installed Docker, and these ports are not busy
* 6379 - for Redis
* 1111 - for localhost:1111/limit/:id

install it locally using:

```
$ clone repo
$ cd go-rate-limiter.git
$ docker-compose up -d --build
```
## API
Simply request localhost:1111/limit/:id

:id represents userId, default limit is 3, if you request with same id 3 times in one minute, it will reject request.
