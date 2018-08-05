# Docker Swarm

> 도커 호스트 클러스터를 구성하고, 클러스터에 컨테이너를 배치해주는 도구

- 표준 Docker API
- 스케쥴링 지원
- 디스커버리 지원 ( 자기들이 ) : Consul, Etcd, ZooKeeper 
- Swarm Manger : 클러스터에 컨테이너를 배치해주는 도구

### Swarm mode

> 도커엔진에 클러스터 관리가 통합된 모드

- 멀티 호스트 네트워크 : 분산환경에서 여러개 노드를 하나의 네트워크로 묶은것
- 서비스 디스커버리 : 멀티 호스트 환경에서 실행된 컨테이너 정보를 제공
- 로드 밸런싱 : 대용량 트래픽을 분산해줌
- 롤링 업데이트 : 새로운 이미지를 순차적으로 업데이트 해 주는 것
- Health Check : 서비스가 정상 상태인지 확인
- 스케일 아웃 : 컨테이너르 원하는 상태로 관리 , Desired State Reconcilation
- 로깅
- 모니터링
- HA

### Swarm mode 중요 개념

- Swarm 클러스터
    - Swarm을 이용해 구축한 클러스터
    - aka Swarm
    
- Node : Swarm 클러스터에 참여하는 도커 엔진 인스턴스
    - Manager Node 
        - swarm 클러스트 상태를 관리하는 노드, 오케스트레이션
        - swarm 메니저 실행
        - service 정의 , Task 할당
    - Worker Node 
        - 매니저 노드의 명령을 받아서 컨테이너 관련 작업을 수행
        - Task 처리
    - Service : 배포의 단위
        - 한개의 서비스는 여러개의 태스크를 가짐
        - 보통 1개의 이미지를 이용해 동일한 타입의 컨테이너를 여러개 실행한다.
        - global services
            - 클러스터에 있는 모든 노드에서 실행되는 서비스
        - replicated services
            - 스케일 아웃을 위해 특정 노드에서 실행하는 서비스
    - Task 
        - 컨테이너 배포 단위, 도커 켄테이너와 컨테이너에서 처리하는 명령
        - swarm 스케쥴링 단위  
        
### 예제

1. node1,2,3 생성

```bash
docker-machine create -d virtualbox node1
```

2. 노드 확인

```bash
docker-machine ls
```

3. 매니저 노드 생성

```bash
eval $(docker-machine env node1)
docker swarm init --advertise-addr 192.168.99.101:2376 # swarm 모드로 초기화
``` 

4. 워커 노드 추가
        
```bash
eval $(docker-machine env node2)
docker swarm join \
--token \

\ 192.168.99.101:2376

```
        
### Swarm Mode 명령어

- swarm init 
- swarm join
- service create
- service inspect
- service ls
- service rm
- service scale
- service ps
- service update 


    
### 서비스 스케일아웃
    
- 컨테이너를 5개로 늘림

```bash

docker $(docker-machine config node1) \ 
service scale myweb=5 

```

### 서비스 롤링업데이트

1. myweb 서비스 이미지를 1.11버전으로 업데이트

```bash
docker $(docker-machine config node1) \ 
service update \
--image nginx:1.11 myweb

```

2. 서비스 확인

```bash
docker $(docker-machine config node1) service ls myweb 
```



# Ingress 네트워크

- Swarm 클러스터 외부에 서비스 포트 공개를 지원
- 모든 노드는 ingress routing mesh에 들어감
- 필수 포트
    - 7946 : 디스커버리 포트
    - 4789 : ingress network 포트 
    