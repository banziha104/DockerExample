## 레이블 

- 모든 리소스를 구성하는 매우 간단하면서도 강력한 쿠버네티스 기능
- 리소스에 첨부하는 임의의 키/값 
- 레이블 셀렉터를 사용하여 각종 리소스를 필터링 할 수 있음
- 리소스는 한 개 이상의 레이블을 가질 수 있음.
- 리소스를 만드는 시점에 전부 레이블을 첨부
- 기존 리소스에도 레이블의 값을 수정 및 추가 가능
- 모든 사람이 쉽게 이해할 수 있는 체계적인 시스템을 구축 가능 

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: label-demo
  labels:
    environment: production
    app: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80

```

- 라벨 포함 정보

```shell script
kubectl get pod --show-labels
```

- 라벨 추가 및 삭제

```shell script
kubectl label pod http-go-v2 rel=beta # 레이블 추가
kubectl label pod http-go-v2 rel=beta --overwrite # 레이블 수정
kubectl label pod http-go-v2 rel- # 삭제

```

<br>

- 레이블 필터링

```shell script
kubectl get pod --show-label -l 'env!=test,rel=beta'
kubectl get -L creation_method
```

<br>


