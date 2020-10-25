# 네임스페이스

- 리소스를 각각의 분리된 영역으로 나누기 좋은 방법
- 여러 네임스페이스를 사용하면 복잡한 쿠버네티스 시스템을 더 작은 그룹으로 분할
- 멀티 테넌트 환경을 분리하여 리소스를 생산, 개발, QA 환경 등으로 사용
- 리소스 이름은 네임스페이스 내에서만 고유 명칭 사용
- kubectl get 을 옵션없이 사용하면 default 네임스페이스에 질의
- 다른 사용자와 분리된 환경으로 타인의 접근을 제한
- 네임스페이스 별로 리소스 접근 허용과 리소스 양도 제어 가능
- --namespace나 -n 을 사용하여 네임스페이스 별로 확인 가능
 
 
 ```shell script
kubectl get ns # 네임스페이스 확인
kubectl get --all-namespaces # 전체 네임스페이스를 대상으로 실행 
kubectl create ns office # 네임스페이스 생성
kubectl create ns office2 --dry-run # 문법 맞는지 확인 
kubectl create ns office2 --dry-run -o yaml # yaml 파일  
kubectl create ns office2 --dry-run -o yaml > office-ns.yaml # yaml로 내보냄 
kubectl create deploy --image nginx -n office
gedit ~/.kube/config # 들어가서 context > context > namespace 를 수정하면 default 네임스페이스 변경
```

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: test-ns
```