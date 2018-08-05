# 볼륨

> 컨테이너는 언제든 꺠질수 있음으로, 외부 볼륨을 사용해야한다. 

- 컨테이너와 데이터 분리
- 컨테이너간 데이터 공유
- I/O 성능 향상 (도커가 모든 걸 관리하고 있는데 볼륨을 쓰면 덜 관리해서 빨라짐)
- 호스트와 컨테이너간 파일 공유 

### 볼륨관리 유형

- 컨테이너 내부에 저장
- 도커 AUFS에 저장
- 도커 호스트 파일 시스템 볼륨 마운트 (*)
- volume-driver를 이용해 네트워크로 연결된 장치에 저장한다. (* 멀티호스트 상황에서는 반드시 사용)

### 볼륨을 직접 만들기 

- docker volume create : 볼륨 만들기

```bash
docker volume create --name test1
```

```bash
# 사용
docker run -it -v test1:/www/test1 ubuntu
```

- docker volume ls : 볼륨목록

```bash
docker volume ls
```

- docker volume inspect : 볼륨조회 
- docker volume rm : 볼륨삭제



