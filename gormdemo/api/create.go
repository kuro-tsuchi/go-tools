package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	mydb "goTool/gormdemo/db"
	"goTool/gormdemo/modules"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

var db *gorm.DB

func Create() {
	p := GenerateProduct()
	result := db.Create(p)
	fmt.Println(p.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

func CreateInBatches() {
	var product *modules.Product
	var productList []*modules.Product
	num := 10
	for i := 0; i < num; i++ {
		product = GenerateProduct()
		productList = append(productList, product)
	}
	err := db.CreateInBatches(productList, 100).Error
	if err != nil {
		log.Error(err.Error())
	}
}

func init() {
	db = mydb.InitDB()
	product := &modules.Product{}
	db.AutoMigrate(product)
}

func GenerateProduct() *modules.Product {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	df := "2006-01-02 15:04:05"
	location, _ := time.LoadLocation("Local")

	//fakerSchema := &modules.FakerSchema{}
	//faker.FakeData(fakerSchema)

	p := &modules.Product{
		Model: modules.Model{
			ID:        uuid.New().String(),
			CreatedAt: time.Now().In(location).Format(df),
			UpdatedAt: time.Now().In(location).Format(df),
		},
		Code: rand.Int63n(1000),
		//Price: decimal.NewFromFloatWithExponent(fakerSchema.Amount, -2),
		Price: decimal.NewFromFloatWithExponent(rand.ExpFloat64()*1e3, -2),
	}

	return p
}
