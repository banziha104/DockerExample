# Docker

> 반 가상화보다 좀 더 경량화된 방식으로, docker이미지에 서버 운영을 위한 프로그램과 라이브러리만 격리해서 설치 가능. OS 자원은 공유

- 베이스 이미지 : 유저랜드만 설치된 파일
- Docker 이미지 : 베이스 이미지에 필요한 프로그램과 라이브러리, 소스를 설치한 뒤 파일을 하나로 만드는 것을 말함
- 레이어 : 도커는 베이스 이미지와 비교 했을 때, 다른 부분
- 기본적으로 도커는 git처럼 버전관리가 가능하며, 기존과 다른 레이어만 이미지 처리하고, 베이스 이미지에 더해 사용하는 방식

<br>

# 리눅스에 설치

1. 자동 설치 스크립트 : 리눅스 배포판 종류를 자동으로 인식하여 docker 패키지를 설치해주는 스크립트를 제공
    - sudo wget -q0- https://get.docker.com/ | sh // docker 설치
    - sudo docker rm `sudo docker ps -aq`   // hello-world 이미지 관련 삭제
    - sudo docker rmi hello-world  // hello-world 이미지 관련 삭제

2. 우분투 설치 : 우분투에서 패키지로 직접 설치하는 방법
    - sudo apt-get update   // 패키지 매니저 업데이트
    - sudo apt-get install docker.io    // 도커 설치
    - sudo ln -sf /usr/bin/docker.io /usr/local/bin/docker // 실행파일을 링크로 해서 사용

# 도커기본 사용법

- 검색
    - sudo docker search nginx : search 명령으로 특 이미지 검색

- 이미지 받기
    - sudo docker pull ubuntu:latest : pull 명령으로 이미지 받기
    - docker images : 받은 이미지 목록 출력
    - docker run -i -t --name hello ubuntu /bin/bash : run 명령어로 컨테이너 생성
    - docker run <옵션> <이미지 이름> <실행할 파일>
    - docker ps -a : 컨테이너 목록 확인
        - a : 정지된 컨테이너 까지 모두 출력
    - docker start <컨테이너 이름>  : 컨테이너 실행
    - docker restart <컨테이너 이름> : 컨테이너 재시작
    - docker attach <컨테이너 이름> : 컨테이너에 접속
    - docker exec <컨테이너 이름>  <명령> <매개변수> : 외부에서 컨테이너 안의 명령 실행
    - docker stop <컨테이너> : 컨테이너 정지
    - docker rm <켄터이너 이름> : 컨테이너 삭제 
    - docker rmi <이미지 이름>:<태그> : 이미지 삭제
    

---

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

# Bash 익히기

- \> : 출력 다이렉션, 덮어씀
- \< : 입력 다이렉션
- \>> : 명령 실행의 표춘 출력을 파일에 추가, 뒷부분에 추가함 
- 2> : 표준 에러를 파일로 저장
- 2>> : 표준 에러를 파일에 추가
- &> : 표준 출력과 표준 에러를 모두 파일로 저장함

- 1>&2 : 표준 출력을 표준 에러로 보냄.
- | : 명령 실행의 표준 출력을 다른 명령의 표준 입력으로 보냄
- $ : 변수사용시에만 씀
- $() : 명령 실행 결과를 변수화함. 명령 실행 결과를 변수에 저장하거나 다른 명령의 매개변수로 넘겨줄 때 사용
- `` : 명령 실행결과를 변수화함 $()와 같음
- && : 한줄에서 여러 개의 명령어를 처리함
- '' : 문자열, 안에 들어 있는 변수는 처리되지 않고 변수명 그대로 사용됌
- "" : 문자열, 안에 있는 변수는 사용가능함
- ${} : 변수 치환 문자열 안에서 변수를 출려갈 때 주로 사용됌
- \ : 한 줄로된 명령을 여러 줄로 표현할 때 사용됌
- {1..10} : 연속된 숫자를 표현함
- {문자열1, 문자열2} : 배열과 같이 문자열으 여러개 지정하여 명령 실행 횟수를 줄임
- <<< : 문자열을 명령의 표준입력으로 보냄
- <<EOF : 여러줄의 명령을 명령의 표준 입력으로 보냄
- export : 설정한 값을 환경 변수로 만듬
- printf : 지정한 형식대로 값을 출력함. 파이프와 연동하여 명령에 값을 입력하는 효과를 낼 수 있음
- sed : 택스트 파일에서 문자열을 변경
- \# : 주석


### 제어문

- if : 조건문
    1. 숫자 비교
        - -eq : 같다
        - -ne : 같지 않다
        - -gt : 초과
        - -ge : 이상
        - -lt : 미만
        - -le : 이하
    2. 문자열 비교 
        - =,== : 같다
        - != : 같지않다
        - -z : 문자열이 null 일때
        - -n : 문자열이 null이 아닐때
        
- for : 반복문
- while : 반복문


```bash

for i in {1..100}
do
if [ $i -le 50 ]; then
    echo "$i is less Than 50"
fi
done

```
    
