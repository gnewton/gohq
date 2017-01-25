package main

import (
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func persist(report *OutageReport, db *gorm.DB) {

	db.Create(report)

	for i, _ := range report.Outages {
		outage := report.Outages[i]
		db.Create(outage)
	}
}

func dbOpen(dbFileName string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbFileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Opening db file: ", dbFileName)

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	sqlite3Config(db.DB())
	return db, nil
}

func sqlite3Config(db *sql.DB) {
	//db.Exec("PRAGMA auto_vacuum = 0;")
	//db.Exec("PRAGMA cache_size=32768;")
	db.Exec("PRAGMA cache_size=65536;")
	db.Exec("PRAGMA count_changes = OFF;")
	db.Exec("PRAGMA cache_spill = ON;")
	//db.Exec("PRAGMA journal_size_limit = 67110000;")
	db.Exec("PRAGMA locking_mode = EXCLUSIVE;")
	//db.Exec("PRAGMA locking_mode = OFF;")
	db.Exec("PRAGMA encoding = \"UTF-8\";")
	//db.Exec("PRAGMA journal_mode = WAL;")

	db.Exec("busy_timeout=0;")
	db.Exec("legacy_file_format=OFF;")

	//db.Exec("PRAGMA mmap_size=1099511627776;")
	db.Exec("PRAGMA page_size = 40960;")

	db.Exec("PRAGMA shrink_memory;")
	db.Exec("PRAGMA synchronous=OFF;")
	//db.Exec("PRAGMA synchronous = NORMAL;")
	//db.Exec("PRAGMA temp_store = MEMORY;")
	//db.Exec("PRAGMA threads = 5;")
	//db.Exec("PRAGMA wal_autocheckpoint = 1638400;")
}

func dbInit(dbFile string) (*gorm.DB, error) {
	db, err := dbOpen(dbFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("%v\n", *db)

	if !db.HasTable(&OutageReport{}) {
		db.CreateTable(&OutageReport{})
	}
	if !db.HasTable(&Outage{}) {
		db.CreateTable(&Outage{})
	}
	return db, nil
}

func containsStamp(stamp string, db *gorm.DB) bool {
	var count int
	db.Where("stamp = ?", stamp).Find(&OutageReport{}).Count(&count)

	return count > 0
}
