# Request ID plugin for Traefik

This traefik plugin set a 'X-Request-Id' header for every incoming HTTP request that
does not have one already.

The generated ids are UUIDv4.

### Dev setup

Docker is used to package the toolchain needed to build, test and manage the plugin.

```bash
# Build and run the developpment container
docker-compose build
docker compose up -d
```

You can start running tests and the linter with the following Make command
```bash
docker-compose exec -it plugin make
```

To run the linter, do:

```bash
docker-compose exec -it plugin make lint
```

To run the unit test, do:

```bash
docker-compose exec -it plugin make test
```
