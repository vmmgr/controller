package v0

import (
	"fmt"
	"github.com/vmmgr/controller/pkg/api/core"
	region "github.com/vmmgr/controller/pkg/api/core/region"
	"github.com/vmmgr/controller/pkg/api/store"
	"gorm.io/gorm"
	"log"
	"time"
)

func Create(region *core.Region) (*core.Region, error) {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return region, fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return region, fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	defer dbSQL.Close()
	err = db.Create(&region).Error
	return region, err
}

func Delete(region *core.Region) error {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	defer dbSQL.Close()

	return db.Delete(region).Error
}

func Update(base int, data core.Region) error {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	defer dbSQL.Close()

	var result *gorm.DB
	if region.UpdateAll == base {
		result = db.Model(&core.Region{Model: gorm.Model{ID: data.ID}}).Updates(core.Region{
			Name:    data.Name,
			Comment: data.Comment,
			Lock:    data.Lock,
		})
	} else {
		log.Println("base select error")
		return fmt.Errorf("(%s)error: base select\n", time.Now())
	}
	return result.Error
}

func Get(base int, data *core.Region) region.ResultDatabase {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return region.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return region.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	defer dbSQL.Close()

	var regionStruct []core.Region

	if base == region.ID { //ID
		err = db.First(&regionStruct, data.ID).Error
	} else if base == region.Name {
		err = db.Where("name = ?", data.Name).Find(&regionStruct).Error
	} else {
		log.Println("base select error")
		return region.ResultDatabase{Err: fmt.Errorf("(%s)error: base select\n", time.Now())}
	}
	return region.ResultDatabase{Region: regionStruct, Err: err}
}

func GetAll() region.ResultDatabase {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return region.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return region.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	defer dbSQL.Close()

	var regions []core.Region
	err = db.Find(&regions).Error
	return region.ResultDatabase{Region: regions, Err: err}
}
