# Docker

> 반 가상화보다 좀 더 경량화된 방식으로, docker이미지에 서버 운영을 위한 프로그램과 라이브러리만 격리해서 설치 가능. OS 자원은 공유

- 베이스 이미지 : 유저랜드만 설치된 파일
- Docker 이미지 : 베이스 이미지에 필요한 프로그램과 라이브러리, 소스를 설치한 뒤 파일을 하나로 만드는 것을 말함
- 레이어 : 도커는 베이스 이미지와 비교 했을 때, 다른 부분
- 기본적으로 도커는 git처럼 버전관리가 가능하며, 기존과 다른 레이어만 이미지 처리하고, 베이스 이미지에 더해 사용하는 방식

---

# 쿠버네티스

https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/01_Kubernetes.md

- ### [설치 및 세팅](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/01_Kubernetes.md)

- ### [Pod](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/02_Pod.md)
  
- ### [Label & Selector](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/03_Label.md)
  
- ### [Replica](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/04_Replica.md)
  
- ### [Deployment](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/05_Deployment.md)
  
- ### [Namespaces](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/06_Namespaces.md)
  
- ### [Services](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/07_Services.md)
  
- ### [Ingress](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/08_Ingress.md)
  
- ### [Network](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/09_Network.md)
  
- ### [Volume](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/10_Volume.md)
  
- ### [Config](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/11_Config.md)
  
- ### [Scheduling](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/12_Scheduling.md)
  


<br>

# 도커

- [DockerBasic](https://github.com/banziha104/DockerExample/tree/master/markdown/fcmarkdown)

- [Docker Machine](https://github.com/banziha104/DockerExample/tree/master/markdown/fcmarkdown)



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

- 전체 컨테이너 종료 : docker rm -f $(docker ps -a -q)
- 전체 이미지 제거 : docker rmi $(docker images -q)


<br>