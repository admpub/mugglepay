# mugglepay

## 注册商家账号

[注册商家账号](https://merchants.mugglepay.com/user/register?ref=admpub)

## 开通麻瓜宝

MugglePay麻瓜宝，基于Telegram的电子钱包，匿名支付0手续费秒级到账。
[打开Telegram开通](https://telegram.me/MugglePayBot?start=TXZR9ZDF)

## 安装

```bash
go get github.com/admpub/mugglepay@latest
```

## 引用

```go
import (
    "github.com/admpub/mugglepay"
    "github.com/admpub/mugglepay/structs"
)
```

## 创建订单

```go
func CreateOrder(c *gin.Context) {
    mgp := mugglepay.New("BitpayxApplicationKey")
    // host := "https://www.example.com"
    serverOrder, _ := mgp.CreateOrder(&structs.Order{
        MerchantOrderID: orderID,
        PriceAmount:     money,
        // PriceCurrency:   "USD",
        // PayCurrency:     "ALIPAY",
        // PayCurrency:     "WECHAT",
        PayCurrency:     "",
        PriceCurrency:   "CNY",
        Title:           "订单标题",
        Description:     "订单描述",
        // 如需法币支付则必须设置正确的回调地址
        // CallbackURL: host + "/payment/notify",
        // CancelURL: host + "/user/code/return",
        // SuccessURL: host + "/user/code/return",
    })
    // 支付宝/微信扫码链接，该函数仅 PayCurrency 为 ALIPAY/WECHAT 时可返回地址
    // 其他情况下均返回加密货币地址
    // aliqr := serverOrder.ParseInvoiceAddress().Invoice.Address
    c.Redirect(http.StatusFound, serverOrder.PaymentURL)
}
```

## 支付回调校验

```go
func Notify(c *gin.Context) {
    body, _ := c.GetRawData()
    var callback structs.Callback
    if err := json.Unmarshal(body, &callback); err == nil {
        mgp := mugglepay.New("BitpayxApplicationKey")
        if mgp.VerifyOrder(&callback) {
            // code ... 
            c.JSON(200, gin.H{"status": 200})
            return
        }
    }
    c.JSON(200, gin.H{"status": 400})
}
```

## 修改支付方式

```go
mgp := mugglepay.New("BitpayxApplicationKey")
sorder, _ := mgp.CheckOut(ServerOrderId, "P2P_BTC")
// 应付金额
money := sorder.Invoice.PayAmount
// 法币支付链接
// 虚拟货币交易地址
address := sorder.Parse().Invoice.Address
// 虚拟货币交易备注
memo := sorder.Invoice.Memo

```

## 打赏
<a href="https://donate.cafe/hanks">
  <img style="width:150px" src="https://newbie.zeromesh.net/donate.7.6.svg" />
</a>
