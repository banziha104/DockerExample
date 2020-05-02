# 쿠버네티스로 컨테이너 실행하기

<br>

- Kubectl 기본 사용법
    - kubectl [command] [type] [name] [flags]
    - command 실행하려는 동작
    - type 지원 타입 pod, service, ingress 등
    - name wkdnjsdlfma
    - flag 부가 옵션
    ```bash
    # 에코서버 실행
    kubectl run echoserver --generator=run-pod/v1 --image="k8s.gcr.io/echoserver:1.10" --port=8080
    
    #에코서버라는 이름의 서비스 생성
    kubectl expose po echoserver --type=NodePort
    
    # 파드 동작 확인 
    kubectl get pods 
    
    # 서비스 동작확인
    kubectl get services
    
    # 호스트와 포트바인딩
    kubectl port-forward svc/echoserver 8080:8080
  
    kubectl delete pod echoserver
    kubectl delete service echoserver
    ```

- kuberctl run 으로 컨테이너 실행하기
    - kubectl run 디플로이먼트이름 --image 컨테이너이미지이름 --port=포트번호 
    - 파드를 생성할때 기본 컨트롤러는 디플로이먼트
    ```bash
    # 컨테이너 실행 
    kubectl run nginx-app --image nginx --port=80
  
    # 파드 개수늘리기
    kubectl scale deploy nginx-app --replicas=2
    
    # 디플로이먼트 삭제
    kubectl delete deployment nginx-app
    ```
    
- 템플릿으로 컨테이너 실행하기 

```bash
kubectl apply -f nginx-app.yaml
```

- 클러스터 외부에서 클러스터 안 앱에 접근하기 
    - 쿠버네티스 내부에서 사용하는 네트워크가 외부와 격리되었기 때문에 접근 불가 
    - 외부에서 접근하기 위해서는 서비스를 생성해야함 
    ```bash
    # 서비스 생성 
    kubectl expose deployment nginx-app --type=NodePort
    
    #부 서비스 생성여
    kubectl get service
    # 해당서비스 자세한설명보기 
    kubectl describe service nginx-app
    ```
