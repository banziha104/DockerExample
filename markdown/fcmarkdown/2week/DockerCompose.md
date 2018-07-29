# Docker Compose

> 여러개 컨테이너로 구성된 어플리케이션을 만들고 관리할 수 있게 해주는 도구

- 수동설치
    - sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
    - chmod +x /usr/local/bin/docker-compose

- 명령어
    - docker-compose up -d --no-recreate : ㅅ

- docker-compose.yml 
- 어플리케이션을 만드는 서비스를 정의하는 yaml 파일

![Compose1](https://github.com/banziha104/DockerExample/blob/master/img/2week/compose1.png)
![Compose2](https://github.com/banziha104/DockerExample/blob/master/img/2week/compose2.png)

- mysql.yml 정의


```yaml
mydb:
  image: mysql:5.6
  enviroment:
    - MYSQL_ROOT_PASSWORD=test1234
  ports:
    - 3306:3306
  volumes:
    - /home/docker/my.cnf:/etc/mysql/conf.d/my.cnf
    - /db_master:/var/lib/mysql
```

- docker-compose -f mysql.yml up -d : mysql.yml 을 기반으로 컨테이너 실행

- docker-compose.yml 정의
- docker-compose up -d : docker-compose.yml을 실행
- docker-compose scale mywas=3 : mywas만 세대 가동
 
 
 
## 버전 2 사용법

- restart : 죽었을때의 옵션
- depend_on : 먼저 실행될것이 있으면 정의해줌


```yaml
version: '2'
services:
  db:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: wordpress
      MYSQL_DATABASE: wordpress
      MYSQL_USER: wordpress
      MYSQL_PASSWORD: wordpress

  wordpress:
    depends_on:
      - db
    image: wordpress:latest
    links:
      - db
    ports:
      - "8000:80"
    restart: always
    environment:
      WORDPRESS_DB_HOST: db:3306
      WORDPRESS_DB_PASSWORD: wordpress
```

### 버전 3 