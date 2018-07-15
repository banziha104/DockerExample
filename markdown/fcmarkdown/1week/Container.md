# 도커 컨테이너

- 이미지를 실행한상태
- 읽기쓰기가 가능한 파일시스템
- 실행된 독립 어플리케이

```bash
# nginx.conf를 변
docker run --name myweb2 -v /home/docker/nginx.conf:/etc/nginx/nginx/conf -d -P nginx

```
