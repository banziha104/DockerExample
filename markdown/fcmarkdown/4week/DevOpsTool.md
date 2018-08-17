# DevOps Tool

- Code  : 형상관리
- Build  : CI
- Test : 자동화된 테스트
- Package : 패키징 도구,패키지 저장소
- Release : 배포자동화
- Configure 
    - 설정관리
    - 인프라 자동화
- Monitor: 성능 모니터링
- 오케스트레이션 
    - Panamax
        - 도커 컨테이너 관리도구
        - Ruby on Rails, CoreOS, github 기반
        - Template 개념, 편리한 UI
    - spotify/helios
        - 도커 이미지, 컨테이너 라이프사이클 관리
        - Java, CLI
        - Job 단위로 처리, 스케쥴링 X
    - d4
    - Skaffold
    - Rancher 
    
- Ansibles
    - 레드햇이 만든 IT 자동화 언어/ 엔진
    - 파이썬 기반
    - 오픈소스
    - Ansible Tower는 상용
    - 리눅스 계열 OS
    - 특징
        - 코딩없이 YAML 형식으로 할일을 정의하기 때문에 이해하기 쉬움
        - 에이전트를 설치하지 않고 호스트에 SSH로 접속할 수 있다면 적용가능
        - 다양한 모듈이 제공되어 작업을 정의하기 쉽다. 
        
    - 구성요소
        - 플레이북 : 앤서블을 어떻게 실행할지를 정의하는 YAML 파일
        - 모듈 : 앤서블에 실행하는 명령 단위
        - 인벤터리 :서버 목록
        - 갤럭시 : 앤서블 커뮤니티 
    

# Skaffold

- 쿠버네티스 기반의 CLI 배포 도구
- 로컬/ 리모트 k8s 클러스터 지원
- 리눅스 , 맥 지원
- 빌드 파이프라인 지원 : 빌드, 푸시, 디플로이
- 설치 (Minikube, Kubernetes 등이 설치되어있어야함) 및 사용
    
    1. 컬로 받아오기
    
    ```bash
    curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-darwin-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin
    ```
    
    2. 깃 클론
    
    ```bash
    git clone https://github.com/GoogleContainerTools/skaffold
    ```
    
    3. 실행
    
    ```bash
    skaffold dev 
    ```
    
    4. skaffold.yaml
    
    ```bash
    apiVersion : skaffold/v1alpha2
    kind: Config
    build:
        artifacts:
        -  imageName: gcr.io/k8s-skaffold/skaffold-example
    deploy:
        kubectl:
            manifests:
                - k8s-*
    profiles:
        - name: gcb
          build:
            googleCloudBuild:
                projectId: k8s-skaffold
               
    
    ```
    
    5. main.go 수정 

# Rancher

> 컨테이너 인프라스트럭처 관리도구, 컨테이너 OS

- Rancher 주요 기능
    - 멀티 호스트 네트워킹: Private + PublicProvider
    - 컨테이너 로드밸런싱 : HA Proxy 기반
    - 스토리지 서비스 : Docker1.9 Volume 호환
    - 서비스 디스커리 : 분산 DNS 기반, 헬스체크 (Agent+Haproxy)
    - 서비스 업데이트 : 롤링업데이트
    - 리소스 관리 : 리소스 모니터링 및 스케쥴링
    - 사용자 관리 : LDAP 연동
    - 멀티 오케스트레이션 앤진 : Cattle, k8s, Swarm
    
- 설치 및 사용

1. rancher-host 머신 생성

```bash
docker-machine create -d virtualbox rancher-host
```

2. rancher 서버 실행

```bash
docker run -d --restart=always -p 8080:8080 rancher/server:stable
```

3. 호스트 추가 : rancher-node1 머신 생성

```bash
docker-machine create -d virtualbox rancher-node1
```

4. 호스트 추가 : rancher agent 설치

```bash
docker run -d --privileged -v /var/run/docker.sock:/var/run/docker.sock \
-v /var/lib/rancher:/var/lib/rancher rancher/agent:v1.0.2


```