package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n===========================", sql)
}

var db *gorm.DB

func main() {
	dsn := "root:123456@tcp(localhost:3306)/goprogramming?charset=utf8mb4&parseTime=True&loc=Local"
	dial := mysql.Open(dsn)
	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(Gender{}, Test{}, Customer{})
	// CreateGender("testfa")

	// GetGenderByName("male")
	// UpdateGender2(1, "Male")
	// DeleteGender(4)

	// db.Migrator().CreateTable(Customer{})

	CreateCustomer("Bob", 1)
	CreateCustomer("Bab", 2)
	GetCustomers()

}

func GetCustomers() {
	customers := []Customer{}
	// tx := db.Preload("Gender").Order("id").Find(&customers)
	tx := db.Preload(clause.Associations).Order("id").Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(customers)
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{Name: name, GenderID: genderID}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
}

func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderByName(name string) {
	gender := Gender{}
	// tx := db.Where("name=?", name).Find(&gender)
	tx := db.First(&gender, "name=?", name)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

type Gender struct {
	ID   int
	Name string `gorm:"unique"`
}

type Test struct {
	gorm.Model
	Code uint
	Name string `gorm:"unique;not null"`
}
