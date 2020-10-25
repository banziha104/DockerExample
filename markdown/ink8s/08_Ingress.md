# 인그레스 

> 하나의 IP이나 도메인의 서비스 제공 

- http 요청의 호스팅 부분을 검사 

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: http-go-ingress
spec:
  rules:
  - host: dldudwnsdl.com
    http:
      paths:
      - path: /
        backend:
          serviceName: http-go-svc
          servicePort: 8080
  - host: web.dldudwnsdl.com
    http:
      paths: 
      - path: /
        backend:
          serviceName: http-go-np
          servicePort: 8080
            
```