# docker-controll

### Init project

```bash
docker-compose up
```

### Endpoints

```
GET `/containers` - return list docker containers
GET `/images` - return list docker images
GET `/containers/stop` - stop and return all docker containers
POST `/images/pull` - pull docker image by ImageID
POST `/container/run` - run container by ContainerID
POST `/container/log` - show container log by ContainerID
```
