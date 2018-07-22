# Docker 

- docker run -d -p 8080:8080 -p 5000:5000 --name myjenkins jenkins : 두 개의 포트를씀
- docker run -p 80:80 -v /home/docker/nginx.conf:/etc/nginx/nginx.conf --link myjenkins:jenkins -d nginx


# 도커 호스트 구성 

1. 독립된 로컬 개발환경 : Vagrant, Docker Machine 을 이용하면 일관된 인터페이스로 구성이 가능
2. VM서버로 도커 Host를 구성 : Host에는 한개 컨테이너만 운영하는게 효과적, VM 유형, 환경에 상관없이 배포가 가능함
3. 싱글서버 도커 Host : 물리서버 1대로 도커 Host를 구성함
4. 도커 Host 클러스트 : 여러 도커 Host 서버를 하나의 Host처럼 관리하기 위해 클러스터를 구성한다. fleet , swarm, kubernets, mesos 

# 컨테이너 데이터 백업

1. 컨테이너의 /var/lib/mysql 에 데이터를 호스트의 /mysqldata로 백업 


```bash
docker  run -d \
> -e MYSQL_ROOT_PASSWORD='test1234' \
> --name mydb \
> -v /mysqldata:/var/lib/mysql mysql
```

# 데이터 컨테이너

> 다른 컨테이너에게 볼륨을 공유하기 위해 만드는 컨테이너, 내용이 쉐어링됨
> 어플리케이션을 실행하지 않음

- docker run --name mydata -v /data/app1 busybox true : 데이터 컨테이너 만들기 busybox는 데이터 컨테이너를 만들때 많이 사용 
- docker run -it --volumes-from mydata ubuntu : --volume-from 옵션은 볼륨을 mydata라는 컨테이너에 맡김, it 옵션은 생성 실행 후 해당 컨테이너에 접속 
-
-
-
-

