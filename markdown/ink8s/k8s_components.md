# 쿠버네티스 

- 주요 컴포넌트  확인

```shell script
kubectl get pods -n kube-system 
cd /etc/kubernetes/manifests/ # system과 관련된 yaml 설정 파일 저장 
```

- etcd: key-value 데이터 베이스


<br>

## Pod

- 컨테이너의 공통 배포된 그룹이며 쿠버네티스의 기본 빌딩 블록을 대표
- 쿠버네티스는 컨테이너를 개별적으로 배포하지 않고 컨테이너의 포드를 항상 배포하고 운영
- 일반적으로 포드는 단일 컨테이너만 포함하지만, 다수의 컨테이너를 포함할 수 있음
- 밀접하게 연관된 프로세스를 함께 실행하고, 하나의 환경에서 동작하는 것처럼 보임. 그러나 다소 격리된 상태로 유지 
- 플랫 인터 포드 네트워크 구조
    - 쿠버네티스 클러스터의 모든 포드는 공유된 단일 플랫, 네트워크 주소 공간에 위치
    - 포드 사이에는 NAT 게이트웨이가 존재 하지 않음 
- 두 가지의 컨테이너가 밀접한 경우에만 한 포드에 컨테이너를 둘, 아닌 경우 포드하나에 컨테이너 하나
- 포드 정의 구성 요소
    - apiVersion : 쿠버네티스의 api의 버전
    - kind : 어떤 리소스의 유형인지 결정
    - metadate : 포드와 관련된 이름, 네임스페이스, 라벨, 그 밖의 정보 존재
    - 스펙 : 컨테이너, 볼륨 등의 정보
    - 상태 : 포드의 상태, 각 컨테이너의 설명 및 상태, 포드 내부의 IP 및 그 밖의 기본 정보
    
```yaml

apiVersion: v1
kind: Pod
metadata:
  name: http-go
spec:
  containers:
  - name: http-go
    image: dldudwnsdl/http-go
    ports:
    - containerPort: 8080

```

- Liveness
    - 컨테이너가 살아있는지 판단하고 다시 시작하는 기능
    - 버그가 생겨도 높은 가용성을 보임
    - 컨테이너의 상태를 스스로 판단하여 교착 상태에 빠진 컨테이너를 재시작
    ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    labels:
      test: liveness
    name: liveness-http
  spec:
    containers:
    - name: liveness
      image: k8s.gcr.io/liveness
      args:
      - /server
      livenessProbe:
        httpGet:
          path: /healthz
          port: 8080
          httpHeaders:
          - name: Custom-Header
            value: Awesome
        initialDelaySeconds: 3
        periodSeconds: 3
    ```
- Readiness Probe
    - 포드가 준비된 상태에 있는지 확인하고 정상 서비스를 시작하는 기능
    - 포드가 적절하게 준비되지 않는 경우 로드밸런싱을 하지 않음
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: goproxy
      labels:
        app: goproxy
    spec:
      containers:
      - name: goproxy
        image: k8s.gcr.io/goproxy:0.1
        ports:
        - containerPort: 8080
        readinessProbe:
          tcpSocket:
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: 8080
          initialDelaySeconds: 15
          periodSeconds: 20
    ```
- Statup Probe
    - 어플리케이션의 시작 시기를 확인하여 가용성을 높이는 기능
    - Liveness와 Readiness의 기능을 비활성화
