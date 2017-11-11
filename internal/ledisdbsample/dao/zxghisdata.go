package dao

import (
"fmt"
"time"
)

type ZxgHisData struct {
	DefGroupId      string
	HGroupId        string
	GroupInfoList   []ZxgHisGroupInfo
	GidSortDic      map[string]float64
	GidStockInfoDic map[string]map[string]ZxgHisStockContent
	GidStockSortDic map[string]map[string]float64
	GidStockNoteDic map[string]map[string]ZxgHisStockNote
	GidStockStarDic map[string][]string
	Date            string
}

type ZxgHisGroupInfo struct {
	Name        string
	StockCount  int
	Version     int64
	Gid         int64
	Source      string
	UpdatedTime int64
}

type ZxgHisStockContent struct {
	IsStarStock  bool
	Price        string
	Timestick    int64
	IsGivenStock bool
}

type ZxgHisStockNote struct {
	StockCount      string
	StockPrice      string
	BrokerageFee    string
	Factorage       string
	StampDuty       string
	TargetPrice     string
	StopLossPrice   string
	Date            string
	Note            string
	TransactionDate string
}

func USER_HISTORY_HASH_KEY(uid string) string {
	return fmt.Sprintf("historyhash_%v", uid)
}

func USER_HISTORY_ITEM_KEY(date time.Time) string {
	return date.Format("20060102")
}

