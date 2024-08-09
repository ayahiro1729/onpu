# onpu

## 環境構築
make build
make up

cd frontend
npm install

cd backend
go mod download

## マイグレーション
cd backend

### マイグレーションファイルの作成
migrate create -ext sql -dir db/migrations -seq create_users_table

### マイグレーションの実行
export DATABASE_URL=postgres://user:password@localhost:5432/onpu?sslmode=disable
migrate -path db/migrations -database "$DATABASE_URL" up

## postgresqlコンテナに接続
docker exec db bash
psql -U user -d onpu

## frontend
### lint実行
cd frontend
npm run lint

## backend
### lint実行
cd backend
golangci-lint run --fix
