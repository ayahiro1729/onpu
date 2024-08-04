# ベースイメージを指定
FROM node:22-alpine

# 作業ディレクトリを設定
WORKDIR /onpu

# package.jsonとpackage-lock.jsonをコピー
COPY ./frontend/package*.json ./

# 依存関係をインストール
RUN npm install

# アプリケーションのソースコードをコピー
COPY ./frontend .

# アプリケーションをビルド
RUN npm run build

# ポートを指定
EXPOSE 3000

# アプリケーションを起動
CMD ["npm", "start"]
