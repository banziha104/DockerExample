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

<br>

# 서비스

> 클러스트 안에서 파드의 집합에 대한 경로나 서비스 디스커버리를 제공하는 리소스

- 서비스는 쿠버네티스 클러스터 안에서만 접근할 수 있다.
- 서비스의 네임 레졸루션은 서비스명.네임스페이스명.svc.local로 연결됨
- ClusterIP : 쿠버네티스 클러스터의 내부 IP 주소를 공개할 수 있으며, 이를 이영해 어떤 파드에서 다른 파드 그룹으로 접근이 가능함. 단, 외부로부터는 접근할 수 없음
- NodePort: 클러스터 외부에서 접근할수 있는 서비스.

```yaml

# 클러스터 IP
apiVersion: v1
kind: Service
metadata:
  name: echo
spec:
  selector:
    app: echo
    release: summer
  ports:
    - name: http
      port: 8080 

```


- 서비스 생성

```bash
kubectl apply -f simple-service.yml
```

<br>

# 인그레스

> NodePort는 L4 레벨까지만 가능, 인그레스는 외부에 대한 노출과 가상 호스트 및 경로 기반. L7 레벨의 제어가 가능


- 클러스터 외부에서 온 HTTP 요청을 서비스로 라이팅을 하기 위한 nginx_ingress_controller 설치

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.16.2/deploy/mandatory.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.16.2/deploy/provider/cloud-generic.yaml
```

<br>

- 설치 확인

```bash
kubectl -n ingress-nginx get service,pod
```

<br>

- 서비스 등록

```bash
# 클러스터 IP
apiVersion: v1
kind: Service
metadata:
  name: echo
spec:
  selector:
    app: echo
    release: summer
  ports:
    - name: http
      port: 8080 
```

<br>

- 인그레스 파일 

```bash
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: echo
spec:
  rules:
    - host: ch05.gihyo.local
      http:
        paths:
          - backend:
              serviceName: echo
              servicePort: 80
``` 


<br>

# freshpod

> 컨테이미지가 업데이트됐는지를 탐지해 파드를 자동으로 배포해줌

- 배포될 파일의 이미지 정책 수정

```bash
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
        imagePullPolicy: IfNotPresent # IfNotPresent로 설정해야 작동함
        ports:
          - containerPort: 8080
```
