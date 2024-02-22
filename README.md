# Where is love

`Where is love` is revolutionary `api` which allows you to find a friend to meet and chat and hopefully get in `love`.

## No no no

- [ ] Password is not encrypted
- [ ] Logging format and definition

## Architecture decisions

- Split the application based on the bounded context like `user`, `discovery` and `swipe`
- Because this is test application keeping separating of contexts
- Logging
- Monitoring
- Because there is not much logic keeping it in the `endpoint` in real implementation that would be in some service
- `postgres` being used as search engine which is not optimal

## Future improvements

- [ ] `test` directory is being used for integration tests and application has to be up and running
- [ ] migrate authentication to `jwt` token
- [ ] think about further splitting `bounded context`s to enable better module separation
- [ ] more documentation
- [ ] create `openapi` specification and generate client to improve integration test readability instead of using base types

## Prerequisite

1. [Docker](https://docs.docker.com/engine/install/)
2. [go](https://go.dev/doc/install)

You have to use `golang` +1.22.

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
