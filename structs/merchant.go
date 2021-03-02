package structs

type Merchant struct {
	AcceptBTC         bool                   `json:"accept_btc"`
	AcceptUSDT        bool                   `json:"accept_usdt"`
	AcceptBCH         bool                   `json:"accept_bch"`
	AcceptETH         bool                   `json:"accept_eth"`
	AcceptEOS         bool                   `json:"accept_eos"`
	AcceptLTC         bool                   `json:"accept_ltc"`
	AcceptBNB         bool                   `json:"accept_bnb"`
	AcceptBUSD        bool                   `json:"accept_busd"`
	AcceptCUSD        bool                   `json:"accept_cusd"`
	AcceptAlipay      bool                   `json:"accept_alipay"`
	AcceptWechat      bool                   `json:"accept_wechat"`
	WalletUserHash    string                 `json:"wallet_user_hash"`
	WalletUserEnabled bool                   `json:"wallet_user_enabled"`
	EmailVerified     bool                   `json:"email_verified"`
	Price             map[string]interface{} `json:"price"`
	Permission        string                 `json:"permission"`
}
