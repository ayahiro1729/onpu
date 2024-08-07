# ベースイメージを指定
FROM node:22-alpine

# 作業ディレクトリを設定
WORKDIR /app

# package.jsonとpackage-lock.jsonをコピー
COPY ./frontend/package*.json ./

# 依存関係をインストール
RUN npm cache clean --force
RUN npm install -g npm@latest
RUN npm install

# アプリケーションのソースコードをコピー
COPY ./frontend .

# ポートを指定
EXPOSE 3000

# アプリケーションを起動
CMD ["npm", "run", "dev"]
