# 쿠버네티스 주요 개념

<br>

| 리소스               | 용도                                                                           |
|----------------------|--------------------------------------------------------------------------------|
| Node                 | 컨테이너가 배치되는 서버                                                       |
| Namespaces         | 쿠버네티스 클러스트 안의 가상의 클러스터                                       |
| Pod                 | 컨테이너의 집합 중 가장 작은 단위로, 컨테이너의 실행 방법을 정의               |
| 레플리카 세트        | 같은 스펙을 갖는 파드를 여러 개 생성하고 관리하는 역할을함.                    |
| Deployments         | 레플리카 세트의 리비전을 관리함                                                |
| Service               | 파드의 집합에 접근하기 위한 경로를 정의                                        |
| Ingress             | 서비스를 쿠버네티스 클러스터 외부로 노출시킨다.                                |
| Config Map             | 설정 정보를 정의하고 파드에 전달한다.                                          |
| Persistent Volume       | 파드가 사용할 스토리지의 크기 및 종류를 정의                                   |
| Persistent Volume Frame | 퍼시스턴트 볼륨을 동적으로 확보                                                |
| Storage Class       | 퍼시스턴트 볼륨이 확보하는 스토리지의 종류를 정의                              |
| Strateful Set    | 같은 스펙으로 모두 동일한 파드를 여러 개 생성하고 관리한다.                    |
| Job                   | 상주 실행을 목저으로 하지 않는 파드를 여러 개 생성하고 정상적인 종료를 보장함. |
| Crone Job               | 크론 문법으로 스케줄링되는 잡                                                  |

<br>

## Cluster & node

- 클러스터는 쿠버네티스의 여러 리소스를 관리하기위한 집합체
- 리소스 중 가장 큰 개념은 노드.
- 노드는 쿠버네티스 클러스터의 관리 대상으로 등록된 도커 호스트.
- 쿠버네티스 클러스터 전체를 관리하는 서버인 마스터가 적어도 하나 이상 있어야 함.
- 쿠버네티스는 노드의 리소스 사용 현황 및 배치 전략을 근거로 적절히 배치한다.
- 노드의 목록조회 

```
kubectl get nodes
```

<br>
 
## Namespace 

- 쿠버네티스는 클러스터 안에 가상 클러스터를 또 다시 만들 수 있다.
- 이를 네임스페이스라고함
- 네임스페이스는 개발팀이 일정 규모 이상일 때 유용함.
- 각 네임스페이스 마다 권한을 부여하면, 다른 네임스페이스를 어지르지않음.
- 네임스페이스 조회

```
kubectl get namespace
```

<br>

## Pod

- 컨테이너가 모인 집합체
- 하나 이상의 컨테이너로 이루어짐
- 결합이 강한 경유 파드를 묶어 일괄 배포함.
- 파드는 하나의 노드위에만 존재할수 있으며 여러 노드에 걸칠 수 없음
- Nginx와 그 뒤에 위치할 애플리케이션 컨테이너를 함게 파드로 묶는 구성이 일반적
- 마스터 노드가 파드를 생성

<br>

- 파드 생성하기

```yaml
apiVersion: v1
kind: Pod # 쿠버네티스의 리소스를 지정하는 유형
metadata:
  name: simple-echo #리소스의 이름
spec: #리소스를 정의하기 위한 속성
  containers:
    - name: nginx
      image: gihyodocker/nginx:latest
      env:
      - name: BACKEND_HOST
        value: localhost:8080
      ports:
      - containerPort: 80
    - name: echo
      image: gihyodocker/echo:latest
      ports:
      - containerPort: 8080
```

```bash
kubectl apply -f simple-pod.yaml
```

<br>

- 파드 상태 조회

```
kubectl get pod
```

<br>

- 파드 접속 

```bash
kubectl exec -it simple-echo sh -c nginx # 컨테이너가 여러개인 경우, -c 옵션으로 컨테이너 지정
```

<br>

- 파드 로그 가져오기

```bash
kubectl log -f simple-echo -c echo 
```

<br>

- 파드 삭제하기

```
kubectl delete pod simple-echo # 직접삭제
kubectl delete -f simple-pod.yaml # 파일을 지정해 삭제
```

## 레플리카세트 

> 레플리카세트는 똑같은 정의를 갖는 파드를 여러 개 생성하고 관리하기 위한 리소스. 

- 어느 정도 규모가 되는 애플리케이션을 구축하려면 같은 파드를 여러 개 실행해 가용성을 확보해야 하는 경우가 생김
- 이런 경우에 사용하는 것이 레프리카세트
- 레플리카세트를 정의한 yaml에 파드의 정의 자체도 작성함으로 따로 둘 필요 없음. 

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: echo
  labels:
    app: echo
spec:
  replicas: 3
  selector: 
    matchLabels:
      app: echo
  template:
    metadata:
      labels:
        app: echo
    spec:
      containers:
      - name: nginx
        image: gihyodocker/nginx:latest
        env:
          - name: BACKEND_HOST
            value: localhost:8080
        ports:
          - containerPort: 80
      - name: echo
        image: gihyodocker/echo:latest
        ports:
          - containerPort: 8080

```

<br>

- 레플리카셋 생성

```yaml
kubectl apply -f simple-replicaset.yaml
```

<br>

- 레플리카셋 삭제

```yaml
kubectl delete -f simple-replicaset.yaml
```

<br>

## 디플로이먼트

> 레프리카세트보다 상위에 해당하는 리소스로 , 어플리케이션 배포의 기본 단위

- 디플로이먼트는 레플리카세트를 관리하고 다루기 위한 리소스
- 레플리카셋과 yaml 이 크게 다르지 않음, 대신 디플로이먼트가 레플리카세트의 리비전 관리를 할 수 있다.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echo
  labels:
    app: echo
spec:
  replicas: 3
  selector:
    matchLabels:
      app: echo
  template:
    metadata:
      labels:
        app: echo
    spec:
      containers:
      - name: nginx
        image: gihyodocker/nginx:latest
        env:
          - name: BACKEND_HOST
            value: localhost:8080
        ports:
          - containerPort: 80
      - name: echo
        image: gihyodocker/echo:latest
        ports:
          - containerPort: 8080
```

<br>

- 실행 
```bash
kubectl apply -f simple-deply
```
