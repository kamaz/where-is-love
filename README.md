# Where is love

`Where is love` is revolutionary `api` which allows you to find a friend to meet and chat and hopefully get in `love`.

## No no no

- [ ] Password is not encrypted
- [ ] We would use proper jwt token
- [ ] Logging format and definition

## Architecture decisions

- Split the application based on the bounded context like `user`, `discovery` and `swipe`
-

## Prerequisite

1. [Docker](https://docs.docker.com/engine/install/)
2. [go](https://go.dev/doc/install)

## How to start app 

Before starting you have to setup the database to do that execute following:

```shell
docker compose --profile tools run db-up
```

After running the command you can just start application by running:

```shell
docker compose up
```

## Other commands

To remove all the migration script you can just run following command:

```shell
docker compose --profile tools run db-down
```

To zzzvlogin to db and sniff around you can just do following:

```shell
docker compose exec db psql -U postgres -d postgres -W
```