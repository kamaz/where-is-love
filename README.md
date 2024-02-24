# Where is love

`Where is love` is revolutionary `api` which allows you to find a friend to meet and chat and hopefully get in `love`.

## 💔 fix

- [ ] Switch token usage from `base64` to `JWT`.
- [ ] Encrypt passwords.
- [ ] Standardise logging format and definition.
- [ ] Unify error handling.
- [ ] Create a configuration file for easy application setup.
- [ ] Separate DB and search functionality to enhance query efficiency.
- [ ] Enable logging.
- [ ] Distinguish between `test` and `integration test`. Currently, the application must be running to execute `go test ./...`.
- [ ] Expand documentation.
- [ ] Improve module separation.
- [ ] Implement auditing and monitoring capabilities.
- [ ] Set up CI/CD.
- [ ] Create an `openapi` specification and generate client to enhance integration test readability, instead of using base types.
- [ ] Hide internal errors from user

## 🏗️ Architecture decisions

- Split the application based on bounded contexts such as `user` and `match`.
- The `user` context allows creating a user and logging in (ideally, this would support OIDC).
- The `match` context enables discovering matches and swiping on people we like.
- Minimal business logic within `endpoints`.
- Use `Postgres` as a single DB, which may not be optimal, especially when adding search functionality. A search engine might be a better solution.
- Distribute data across bounded contexts.

## Prerequisite

1. [Docker](https://docs.docker.com/engine/install/)
2. [go](https://go.dev/doc/install) version 1.22 or higher.

## How to start app 

First, set up the database by executing the following command:

```shell
docker compose --profile tools run db-up
```

Then, start the application with:

```shell
docker compose up
```

After the migration is applied and the application is running, run the full set of tests to ensure everything works correctly:

Note: Since users and data are generated randomly, an integration test may fail on the initial run (low probability, but possible). If this happens, simply rerun the test.

```shell
go test -count 1 ./...
```

## Other commands

To remove all migration scripts, run the following command:

```shell
docker compose --profile tools run db-down
```

To log in to the database:

```shell
docker compose exec db psql -U postgres -d postgres -W
```

To clean up and remove all images:

```shell
docker compose down --remove-orphans --rmi all -v
```
