FROM golang:1.22-alpine

WORKDIR /app

ARG ENV=development

# 依存関係を事前にコピーしてインストール
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
RUN go mod tidy

# ソースコードをコピー
COPY ./backend .

# 本番用ビルド
RUN if [ "$ENV" = "production" ] ; then \
        go build -o main ./cmd/server ; \
    fi

# 開発環境用のツールインストール
RUN if [ "$ENV" = "development" ] ; then \
        go install github.com/cespare/reflex@latest ; \
    fi

EXPOSE 8080

# 環境に応じて起動コマンドを変更
CMD if [ "$ENV" = "production" ] ; then \
        ./main ; \
    else \
        reflex -r '\.go$' -s go run ./cmd/server ; \
    fi
