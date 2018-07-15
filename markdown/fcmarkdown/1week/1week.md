# Dokcer

> **리눅스 컨테이너** 기술을 이용해 어플리케이션 기술을 이용해 어플리케이션 **패키징,배포**를 지원하는 **경량**의 **가상화** 오픈소스 프로젝트

- 리눅스 컨테이너 : 단일 리눅스 호스트상에서 여러개의 격리된 리눅스 시스템을 실행하기 위한 **OS 수준의 가상화**
    - Libvirt : OS 수준의 가상화 오픈소스 기술, Docker의 Low Level 단의 기술은 Libvert로 이루어짐
    - libContainer :  

- Container와 Hypervisor 
    - Hypervisor : GuestOS를 사용
    - Container : HostOS를 사용
    - 상호보완적으로 성장중
    
- Immutable Infrastructure 
    - 이미지 기반의 어플리케이션 배포 패러다임
    - 많은 서버를 동적으로 관리하는 클라우드 환경에서 효과적이고 유연한 배포 방식
    

---

# Docker 제품군 및 제품환경

1. Dokcer Client
    - 도커 엔진과 통신하는 프로그램
    - Kitemetic
    
2. Docker Host OS : 도커 엔진을 설치할수 있는 환경, 64비트 리눅스 커널 버전 3.10 이상 환경
    - Ubuntu, CentOS, RHEL
    - Container OS(구 CoreOS)
    - Debian , Suse , Fedora, etc
    
3. Docker Engine 
    - 어플리케이션을 컨테이너로 만들고 실행하게 해주는 데몬
    - Swarm, Kubernetes와 통합 
    - RESTful 기반
    - Client Docker 
    
4. Docker Machine : 로컬, 리모트 서버에 도커 엔진을 설치하고 , 설정을 자동으로 해주는 프로비저닝 클라이언트 

5. Docker Hub 
    - 도커 이미지를 관리하는 저장소
    - 오픈소스 공식 이미지 관리

6. Docker Trusted Registry
    - 도커 이미지 저장소를 구축
    - 설치형, 인증 지원
    - Registry (무료) , DTR(유료)
    
7. Docker Universal Control Plan 
    - 도커 클러스터 관리도구
    - 유료 
    
8. Docker Data Center 
    - DUCP + DTR
    
9. Docker Cloud 
    - 도커 이미지, 컨테이너 관리를 지원
    - 공개/비공개  클라우드 관리 가능
    
10. Docker Choice : 클라우드, OS에 비종속적
11. Agility : 생산성, 효율증가
12. Security : 어플리케이션으로부터 데이터 분리

    
    