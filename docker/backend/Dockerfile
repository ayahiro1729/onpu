# ベースイメージを指定
FROM golang:1.22-alpine

# CompileDaemonをインストール
RUN go install github.com/githubnemo/CompileDaemon@latest

# 作業ディレクトリを設定
WORKDIR /app

# Goのモジュールファイルをコピー
COPY ./backend/go.* ./

# 依存関係を取得
RUN go mod download
RUN go mod tidy

# アプリケーションのソースコードをコピー
COPY ./backend .

# ポートを指定
EXPOSE 8080

# アプリケーションを起動
CMD ["CompileDaemon", "--build=go build -o /app/main /app/main.go", "--command=/app/main"]
