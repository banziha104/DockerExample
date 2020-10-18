# 서비스 

- 포드는 일시적으로 생성한 컨테이너의 집합
- 떄문에 포드가 지속적으로 생겨났을때 서비스를 하기에 적합하지 않음
- IP 주소의 지속적인 변동,로드밸런싱을 관리해줄 또 다른 개체가 필요
- 이를 해결하는 것이 **서비스**라는 리로스가 존재 
- 서비스 세션 고정하기
    - 서비스가 다수의 포드로 구성하면 웹서비스의 세션이 유지되지 않음.
    - 이를 위해 처음 들어왔던 클라이언트 IP를 유지해주는 방법이 필요
    - sessionAffinity: ClientIP 라는 옵션을 주면 서비스 세션이 고정됨.
- 다중 포트 연결 방법 : 포트에 그대로 나열해서 사용
    
```shell script
kubectl expose # 가장 쉽게 서비스를 만듬
kubectl create -f http-go-svc.yaml # 그러나 대부분 yaml로 만듬 
``` 

<br>

# 서비스 외부 노출

- NodePort(Type): 노드의 자체 포트를 사용하여 포드로 리다이렉션
- LoadBalancer(Type): 외부 게이트웨이를 사용해 노드 포트로 리다이렉션
- Ingress: 하나의 IP 주소를 통해 여러 서비스를 제공하는 특별한 매커니즘 
- 노드 포트 생성
    - 서비스 yaml 파일을 작성
    - type에 NodePort를 지정
    - 30000-32767 포트만 사용가능
    

```shell script

```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: http-go-np
spec:
  type: NodePort
  selector:
    run: http-go
  ports:
    - protocol: TCP
      port: 80 # 서비스의 포트
      targetPort: 8080 # 포드의 포트
      nodePort: 30001 # 최종적으로 서비스되는 포트 
```