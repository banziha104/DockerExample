
# 도커파일 작성법 및 사용법

- docker : <명령> <매개 변수>로 작동하며 #은 주석 
- .dockerignore : 컨텍스트에서 파일이나 디렉터리를 제외하고 싶을 때 사용


#### FROM 

> 어떤 이미지를 기반으로 이미지를 생성할지 설정함

```dockerfile

# 항상 설정하고 맨 처음에 와야함, 로컬에 있으면 바로 사용하고 없으면 Docker Hub에서 받아옮
FROM ubuntu 

```

#### MAINTAINER

> 이미지를 생성한 사람의 정보를 설정함. 형식은 자유

```dockerfile

MAINTAINER Lee,Youngjoon <banziha104@gmail.com>

```

#### RUN

> FROM 에서 설정한 이미지 위에서 스크립트 혹은 명령을 실행함.

```dockerfile
RUN apt-get install -y nginx
RUN echo "Hello Docker" > /tmp/hello
RUN culr -sSL https://golang.org/dl/go1.3.1.src.tar.gz | tar -v -C /usr/local -xz
RUN git clone https://githuc.com/docker/docker.git

# 셸 없이 바로 실행
RUN ["apt-get","install","-y","nginx"]
RUN ["/user/local/bin/hell", "--help"]


```

#### CMD

> 컨테이너가 시작외었을 때 스크립트 혹은 명령을 실행함, 정지된 컨테이너를 시작할떄나 컨테이너를 생성하면 실행됌

```dockerfile

CMD touch /home/hello/hello.txt

# 셸없이 바로 실행하기
CMD ["redis-server"]

# 셸없이 바로 실행할 때 매개 변수 설정
CMD ["redis-server","--datadir=/var/lib/mysql"]

# 엔트리포인트 이용

ENTRYPOINT ["echo"]
CMD ["Hello"]


```

#### ENTRYPOINT

> 컨테이너가 시작되었을 떄 스크립트 혹은 명령을 실행함, CMD와 같지만, docker run 명령에서 동작 방식이다름 
> <br> docker run <이미지> <실행할 파일> 형식에서 실행할 파일을 설정하면 CMD는 무시됌

```dockerfile

ENTRYPOINT touch /home/hello/hello.txt

ENTRYPOINT ["/home/hello/hello.sh"]

```

#### EXPOSE 

> 호스트와 연결할 포트 번호를 설정함, docker run의 --expose 와 동일, 외부로 노출되지 않기떄문에
> <br> docker run -p 80 과 같이 포트로 설정해야함

```dockerfile
EXPOSE 80
EXPOSE 443

```

#### ENV
 
> ENV는 환경 변수를 설정함

```dockerfile

ENV GOPATH /go
ENV PATH /go/bin:$PATH

```

#### ADD 

> 파일에 이미지를 추가함

- 복사할 파일 경로에 인터넷에 있는 파일의 URL 설정가능
- 디렉터리도 가능
- 로컬에 있는 압축파일은 풀어서 추가됌
- 인터넷에서 받을 경우 tar 파일이 그래도 추가됌
- 이미지 경로는 항상 절대 경로로 설정해야함

```dockerfile
# 복사할 파일 경로, 이미지에서 파일이 위치할 경로
# 복사할 파일경로는 컨텍스트아래를 기준으로하며, 컨텍스트 바깥의 파일이나 디렉터리등은 불가능함
ADD hello-entrypoint.sh /entrypoint.sh 
```

#### COPY

> COPY는 파일을 이미지에 추가함, ADD완달리 압축을해제하거나 파일 URL도 사용할 수 없음

```dockerfile

COPY hell-entrypoints.sh /entrypoint.sh

```


#### VOLUME 

> 디렉터리의 내용을 컨테이너가 아닌 호스트에 저장하도록 설정함

```dockerfile
# <컨테이너 디렉토리> 
VOLUME /data

# <컨테이너 디렉터리, 컨테이너 디렉터리2>
VOLUME ["/data","/var/log/hello"] 

# 데이터 볼륨을 호스트의 특정 디렉터리와 연결하려면 -v 옵션 실행
RUN sudo docker run -v /root/data:/data example 


```

#### USER 

> 명령을 실행할 사용자 계정을 설정함

```dockerfile

USER root

```

#### WORKDIR

> RUN,CMD,ENTRYPOINT의 명령이 실행될 디렉터리를 설정함

```dockerfile

WORKDIR /var/www

```

#### ONBUILD

> 생성한 이미지를 기반으로 다른 이미지가 생성될 때 명령을 실행, ONBUILD로 최초에 ONBUILD를 사용한 상태에서는 아무 명령도
> <br> 실행하지 않음, 다음 번에 이미지가 FROM으로 사용될 떄 실행할 명령을 예약하는 기능

```dockerfile
ONBUILD RUN touch /hello.txt
```

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