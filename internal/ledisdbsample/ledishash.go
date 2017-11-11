package ledisdbsample

import (
	"jerome.com/gobaselearning/internal/ledisdbsample/dao"
	"time"
	"github.com/siddontang/ledisdb/ledis"
	lediscfg "github.com/siddontang/ledisdb/config"
	"fmt"
	"jerome.com/Hello/json"
	//"github.com/pelletier/go-toml"
)

func SetHash(key string, data []byte) {

	//itemkeystart := time.Now().UnixNano()

	cfg := lediscfg.NewConfigDefault()
	cfg.Addr = "172.31.32.248:7380"
	//cfg.DBName = "goleveldb"
	fmt.Println(cfg)
	l, err := ledis.Open(cfg)
	fmt.Println(err)
	db, err1 := l.Select(0)
	fmt.Println(err1)

	db.Set([]byte("setkey"),[]byte("1"))
	r, err2 := db.HSet([]byte(key), []byte("f1"), data)
	fmt.Println(r,err2)



}

func ComposeZxgHistoryData() dao.ZxgHisData {
	zxghisdata := dao.ZxgHisData{}

	GroupInfoList := make([]dao.ZxgHisGroupInfo, 0)
	GidSortDic := make(map[string]float64)

	GidStockInfoDic := make(map[string]map[string]dao.ZxgHisStockContent)
	GidStockSortDic := make(map[string]map[string]float64)
	GidStockNoteDic := make(map[string]map[string]dao.ZxgHisStockNote)
	GidStockStarDic := make(map[string][]string)

	for i := 0; i < 20; i ++ {
		zxgssdbgrpInfo := dao.ZxgHisGroupInfo{
			Name:        fmt.Sprintf("自选股%v", i),
			StockCount:  200,
			Version:     0,
			Gid:         int64(i),
			Source:      "test",
			UpdatedTime: time.Now().UnixNano()}
		GroupInfoList = append(GroupInfoList, zxgssdbgrpInfo)

		GidSortDic[fmt.Sprintf("%v", i)] = float64(i)

		subgidStockContentMap := make(map[string]dao.ZxgHisStockContent)
		subgidStockSortMap := make(map[string]float64)
		subgidStockNoteMap := make(map[string]dao.ZxgHisStockNote)
		stockstarSlice := make([]string, 0)
		for j := 0; j < 200; j ++ {

			zxgssdbstkcontent := dao.ZxgHisStockContent{
				IsStarStock:  false,
				Price:        "--",
				Timestick:    time.Now().UnixNano(),
				IsGivenStock: false}
			subgidStockContentMap[fmt.Sprintf("%v", j)] = zxgssdbstkcontent
			subgidStockSortMap[fmt.Sprintf("%v", j)] = float64(j)

			zxgssdbnote := dao.ZxgHisStockNote{
				StockCount:      "500",
				StockPrice:      "12",
				BrokerageFee:    "1",
				Factorage:       "2",
				StampDuty:       "0.2",
				TargetPrice:     "15",
				StopLossPrice:   "10",
				Date:            "20171015",
				Note:            "xxxxxxxx",
				TransactionDate: "20171015"}
			subgidStockNoteMap[fmt.Sprintf("%v", j)] = zxgssdbnote
			stockstarSlice = append(stockstarSlice, fmt.Sprintf("%v", j))
		}
		GidStockInfoDic[fmt.Sprintf("%v", i)] = subgidStockContentMap
		GidStockSortDic[fmt.Sprintf("%v", i)] = subgidStockSortMap
		GidStockNoteDic[fmt.Sprintf("%v", i)] = subgidStockNoteMap
		GidStockStarDic[fmt.Sprintf("%v", i)] = stockstarSlice
	}
	zxghisdata.DefGroupId = "0"
	zxghisdata.HGroupId = "0"
	zxghisdata.Date = "20171107"

	zxghisdata.GroupInfoList = GroupInfoList
	zxghisdata.GidSortDic = GidSortDic
	zxghisdata.GidStockInfoDic = GidStockInfoDic
	zxghisdata.GidStockSortDic = GidStockSortDic
	zxghisdata.GidStockNoteDic = GidStockNoteDic
	zxghisdata.GidStockStarDic = GidStockStarDic

	return zxghisdata
}

func TestLedisDb() {
	zxghistdata := ComposeZxgHistoryData()
	byteData, _ := json.Marshal(zxghistdata)
	SetHash("hashkey", byteData)
}

