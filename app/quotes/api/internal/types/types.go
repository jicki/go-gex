// Code generated by goctl. DO NOT EDIT.
package types

type KlineListReq struct {
	StartTime int64  `json:"start_time"` //开始时间 秒时间戳
	EndTime   int64  `json:"end_time"`   //结束时间
	KlineType int32  `json:"kline_type"` //k线类型
	Symbol    string `json:"symbol"`     //交易对
}

type Kline struct {
	Open       string `json:"open"`        //开
	High       string `json:"high"`        //高
	Low        string `json:"low"`         //低
	Close      string `json:"close"`       //收
	Amount     string `json:"amount"`      //成交量
	Volume     string `json:"volume"`      //成交额
	StartTime  int64  `json:"start_time"`  //开始时间
	EndTime    int64  `json:"end_time"`    //结束时间
	PriceRange string `json:"price_range"` //涨跌幅
	Symbol     string `json:"symbol"`      //交易对
}

type KlineListResp struct {
	KineList []*Kline `json:"kine_list"`
}

type GetDepthListReq struct {
	Symbol string `json:"symbol"` //交易对
	Level  int32  `json:"level"`  //档位
}

type Position struct {
	Qty    string `json:"qty"`    //数量
	Price  string `json:"price"`  //价格
	Amount string `json:"amount"` //金额
}

type GetDepthListResp struct {
	Version string      `json:"version"` //当前版本号
	Asks    []*Position `json:"asks"`    //卖盘
	Bids    []*Position `json:"bids"`    //买盘
}

type GetTickerListReq struct {
	Symbol string `json:"symbol"` //交易对
}

type Ticker struct {
	LastPrice   string `json:"last_price"`   //最新价
	High        string `json:"high"`         //高
	Low         string `json:"low"`          //低
	Amount      string `json:"amount"`       //成交量
	Volume      string `json:"volume"`       //成交额
	PriceRange  string `json:"price_range"`  //涨跌幅
	Last24Price string `json:"last24_price"` //24小时前的价格
	Symbol      string `json:"symbol"`       //交易对
}

type GetTickerListResp struct {
	TickerList []*Ticker `json:"ticker_list"`
}

type GetTickReq struct {
	Symbol string `json:"symbol"`         //交易对
	Limit  int32  `json:"limit,optional"` //获取多少条
}

type TickInfo struct {
	Price        string `json:"price"`  //价格
	Qty          string `json:"qty"`    //数量
	Amount       string `json:"amount"` //金额
	Timestamp    int64  `json:"timestamp"`
	Symbol       string `json:"symbol"`
	TakerIsBuyer bool   `json:"taker_is_buyer"`
}

type GetTickResp struct {
	TickList []*TickInfo `json:"tick_list"`
}