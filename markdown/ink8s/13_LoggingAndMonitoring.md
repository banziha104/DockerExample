# 쿠버네티스 모니터링 시스템과 아키텍쳐

- 쿠버네티스를 지원하는 다양한 모니터링 플랫폼
    - Heapster: Deprecated
    - Metrics Service : 힙스터를 deprecated하고 모니터링 표준으로 메트릭서버 도입
    - cAdviosr: kubelet에 모니터링하는 용도로 사용됨
    - 프로메테우스
    - eks
- 리소스 모니터링 도구
    - 쿠버네티스 클러스터 내의 애플리케이션 성능을 검사
    - 쿠버네티스는 각 레벨에서 애플리케이션의 리소스 사용량에 대한 상세 정보를 제공
    - 애플리케이션의 성능을 평가하고 병목 현상을 제거하여 전체 성능을 향상을 도모
    
- 로컬에 메트릭스 서버 설치 

```shell

kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.3.7/components.yaml
kubectl edit deployments.apps -n kube-system metrics-server

```

<br>

# 애플리케이션 로그 관리

- 로그는 컨테이너 단위로 로그 확인 가능
- 싱글 컨테이너의 포드의 경우, 포드까지만 지정하여 확인
- 멀티 컨테이너의 경우, 포드 뒤에 컨테이너 이름까지 전달하여 로그 확인
- kubectl logs <pod name> <container name>
- kubeapi가 정상 동작하지 않는 경우
    - 쿠버네티스에서 돌아가는 리소스들은 모두 docker를 사용
    - 따라서 docker의 로깅 기능을 사용
    - docker ps -a를 사용하여 조회 가능
    - docker logs <container id>를 사용하여 로그 확인 가능

<br>

# 큐브 대시보드 설치와 사용 

```shell

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.4/aio/deploy/recommended.yaml
kubectl edit service/kubernetes-dashboard -n kubernetes-dashboard # ClusterIP -> NodePort
# https://127.0.0.1:<Port>로 접속
kubectl get secret -n kubernetes-dashboard
kubectl describe secret -n kubernetes-dashboard kubernetes-dashboard-<id> 

``` 

<br>

# 프로메테우스

- GKE에 프로메테우스를 설치하려면 최소 노드당 CPU 2 이상의 퍼포먼스가 필요 
- 메트릭을 자동으로 수집하여 Prometheus 서버에서 수집하고 Grafana에 제공 
- 아키텍쳐
  - Prometheus StatefulSet
    - 구성된 모든 소스를 주기적으로 쿼리하여 구성된 모든 메트릭을 수집
    - 각 프로메테우스 Pod은 PVC에 저장 
  - Prometheus Node Export DaemonSet
    - 포드를 실행하고 있는 호스트 파일 시스템의 /sys, /proc을 수집 및 모니터링하여 노드의 하드웨어 및 운영 체제에 대한 값을 측정하고 메트릭을 NodeExport 포드의 포드 9100에 표시
  - Kube State Metrics Deployment
    - Kubernetes API 서버를 수신하고 리소스 (배포, 노드, 포드 등)와 관련된 메트릭을 생성
    - /metrics 포트 8080에서 메트릭을 표시
    - 프로메테우스 서버는 메트릭을 사용
  - Prometheus Alert Manager
    - 프로메테우스 서버에서 발생하는 경고를 수신하고 ConfigMap에 지정된 구성에 따라 처리
  - Granfana StatefulSet
    - 프로메테우스에 메트릭에 대한 쿼리를 위한 사용자 인터페이스를 제공하고 사전 구성된 대시 보드의 메트릭을 시각화

```shell

```

<br>

# 헬름 
