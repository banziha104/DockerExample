# 도커머신

> 가상서버에 도커엔진을 설치할수 있게 도와주는 도구, 가상머신 필요!!



- 설치


```bash
curl -L https://github.com/docker/machine/releases/download/v0.13.0/docker-machine-`uname -s`-`uname -m` >/tmp/docker-machine &&
chmod +x /tmp/docker-machine &&
sudo cp /tmp/docker-machine /usr/local/bin/docker-machine
```

- docker-machine create --driver virtualbox my-dev : my-dev virtualbox에 my-dev를 생성
- config : 설정하기
- docker-machine env my-dev : my-env의 ip,위치, 이름등을 알려줌 
- eval $(docker-machine env my-dev) : 현재 쉘에 서버 연결정보를 설정 
- docker-machine scp nginx.conf my-dev:~/dev/config : 로컬파일을 해당머신에 위치에 복사함
- docker-machine inspect my-dev : my-dev를 구현함
- docker-machine ip my-dev : ip 조회
- docker-machine kill my-dev : 삭제
- docker-machine restart my-dev : 재시작
- docker-machine provision my-dev : 최신 이미지로 꺠끗하게 지움
- docker-machine ssh my-dev : my-dev로 ssh 연결함
- start : 
- stop : 
- status : 
- docker-machine upgrade : 최신으로 업그레이드
- docker-machine url : 해당 머신에 접속할 수 있는 url 을볼 수 있음