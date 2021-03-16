## ゲーム名

コンバット

## URL リンク

https://ishishow.github.io/tank/

SwaggerEditor: <https://editor.swagger.io> <br>
定義ファイル: `./api-document.yaml`<br>

※ Firefox はブラウザ仕様により上記サイトから localhost へ向けた通信を許可していないので動作しません

- https://bugzilla.mozilla.org/show_bug.cgi?id=1488740
- https://bugzilla.mozilla.org/show_bug.cgi?id=903966

## 事前準備

### docker-compose を利用した MySQL と Redis の準備

#### MySQL

```
$ docker-compose up mysql
```

#### Redis

```
$ docker-compose up redis
```

を実行することでローカルの Docker 上に MySQL サーバが起動します。

### API 用のデータベースの接続情報を設定する

Mac の場合

```
$ export MYSQL_USER=root \
    MYSQL_PASSWORD=ca-tech-dojo \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=dojo_api
```

Windows の場合

```
$ SET MYSQL_USER=root
$ SET MYSQL_PASSWORD=ca-tech-dojo
$ SET MYSQL_HOST=127.0.0.1
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=dojo_api
```

## API ローカル起動方法

```
$ go run ./cmd/main.go
```
