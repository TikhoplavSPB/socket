# Socket

Simple TCP broadcasting server

## Deployment

### Dev deployment

Prerequisites:
- Docker
- Docker-compose

```
docker-compose up
```

Dev build will be available on port 8080.

### Prod deployment

Prerequisites:
- Docker

```
docker service create --name tikhoplav/socket -p 8080:8080 tikhoplav/socket
```