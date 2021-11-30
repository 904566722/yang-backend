package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"
    "yang-backend/pkg/book_keeping/models"
    "yang-backend/pkg/config"
)

var DB *gorm.DB

func InitDB()  {
    dbLogLever := logger.Warn
    if config.Config.EnableDbLog {
        dbLogLever = logger.Info
    }
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second, // 慢 SQL 阈值
            LogLevel:      dbLogLever,  // Log level
            Colorful:      true,        // 彩色打印
        },
    )
    db, err := gorm.Open(mysql.Open(config.Config.DbPath), &gorm.Config{Logger: newLogger, PrepareStmt: true})
    if err != nil {
        log.Fatalf("gorm open db failed.\ndatabase path: %s\nerr:%v\n", config.Config.DbPath, err)
    }
    DB = db
}

// InitTable 初始化表
func InitTable() {
    err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;").AutoMigrate(
        &models.Income{},
        &models.Outcome{},
    )
    if err != nil {
        log.Fatalf("gorm initialize table failed\nerr:%v\n", err)
    }
}