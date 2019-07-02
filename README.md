# globalip-route53

グローバルIPアドレスをRoute53のAレコードに登録するアプリです。

## 使い方

### 環境変数を設定
```
export AWS_ROUTE53_DOMAIN=<route53のドメインのID(14文字のID)>
export AWS_ROUTE53_RECORD=<レコード名(www.sample.com)>
export AWS_ACCESS_KEY_ID=<シークレットキー>
export AWS_SECRET_ACCESS_KEY=<アクセスキー>
```

### 普通に使う方法
```
go get -u -v github.com/ryicoh/globalip-route53
$GOATH/bin/globalip-route53
```

```
git clone https://github.com/ryicoh/globalip-route53
cd globalip-route53
go build
./globalip-route53
```

### Dockerから使う方法
```
docker run -it --rm \
  -e AWS_ROUTE53_DOMAIN=<route53のドメインのID(14文字のID)> \
  -e AWS_ROUTE53_RECORD=<レコード名(www.sample.com)> \
  -e AWS_ACCESS_KEY_ID=<シークレットキー> \
  -e AWS_SECRET_ACCESS_KEY=<アクセスキー> \
  ryicoh/globalip-route53
```
