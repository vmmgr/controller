package v0

import (
	"fmt"
	"github.com/vmmgr/controller/pkg/api/core"
	"github.com/vmmgr/controller/pkg/api/core/node/storage"
	"github.com/vmmgr/controller/pkg/api/store"
	"gorm.io/gorm"
	"log"
	"time"
)

func Create(storage *core.Storage) (*core.Storage, error) {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return storage, fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return storage, fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())
	}
	defer dbSQL.Close()

	err = db.Create(&storage).Error
	return storage, err
}

func Delete(storage *core.Storage) error {
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

	return db.Delete(storage).Error
}

func Update(base int, data core.Storage) error {
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
	if storage.UpdateAll == base {
		result = db.Model(&core.Storage{Model: gorm.Model{ID: data.ID}}).Updates(core.Storage{
			NodeID:      data.NodeID,
			AdminOnly:   data.AdminOnly,
			Type:        data.Type,
			Path:        data.Path,
			MaxCapacity: data.MaxCapacity,
			Comment:     data.Comment,
		})
	} else {
		log.Println("base select error")
		return fmt.Errorf("(%s)error: base select\n", time.Now())
	}
	return result.Error
}

func Get(base int, data *core.Storage) storage.ResultDatabase {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	defer dbSQL.Close()

	var storageStruct []core.Storage

	if base == storage.ID { //ID
		err = db.First(&storageStruct, data.ID).Error
	} else if base == storage.NodeID { //Node?????????Storage??????
		err = db.Where("node_id = ?", data.NodeID).Find(&storageStruct).Error
	} else if base == storage.AdminOnly { //Node StorageID ???VMID????????????
		err = db.Where("admin_only = ?", data.AdminOnly).Find(&storageStruct).Error
	} else if base == storage.Name { //Name
		err = db.Where("lock = ?", data.Name).Find(&storageStruct).Error
	} else {
		log.Println("base select error")
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: base select\n", time.Now())}
	}
	return storage.ResultDatabase{Storage: storageStruct, Err: err}
}

func GetAll() storage.ResultDatabase {
	db, err := store.ConnectDB()
	if err != nil {
		log.Println("database connection error")
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	dbSQL, err := db.DB()
	if err != nil {
		log.Printf("database error: %v", err)
		return storage.ResultDatabase{Err: fmt.Errorf("(%s)error: %s\n", time.Now(), err.Error())}
	}
	defer dbSQL.Close()

	var storages []core.Storage
	err = db.Preload("Node").Find(&storages).Error
	return storage.ResultDatabase{Storage: storages, Err: err}
}
