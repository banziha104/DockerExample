# Vagrant

> 다양한 가상머신을 동일한 방식으로 만들고 관리할 수 있게 해주는 도구
> 루비 언어를 따르는 Vagrantfile 에 필요한 설정을 정의

- Vagrant Box : Vagrant 로 만들 수 있는 가상서버 유형
- Vagrantfile에 포함되는 내용
    - Box : 패키징할 서버 유형
    - 포트 바인딩 
    - 네트워크 주소
    - 이름,CPU,Memory
    - 부트 스트래핑 : 어플리케이션 설치
    - 프로비저닝 : 어플리케이션 설정
    - 폴더 마운트 
    
    
### CoreOS

> Docker 호스트에 특화된 리눅스

- 최소화된 메모리 사용
- systemd(리눅스 구동시 서비스를 자동으로 뜨게 해), etcd(Key-value 스토어), fleet()
- CoreUpdate(상용)
- https://coreos.com/

