package entity

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// export DB_DSN
//  export USER_DB_DNS="dbuser:dbpass@tcp(dbhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func TestTCloudAgencyEntityInit(t *testing.T) {
	// TCloud_AGENCY_DB_DNS
	dsn, ok := os.LookupEnv("USER_DB_DNS")
	if !ok {
		t.Log("DB DNS not set")
		return
	}
	t.Logf("FORECAST_DB_DNS=%s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// init tables
	initTables := []schema.Tabler{
		&TUser{},
	}

	// drop table
	for _, table := range initTables {
		tableName := table.TableName()
		// drop table
		tx := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
		if err := tx.Error; err != nil {
			t.Errorf("drop table[%s] got err: %s", tableName, err)
			continue
		}

		// auto migrate
		if err := db.AutoMigrate(table); err != nil {
			t.Errorf("auto migrate table[%s] got err: %s", tableName, err)
			continue
		}

		// log
		t.Logf("init table[%s] success", tableName)
	}

	t.Log("init success!")
}
