# globalip-route53

グローバルIPアドレスをRoute53のAレコードに登録するアプリです。

## 使い方

### 普通に使う方法
```
go get -uv github.com/ryicoh/globalip-route53
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
docker run -it --rm ryicoh/globalip-route53
```
