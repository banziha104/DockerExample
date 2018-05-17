# Go Echo Server 배포

- nginx 설치

```bash

docker run -p 8080:80 nginx

```


- mysql 설치

```bash

docker \
  run \
  --detach \
  --env MYSQL_ROOT_PASSWORD=1234\
  --env MYSQL_USER=${MYSQL_USER} \
  --env MYSQL_PASSWORD=${MYSQL_PASSWORD} \
  --env MYSQL_DATABASE=${MYSQL_DATABASE} \
  --name ${MYSQL_CONTAINER_NAME} \
  --publish 3306:3306 \
  mysql;

```

```bash
docker \
  run \
  --detach \
  --env MYSQL_ROOT_PASSWORD=1234\
  -p 3306:3306 \
  mysql;
```

- mongodb 설치

```bash
docker run --name mongo -p 27017:27017 -d mongo
```
