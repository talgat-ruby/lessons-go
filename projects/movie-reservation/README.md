# Movie Reservation

docker compose run

```shell
$ docker compose --env-file=./.env up --build
```

curl for api request

```shell
$ curl 'http://127.0.0.1:4013/movies'
```

migrations

```shell
docker compose --profile tools run --rm migrate create -ext sql -dir ./migrations NAME_OF_MIGRATION_FILE
docker compose --profile tools run --rm migrate {up,down}
```

seeds

```shell
go run ./internal/cli/... seed
```
