Docker/k8s実践コンテナ開発   cf41, 72, 73

Docker-compose.ymlに以下の場合だと事前に
イメージをbuildをしておく必要がある

image: example/echo:latest

docker image build -t example/echo:latest .

次に以下のコマンドでサーバーを立ち上げる

docker-compose up -d

サーバーを停止　削除は以下のコマンド

docker-compose down