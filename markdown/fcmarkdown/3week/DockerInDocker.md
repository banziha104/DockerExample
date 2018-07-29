# 도커안에 도커

> Docker 컨테이너가 Docker 명령어를 실행할 수 있게함

- 준비사항
    - Docker : 소켓 파일에 대한 접근권한
    - Docker : 클라이언트, 실행권한
    
```bash
docker run -d --name jenkins -p 8080:8080 -v /var/run/docker/sock:/var/run/docker.sock k16wire/docker-jenkins  # 이미지 빌드용 jenkins 컨테이너 실행
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword

```