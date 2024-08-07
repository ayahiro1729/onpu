# ベースイメージを指定
FROM golang:1.22-alpine

# 作業ディレクトリを設定
WORKDIR /app

# キャッシュディレクトリを作成し、環境変数を設定
ENV GOCACHE=/app/.cache/go-build

# アプリケーションのソースコードをコピー
COPY ./backend .

# 依存関係を取得
RUN go install github.com/cespare/reflex@latest
RUN go mod download
RUN go mod tidy
RUN mkdir -p ./bin/ && go build -o ./bin/server ./cmd/server
RUN chown -R 10000:10000 ./bin/server

RUN apk update && \
    apk add --no-cache git

# ポートを指定
EXPOSE 8080

USER 10000:10000

# アプリケーションを起動
CMD ["sh", "-c", "reflex -r '(\\.go$|go\\.mod)' -s ./bin/server"]
