package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// dsn := "bridge:!WxY2P!Y^nVff#C@tcp(ec2-16-162-106-158.ap-east-1.compute.amazonaws.com:7006)/bridgedb?charset=utf8mb4&parseTime=True&loc=Local"

// BridgeOrder 表示跨链订单的模型
type BridgeOrder struct {
	Id                int64      `orm:"id,primary"`          //
	SourceAddress     string     `orm:"source_address"`      // 转出币种
	TargetAddress     string     `orm:"target_address"`      // 转入币种
	SourceCoinAddress string     `orm:"source_coin_address"` // 转出钱包地址
	TargetCoinAddress string     `orm:"target_coin_address"` // 转入钱包地址
	SourceChainId     string     `orm:"source_chainId"`      // 转出链
	TargetChainId     string     `orm:"target_chainId"`      // 转入链
	TransactionHash   string     `orm:"transaction_hash"`    // 交易哈希
	OrderId           string     `orm:"order_id"`            // 跨链订单id
	Amount            float64    `orm:"amount"`              // 数量
	Status            int        `orm:"status"`              // 跨链记录状态
	CreateAt          *time.Time `orm:"create_at"`           // 发起时间
	UpdateAt          *time.Time `orm:"update_at"`           // 结束时间
	BlockNumber       int64      `orm:"block_number"`        // 区块高度
	Fee               float64    `orm:"fee"`                 // 手续费
	Remark            string     `orm:"remark"`              // 评论
}

func (BridgeOrder) TableName() string {
	return "bridge_order"
}

type Coin struct {
	Id        int64      `orm:"id,primary" json:"id"`      //
	Name      string     `orm:"name" json:"name"`          // 代币名称
	Symbol    string     `orm:"symbol" json:"symbol"`      // 代币简称
	ChainId   string     `orm:"chain_id" json:"chainId"`   // 链Id
	Address   string     `orm:"address" json:"address"`    // 代币地址
	IsEnable  int        `orm:"isEnable" json:"isEnable"`  // 是否上架
	CreateAt  *time.Time `orm:"create_at" json:"createAt"` //
	UpdateAt  *time.Time `orm:"update_at" json:"updateAt"` //
	TokenType string     `orm:"type" json:"tokenType"`     // 币种类型
	Decimals  int        `orm:"decimals" json:"decimals"`  // 精度
	Icon      string     `orm:"icon" json:"icon"`          // 代币图标
}

func (Coin) TableName() string {
	return "coin"
}

// DailyBridgeStats 表示每日跨链统计结果的模型
type DailyBridgeStats struct {
	ID        int64     `gorm:"primaryKey"`
	Coin      string    `json:"coin"`      // 币种 (SourceAddress)
	InAmount  float64   `json:"inAmount"`  // 跨链转入
	OutAmount float64   `json:"outAmount"` // 跨链转出
	Balance   float64   `json:"balance"`   // 跨链差额
	Fee       float64   `json:"fee"`       // 跨链手续费
	StatDate  time.Time `json:"statDate"`  // 统计日期
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}

func (DailyBridgeStats) TableName() string {
	return "daily_bridge_stats"
}

//func main() {
//	// Database connection string
//	dsn := "bridge:!WxY2P!Y^nVff#C@tcp(ec2-16-162-106-158.ap-east-1.compute.amazonaws.com:7006)/bridgedb"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Define the time range
//	startOfDay := "2023-06-07 00:00:00"
//	endOfDay := "2023-06-08 00:00:00"
//
//	// Execute the SQL query
//	rows, err := db.Query(`SELECT id, source_address, target_address, source_coin_address, target_coin_address,
//                                  source_chainId, target_chainId, transaction_hash, orderId, amount, status,
//                                  create_at, update_at, block_number, fee, remark
//                           FROM bridge_order
//                           WHERE status = ? AND update_at >= ? AND update_at < ?`, 5, startOfDay, endOfDay)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	var orders []BridgeOrder
//
//	// Iterate over the result set
//	for rows.Next() {
//		var order BridgeOrder
//		err := rows.Scan(&order.Id, &order.SourceAddress, &order.TargetAddress, &order.SourceCoinAddress,
//			&order.TargetCoinAddress, &order.SourceChainId, &order.TargetChainId, &order.TransactionHash,
//			&order.OrderId, &order.Amount, &order.Status, &order.CreateAt, &order.UpdateAt, &order.BlockNumber,
//			&order.Fee, &order.Remark)
//		if err != nil {
//			log.Fatal(err)
//		}
//		orders = append(orders, order)
//	}
//
//	// Check for errors from iterating over rows
//	if err = rows.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//	// Print the retrieved orders
//	for _, order := range orders {
//		fmt.Printf("%+v\n", order)
//	}
//}

func main() {
	// 连接数据库
	dsn := "bridge:!WxY2P!Y^nVff#C@tcp(ec2-16-162-106-158.ap-east-1.compute.amazonaws.com:7006)/bridgedb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//getCoin(db)
	// 启动定时任务
	go scheduleDailyTask(db, "2024-06-07")

	// 阻止程序退出
	select {}
}

func scheduleDailyTask(db *gorm.DB, startDate string) {
	layout := "2006-01-02"
	startTime, err := time.Parse(layout, startDate)
	if err != nil {
		panic("Invalid start date format. Please use YYYY-MM-DD")
	}

	for {
		now := time.Now()
		if now.After(startTime) {
			calculateDailyStatistics(db, startTime)
			startTime = startTime.AddDate(0, 0, 1)
		}

		// 计算距离下一个凌晨的时间
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		duration := next.Sub(now)

		// 等待到下一个凌晨
		time.Sleep(duration)
	}
}

func getCoin(db *gorm.DB) {
	// 查询当天完成的跨链订单
	var coins []Coin
	db.Find(&coins)
	fmt.Println("hee")
	fmt.Println(coins)
}

func calculateDailyStatistics(db *gorm.DB, statDate time.Time) {
	startOfDay := time.Date(statDate.Year(), statDate.Month(), statDate.Day(), 0, 0, 0, 0, statDate.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	// 查询当天完成的跨链订单
	var orders []BridgeOrder
	db.Where("status = ? AND update_at >= ? AND update_at < ?", 5, startOfDay, endOfDay).Find(&orders)

	coinStats := make(map[string]map[string]float64)

	for _, order := range orders {
		coin := order.SourceCoinAddress
		amountAfterFee := order.Amount - order.Fee

		if coinStats[coin] == nil {
			coinStats[coin] = make(map[string]float64)
		}

		// 跨链转出
		coinStats[order.SourceAddress]["out"] += amountAfterFee
		// 跨链转入
		coinStats[order.TargetAddress]["in"] += amountAfterFee
		// 跨链手续费
		coinStats[order.SourceAddress]["fee"] += order.Fee
	}

	// 将统计结果存储到数据库中
	for coin, stats := range coinStats {
		in := stats["in"]
		out := stats["out"]
		fee := stats["fee"]
		balance := in - out

		chainStat := DailyBridgeStats{
			Coin:      coin,
			InAmount:  in,
			OutAmount: out,
			Balance:   balance,
			Fee:       fee,
			StatDate:  startOfDay,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// 保存或更新统计结果
		db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&chainStat)

		fmt.Printf("币种 %s 的统计数据已保存：\n", coin)
		fmt.Printf("  跨链转入: %.2f\n", in)
		fmt.Printf("  跨链转出: %.2f\n", out)
		fmt.Printf("  跨链差额: %.2f\n", balance)
		fmt.Printf("  跨链手续费: %.2f\n", fee)
	}
}
