# bitFlyer

株式会社bitFlyer様が提供している、bitFlyer Lightning APIのSDKをGoで実装してみました。

[APIマニュアル](https://lightning.bitflyer.jp/docs)

## Go version

Go 1.5

## Installation

```
go get github.com/pokochi/bitFlyer
```

## Operating instructions

* Public API

```
//引数なし
b := bitFlyer.PublicApi{}
err, s := b.GetHealth()

//引数あり
b := bitFlyer.PublicApi{}
m := make(map[string]string)
m["product_code"] = "BTC_JPY"

err, s := b.GetBoard(m)
```

* Private API

```
//引数なし
b := bitFlyer.NewPermissions("Your API KEY", "Your API SECRET")
err, s := b.GetPermissions()

//引数あり
b := bitFlyer.NewTrade("Your API KEY", "Your API SECRET")
m := make(map[string]string)
m["product_code"]     = "BTC_JPY"
m["child_order_type"] = "MARKET"
m["side"]             = "BUY"
m["size"]             = "0.01"
m["minute_to_expire"] = "10"

err, s := b.SendChildOrder(m)
```

## Todo

* [新規の親注文を出す（特殊注文）](https://lightning.bitflyer.jp/docs#新規の親注文を出す特殊注文)の実装。
* エラー処理のリファクタ

## License

MIT
