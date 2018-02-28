# 리눅스에 설치

1. 자동 설치 스크립트 : 리눅스 배포판 종류를 자동으로 인식하여 docker 패키지를 설치해주는 스크립트를 제공
    - sudo wget -q0- https://get.docker.com/ | sh // docker 설치
    - sudo docker rm `sudo docker ps -aq`   // hello-world 이미지 관련 삭제
    - sudo docker rmi hello-world  // hello-world 이미지 관련 삭제

2. 우분투 설치 : 우분투에서 패키지로 직접 설치하는 방법
    - sudo apt-get update   // 패키지 매니저 업데이트
    - sudo apt-get install docker.io    // 도커 설치
    - sudo ln -sf /usr/bin/docker.io /usr/local/bin/docker // 실행파일을 링크로 해서 사용