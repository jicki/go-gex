// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameKline = "kline"

// Kline mapped from table <kline>
type Kline struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	StartTime int64  `gorm:"column:start_time;not null;comment:k线开始时间" json:"start_time"`
	EndTime   int64  `gorm:"column:end_time;not null;comment:k线结束时间" json:"end_time"`
	Symbol    string `gorm:"column:symbol;not null;comment:交易对" json:"symbol"`
	SymbolID  int32  `gorm:"column:symbol_id;not null;comment:交易对id" json:"symbol_id"`
	KlineType int32  `gorm:"column:kline_type;not null;comment:k线类型1分钟 5分钟" json:"kline_type"`
	Open      string `gorm:"column:open;not null;comment:开盘价" json:"open"`
	High      string `gorm:"column:high;not null;comment:k线内最高价" json:"high"`
	Low       string `gorm:"column:low;not null;comment:k线内最低价" json:"low"`
	Close     string `gorm:"column:close;not null;comment:收盘价" json:"close"`
	Amount    string `gorm:"column:amount;not null;comment:成交量(基础币数量)" json:"amount"`
	Volume    string `gorm:"column:volume;not null;comment:成交额(计价币数量)" json:"volume"`
	Range     string `gorm:"column:range;not null;comment:涨跌幅" json:"range"`
}

// TableName Kline's table name
func (*Kline) TableName() string {
	return TableNameKline
}