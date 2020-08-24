# t

ツイートするためだけの自分用コマンド。

## howto
1. .envファイルを用意
```
export CONSUMER_KEY=""
export CONSUMER_SECRET=""
export ACCESS_TOKEN=""
export ACCESS_SECRET=""
```
2. `go build -o t main.go`
3. `t --text "hoge"`
