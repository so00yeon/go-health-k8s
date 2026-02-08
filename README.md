# go-health-k8s

Go 언어 학습과 Kubernetes 배포 흐름을 익히기 위한 개인 학습 프로젝트입니다.

## 기술 스택

- Go 1.22+
- Docker (멀티스테이지 빌드)
- Kubernetes (minikube)

## 아키텍처

```
┌──────────────┐      ┌──────────────────────────┐
│   Client     │─────▶│  K8s Service (ClusterIP)  │
└──────────────┘      └────────────┬─────────────┘
                                   │
                      ┌────────────▼─────────────┐
                      │   K8s Deployment (x2)     │
                      │  ┌──────────────────────┐ │
                      │  │  Go HTTP Server       │ │
                      │  │  ├── /health          │ │
                      │  │  └── /version         │ │
                      │  └──────────────────────┘ │
                      │  ConfigMap (env vars)      │
                      └────────────────────────────┘
```

## 프로젝트 구조

```
go-health-k8s/
├── cmd/server/             # 서버 진입점
├── internal/
│   ├── config/             # 환경변수 설정
│   ├── handler/            # HTTP 핸들러 + 테스트
│   ├── middleware/          # 미들웨어 (로깅)
│   └── router/             # 라우터
├── k8s/base/               # Kubernetes 매니페스트
│   ├── configmap.yaml
│   ├── deployment.yaml
│   └── service.yaml
├── .github/workflows/      # CI/CD
├── Dockerfile
├── .dockerignore
├── .gitignore
├── go.mod
└── README.md
```

## 실행

```bash
go run cmd/server/main.go
```

환경변수로 설정을 변경할 수 있습니다:

```bash
PORT=3000 APP_ENV=production APP_VERSION=1.0.0 go run cmd/server/main.go
```

## API

| Method | Path      | 설명             |
|--------|-----------|------------------|
| GET    | /health   | 헬스체크          |
| GET    | /version  | 버전 및 환경 정보  |

## Docker

```bash
# 빌드
docker build -t go-health-k8s:latest .

# 실행
docker run -p 8080:8080 -e APP_ENV=production go-health-k8s:latest

# 확인
curl http://localhost:8080/health
```

## Kubernetes (minikube)

```bash
# minikube 시작
minikube start

# minikube 내부 Docker 데몬 사용
eval $(minikube docker-env)

# 이미지 빌드 (minikube Docker에서 직접)
docker build -t go-health-k8s:latest .

# 매니페스트 적용
kubectl apply -f k8s/base/configmap.yaml
kubectl apply -f k8s/base/deployment.yaml
kubectl apply -f k8s/base/service.yaml

# 상태 확인
kubectl get pods
kubectl get svc

# 포트포워딩으로 접근
kubectl port-forward svc/go-health-k8s 8080:80
curl http://localhost:8080/health
```

## 테스트

```bash
go test ./...
```

## 배운 점

- **Go**: 패키지 구조(`cmd/` / `internal/`), 인터페이스 기반 설계, `net/http` 표준 라이브러리만으로 충분한 HTTP 서버 구축이 가능하다는 점
- **Docker**: 멀티스테이지 빌드로 최종 이미지 크기를 최소화하는 방법, `distroless` 베이스 이미지의 장점
- **Kubernetes**: Deployment/Service/ConfigMap의 역할 분리, `envFrom`으로 설정 주입, liveness/readiness probe를 통한 컨테이너 상태 관리
- **CI/CD**: GitHub Actions에서 테스트 → 이미지 빌드 → 레지스트리 푸시까지 이어지는 파이프라인 구성
