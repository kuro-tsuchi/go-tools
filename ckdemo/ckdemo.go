package ckdemo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math/rand"
	"time"
)

var db *gorm.DB

func MainExec() {
	var err error
	dsn := "tcp://192.168.1.39:9000?database=test&read_timeout=10&write_timeout=20"
	db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Error(err.Error())
	}
	var count int64
	info := &OrderInfo{}
	if err = db.Model(info).Count(&count).Error; err != nil {
		log.Error(err.Error())
	}
	if count == 0 {
		db.AutoMigrate(info)
		initDB(db)
	}
	//log.Infof("total number of orders: %d", count)

	err = searchInfo()
	sumInfo()

	//sum(toDecimal32(money,2))
	type tempData struct {
		UserId int64           `json:"user_id"`
		Vip    string          `json:"vip"`
		Total  decimal.Decimal `json:"total"`
	}

	//var data []tempData
	//err = db.Raw("SELECT user_id , vip, sum(money) total from order_info group by  user_id , vip").Find(&data).Error
	rows, err := db.Model(info).Select("user_id , vip, sum(money) as total").Group("user_id , vip").Rows()
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(rows)
	dataList := make([]*tempData, 0)
	for rows.Next() {
		dest := &tempData{}
		db.ScanRows(rows, dest)
		//data := new(tempData)
		//rows.Scan(&data.UserId, &data.Vip, &data.Total)
		//dataList= append(dataList, data)
		fmt.Println(dest.Total)
	}
	fmt.Printf("%#v", &dataList)

}

func sumInfo() {
	//type Data struct {
	//	total float32
	//}
	//var data Data

	var total decimal.Decimal
	fmt.Println(total)
	rows, err := db.Raw("SELECT  sum(money) total from order_info").Rows()
	if err != nil {
		log.Error(err.Error)
	}
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			log.Error(err.Error())
		}

	}
		fmt.Printf("total %v",  total)
}

func searchInfo() error {
	info := OrderInfo{}
	if err := db.First(&info).Error; err != nil {
		log.Error(err.Error())
		return err
	}
	fmt.Printf("==================searchInfo:%#v\n", info.Money)
	return nil
}

func initDB(db *gorm.DB) {
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		param := rand.Int63() % 3
		err := db.CreateInBatches(&OrderInfo{
			Id:     uuid.New().String(),
			UserId: param,
			//Money:  decimal.NewFromFloatWithExponent(rand.Float64()*100, -2),
			Vip: fmt.Sprintf("VIP%d", param),
			//CreateTime: time.Now().UnixMicro() / 1e3,
		}, 10).Error
		if err != nil {
			log.Error(err.Error())
		}
	}
}

type OrderInfo struct {
	Id     string  `json:"id"`
	UserId int64   `json:"userId"`
	Money  decimal.Decimal `json:"money" sql:"type:decimal(18,2);"`
	Vip    string  `json:"vip"`
	//CreateTime int64  `json:"create_time"`
}
