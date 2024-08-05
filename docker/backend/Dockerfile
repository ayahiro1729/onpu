# ベースイメージを指定
FROM golang:1.22-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Goのモジュールファイルをコピー
COPY ./backend/go.* ./

# 依存関係を取得
RUN go mod download
RUN go mod tidy

# アプリケーションのソースコードをコピー
COPY ./backend .

# アプリケーションをビルド
RUN go build -o /app/main .

# ポートを指定
EXPOSE 8080

# アプリケーションを起動
CMD ["/app/main"]
