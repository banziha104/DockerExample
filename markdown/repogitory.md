# Docker Hub 및 개인 저장소

### 로컬

- docker pull registry:latest : 이미지를 가져옮
- docker tag hello:0.1 localhost:5000/hello:0.1 : 태그생성, 이미지를 올릴 때는 태그를 항상 먼저 생성 <이미지 이름>:<태그> <Docker 저장소 URL>/<이미지 이름>:<태그>
- docker push localhost:5000/hello:0.1 : 이미지를 올림

### 기본 인증

> Docker 레지스트리에는 로그인 기능이 없음으로, NGINX의 기본 인증 기능을 사용해야함

### 도커 컨테이너 연결하기 

- 별칭 : db와 같이 설정한 이름
- 주소 : 192.168.0.1 과 같은 주소 값

```bash
// --link <컨테이너 이름>:<별칭> 을 이용해 연결
sudo docker run --name web -d -p 80:80 --link db:db nginx

```

### 다른 서버의 Docker 컨테이너에 연결하기 (앰배서더 컨테이너)

- docker 파일 생성

```dockerfile
CMD env | grep _TCP = |\
    sed 's/.*_PORT_\([0-9]*\)_TCP=tcp:\/\/\(.*\):\(.*\)/socat\
    TCP4-LISTEN:\1,fork,reuseaddr TCP4: \2:\3 \&/` \
    | sh && top
```

- sudo docker run -d --link redis:redis --name redis_ambassador -p 6379:6379 svendowideit/ambassador
    - -d 옵션으로 컨테이너를 백그라운드로 실행시킴
    - --link redis:redis : redis컨테이너를 redis 별칭으로 연결
    - ---name redis_ambassador로 이름을 지정
    - -p 6379:6379 옵션으로 컨테이너의 6379번 포트와 호스트의 6379 연결 하고 외부에 노출