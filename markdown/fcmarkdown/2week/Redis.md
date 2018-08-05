# 레디스

- Remote Dicationary System
- 메모리 기반의 Key/Value 스토어
- 저장 가능한 데이터 용량은 물리적인 메모리 크기를 넘어설 수 없음
- 서버 재시작을 위해 Disk에 데이터를 저장 

```bash
docker run -d -p 6379:6379 --name fc-redis redis
```

```bash
docker run -d --name fc-redis2  \
> -v /apps/redis:/data \
> redis redis-server --appendonly yes
```

```bash
docker run -it --link fc-redis:redis --rm redis \
> redis-cli -h redis -p 6379
```

### 레디스 이미지빌드 

1. 설정파일 작성

```text

port 6379
timeout 300
databases 1
requirepass password1
maxclients 1000
appendonly no
tcp-keepalive 60
syslog-enabled yes

```

2. 도커파일 작성

```dockerfile
FROM redis:3.0

COPY redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf"]

```

### 레디스 클러스트 구성

```bash
docker run -d -P grokzen/redis-cluster:3.0.6 
```