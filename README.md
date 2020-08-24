# t

ツイートするためだけの自分用コマンド。

## howto
1. 環境変数に以下を登録
```
export T_CONSUMER_KEY=""
export T_CONSUMER_SECRET=""
export T_ACCESS_TOKEN=""
export T_ACCESS_SECRET=""
```
2. `go build -o t main.go`
3. `t --text "hoge"`
