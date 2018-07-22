# 도커 호스트 구성 

1. 독립된 로컬 개발환경 : Vagrant, Docker Machine 을 이용하면 일관된 인터페이스로 구성이 가능
2. VM서버로 도커 Host를 구성 : Host에는 한개 컨테이너만 운영하는게 효과적, VM 유형, 환경에 상관없이 배포가 가능함
3. 싱글서버 도커 Host : 물리서버 1대로 도커 Host를 구성함
4. 도커 Host 클러스트 : 여러 도커 Host 서버를 하나의 Host처럼 관리하기 위해 클러스터를 구성한다. fleet , swarm, kubernets, mesos 


# 컨테이너 데이터 백업
