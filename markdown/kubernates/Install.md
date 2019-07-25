# 쿠버네티스

> 컨테이너 운영을 자동화하기 위한 컨테이너 오케스트레이션 도구 

<br>


--- 

## Install( MacOS )
http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetesdashboard:/proxy/

-kuberctl 설치 
     
```bash
curl -L0 https://storage.googleapis.com/kubernetes-release/release/v1.10.4/bin/darwin/amd64/kubectl \
&& chmod +x kubectl \
&& mv kubectl /usr/local/bin
```

- dashboard 설치

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
```

- dashboard 테스트
```bash
kubectl get pod --namespace=kube-system -l k8s-app=kubernetes-dashboard
```

- access_token 만들기

