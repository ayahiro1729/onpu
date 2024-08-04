# onpu

## 環境構築
make build
make up

cd frontend
npm install

cd backend
go mod download

## frontend
### lint実行
cd frontend
npm run lint

## backend
### lint実行
cd backend
golangci-lint run --fix
