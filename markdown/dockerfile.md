
# 도커파일 작성법 및 사용법

- docker 파일 작성 

```dockerfile

# 어떤 이미지를 기반으로 할건지 설정
FROM ubuntu:14.04

# 메인터이너 정보
MAINTAINER Foo Bar <foo@bar.com>

# 셸 스크립트 혹은 명령을 실행
RUN apt-get update
RUN apt-get install -y nginx
RUN echo "\ndaemon off" >> /etc/nginx/nginx.conf
RUN chown -R www-data:www-data /var/lib/nginx

# 호스트와 공유할 디랙토리 명령
VOLUME ["/data","/etc/nginx/site-enabled","/var/log/nginx"]

# CMD 에서 설정한 파일이 실행될 디렉터리
WORKDIR /etc/nginx

# 컨테이너가 시작외었을 떄 실행할 실행 파일 또는 셸스크립트
CMD ["nginx"]

# Host와 연결할 포트 번호
EXPOSE 80
EXPOSE 443
```
    
- build 명령으로 이미지 생성하기
    - docker build <옵션> <Dockerfile 경로>
    
<br>