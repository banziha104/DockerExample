# RESTful APi

- 주요 용어
    - 리소스 : 객체나 관련있는 데이터 집합, 리소스를 대상으로 처리가 이루어짐
    - 콜렉션 : 리소스의 집합
    - URL : 리소스에 대한 경로, 처리가 일어나는 것을 의미 

- API Endpoint
    - /getAllEmployees
    - GET /companies/3/employees
    
- HTTP 메소드
    - GET, POST, PUT, DELETE

- HTTP 응답코드
    - 2xx : 성공
    - 3xx : 리다이렉트
    - 4xx : 클라이언트 에러
    - 5xx : 서버 에러
    
- Versioning (* 이거참 중요한 것 같다)
    - api.yourservice.com/v1/xxxx
    
- 도메인 주도 설계 (DDD)
    - 도메인 : 특정 업종, 비즈니스,협업의 경험/지식
    - 소프트웨어는 현실문제를 해결하는것, 현실을 잘 반영하는 것이 좋은 소프트웨어를 만드는 지름길
    - 도메인 분석 -> 도메인 내 업무 문제를 정의하고 이를 모델로 표현하는 일
    - 도메인 전문가와 SW 전문가 모두가 이해할 수 있는 보편적인 언어를 사용해야함

- BoundedContext
    - 도메인 주도 개발 패턴, 큰 조직이나 모델을 설계하는 방식, 큰 모델을 독립적으로 동작하는 Context로 구분, Context는 다른 Context와 상호작용한다.
    