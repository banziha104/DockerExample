# 아키텍쳐

- Docker 는 Git 처럼 다른 부분만 태그로 관리해서 실행함 (처음 실행 이후 다음 실행은 매우 빠름)

- Namesapaces : 프로세스별 리소스 **리소스 격리**
    - 가벼운 프로세스 가상화 기술
    - 격리된 작업공가능ㄹ 구성
    - 유형 
        - mny : 파일시스템
        - pid : 프로세스
        - net : 네트워크 스택
        - ipc : System V IPC
        - ust : hostname
        - user : UIDS

- Control groups (cgroups) : 리소스 *관리*
    - 실행중인 어플리케이션이 원하는 만큼 리소스를 사용하게 한다
    - 특정 컨테이너가 지정한 만큼 리소스를 쓰도록 제한
    - 2006년 구글에서 시작
    - 관리소스 
        - memory : mm/memocontrol.c
        - cpuser : kernel/spuset.c
        - net_prio : net/core/netprio_cgroup.c
        - devices : security/device_cgroup.c
        
- Union Filesystem 
    - 다른 파일 시스템을 Union mount 하도록 해주는 리눅스 파일 시스템 서비스
    - 브랜치로하는 여러 파일 시스템을 하나의 시스템처럼 사용할 수 있게 해줌
    
- 권한수준
    - root 사용자와 동일한 수준의 권한이 필요
    - docker 데몬을 실행하기 위해서는 root 사용자필요
    - TCP 포트가 아닌 Unix 소켓과 데몬을 바인딩하기 때문
    - 사용자, 사용자 그룹 모두 docker 라는 이름을 사용
    