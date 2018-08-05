# Docker Network

> 도커가 시작될때 호스트 머신에 docker0라고 부르는 가상 인터페이스를 생성함, docker0에 사설 IP가 랜덤하게 배정된다.

- bridge : docker0
- overlay : 멀티 호스트간에 연결해주는 네트워크

- docker network create

```bash
# 생성
docker network create red
docker network create blue

# 사용 
docker run -itd --net=red --name ubuntu1 ubuntu
docker run -itd --net=blue --name ubuntu3 ubuntu
```

- docker network ls
- docker network inspect

```bash
docker network inspect bridge
```


- docker network rm
- docker network connect

```bash
docker network connect blue ubuntu1
```

- docker network disconnect

### bride network 

> 컨테이너는 동일 호스트내에 위치해야하며, 사용자 정의 bridege 네트워크에 포함된 컨테이너는 컨테이너 이름으로 통신 가능

### overlay network

- 다른 도커 호스트에서 각각 실행중인 컨테이너들이 서로 통신할 수 있게 해준다.
- 도커 엔진은 overlay 네트워크 드라이버를 통해 멀티 호스트 네트워크 지원


- 요건
    - Key-value 스토어 : Consul, Etcd,Zookeeper
    - Key-value 스토어에 연결된 호스트 클러스터
    - 커널 버전 3.16 이상
    - 각 호스트 도커 엔진 설정
    
1. mhl-consul 머신 추가

```bash
docker-machine create -d virtualbox mhl-consul
```

2. consul 컨테이너 추가

```bash
docker $(docker-machine config mhl-consul) run -d -p 8500:8500 -h consul progrium/consul -server -bootstrap 
```

3. mhl-demo1 머신을 Consul 클러스터에 추가

```bash
docker-machine create \ 
-d virtualbox \
--engine-opt="cluster-store=consul://$(docker-machine ip mhl-consul):8500" \
--engine-opt="cluster-advertise=eth1.0" mhl-demo1

```

4. mhl-demo1,2 네트워크 조회

```bash
docker $(docker-machine config mhl-demo1) network ls 
docker $(docker-machine config mhl-demo1) network create -d overlay frontend
docker $(docker-machine config mhl-demo2) network ls 

```

5. mhl-demo1 frontend 네트워크에 nginx 추가

```bash
docker $(docker-machine config mhl-demo1) run -itd --name=web --net=frontend nginx 
docker $(docker-machine config mhl-demo2) run -it --rm --net=frontend busybox wget -q0- http://web\
```