# 리눅스에 설치

1. 자동 설치 스크립트 : 리눅스 배포판 종류를 자동으로 인식하여 docker 패키지를 설치해주는 스크립트를 제공
    - sudo wget -q0- https://get.docker.com/ | sh // docker 설치
    - sudo docker rm `sudo docker ps -aq`   // hello-world 이미지 관련 삭제
    - sudo docker rmi hello-world  // hello-world 이미지 관련 삭제

2. 우분투 설치 : 우분투에서 패키지로 직접 설치하는 방법

    - sudo apt-get update   // 패키지 매니저 업데이트
    - sudo apt-get install docker-engine 도커 설치
    - sudo apt-get upgrade docker-engin 
    - sudo apt-get install docker.io    // 도커 설치
    - sudo ln -sf /usr/bin/docker.io /usr/local/bin/docker // 실행파일을 링크로 해서 사용
    
3. EC2 에서 Docker 사용
    - Configure Instance -> Advanced Details 클릭
    - 아래 메시지 입력
    
    
```shell

#cloud-config

packages:
 - curl
  
runcmd:
 - [ sh, -c, "curl https://get.docker.com/ | sh" ]
 - [ sh, -c, "usermod -aG docker ubuntu" ]

```
    
4. Elastic Beanstalk에서 Docker 사용

- 새어플리케이션을 만듬
- Environment tier : Web Server (인터넷에서 접속할 수 있는 웹 서버)로 만듬 (Worker : 백그라운드 작업 환경)
- Predefined configuration : 개발 언어 또는 플랫폼 -> Docker 선택
- Environment Type : 기본 
    
5. Docker hub 공개 저장소 이미지 사용

- Dockerrun.aws.json 파일을 통해 공개 저장소에 저장된 이미지를 그대로 사용가능함
- Nginx 예제

```javascript
var c = 
{
  "AWSEBDockerrunVersion" : "1",
  "Image" : {
    "Name" : "nginx:latest",
    "Update" : "true"
  },
  "Ports" : [
    {
      "ContainerPort" : "80"
    }
  ],
  "Volumes" :[
    {
      "HostDirectory" : "/var/www",
      "ContainerDirectory" : "/var/www"
    }
  ],
  "Logging" : "/var/log/nginx"
}
```

6. Docker Hub 개인 저장소 이미지 사용하기

- sudo docker login 으로 로그인
- 
