# Docker ci tool stack

> ci 환경을 미리 다만들어놓은 것

![CI](https://github.com/banziha104/DockerExample/blob/master/img/2week/docker-ci.png)


- docker-machine create -d virtualbox docker-ci : docker ci 머신을 생성
- eval $(docker-machine env docker-ci) : 도커 머신 연결
- git clone https://github.com/marcelbirkner/docker-ci-tool-stack.git : github 레파지토리 연결
- docker-compose up -d : tool-stack 실행


![CI](https://github.com/banziha104/DockerExample/blob/master/img/2week/docker-ci5.png)


## GitLab

> 이슈관리, 코드리뷰, CI/CD를 지원하는 통합 개발환경 서버

- 데이터 저장위치
    - /var/opt/gitlab : 어플리케이션 데이터 저장
    - /var/log/gitlab : 로그저장
    - /etc/gitlab : 설정파일 저장 
    - 포트 80,443,22 