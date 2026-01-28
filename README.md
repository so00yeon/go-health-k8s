# go-health-k8s

Go 언어 학습과 Kubernetes 배포 흐름을 익히기 위한 개인 학습 프로젝트입니다.

## 기술 스택

- Go 1.22+
- Docker (멀티스테이지 빌드)
- Kubernetes (minikube)

## 프로젝트 구조

```
go-health-k8s/
├── cmd/server/          # 서버 진입점
├── internal/
│   ├── config/          # 환경변수 설정
│   ├── handler/         # HTTP 핸들러
│   ├── middleware/       # 미들웨어 (로깅)
│   └── router/          # 라우터
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

## 테스트

```bash
go test ./...
```
