# Docker 호스트 모니터링

### htop 

> 호스트를 모니터링함

```bash
docker run -it --rm --pid host tehbilly/htop # rm 나오면서 지워짐 q를 누르면 지워짐 
docker run -it --rm --pid=container:myweb tehbilly/htop # myweb 컨테이너를 모니터링함 
```

### sysdig

> 호스트를 모니터링함2 다소복잡함, 리소스 모니터링을 다해야함, 무거운대신 기능이많음

```bash
docker run -it --rm --name=sysdig --privileged=true \
--volume=/var/run/docker.sock:/host/var/run/docker.sock \
--volume=/dev:/host/dev \
--volume=/proc:/host/proc:ro \
--volume=/boot:/host/boot:ro \
--volume=/lib/modules:/host/lib/modules:ro \
--volume=/usr:/host/usr:ro \
sysdig/sysdig

```


### Glances

> 파이썬으로 만든 시스템 모니터링 도구 (*), CLI, API,Web인터페이스 지원

```bash

# glance 실행
docker run -it --pid=host \
-v /var/run/docker.sock:/var/run/docker.sock:ro \  # readonly
--name grances \
nicolargo/glances

```

```bash
docker run -d -p 61208:61208 -p 61209:61209 \
-e "GLANCES_OPT=-w" --pid=host -v /var/run/docker.sock:/var/run/docker.sock:ro \
--name glances nicolargo/glances

```

# Docker 컨테이너 모니터링

- 컨테이너 로깅 기준
    - 로그를 이벤트 스트림으로 취급
    - 로그파일로 관리하지 말고 stdout으로 출력
    - 로그파일은 앱이 아닌 실행환경으로 관리한다
   
- ELK를 이용한 로그분석 환경 구축 
    - Elastic Search : 아파치 루신 기반의 검색엔진 서버
    - Logstash : 로그와 같은 데이터를 실시간으로 수집해주는 엔진
    - Kibina : 엘라스틱 서치 결과를 분석하고 표현해주는 플랫폼
    