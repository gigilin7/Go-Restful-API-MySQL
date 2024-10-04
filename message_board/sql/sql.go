package sql

import (
	"fmt"
	"os"

	"message/model"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connect *gorm.DB

type conf struct {
	Host     string `yaml:"host"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func (c *conf) getConf() (*conf, error) {
	// 讀取 config/connect.yaml 檔案
	yamlFile, err := os.ReadFile("sql/connect.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading yaml file: %w", err)
	}

	// 將讀取的字串轉換成結構體 conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling yaml file: %w", err)
	}
	return c, nil
}

// 初始化連線資料庫
func InitMySql() error {
	var c conf

	// 獲取 yaml 配置引數
	conf, err := c.getConf()
	if err != nil {
		return err
	}

	// 將 yaml 配置引數拼接成連線資料庫的 URL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DbName,
	)

	// 連線資料庫
	Connect, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}

	// 自動遷移，創建或更新資料表
	err = Connect.AutoMigrate(&model.Message{})
	if err != nil {
		return fmt.Errorf("error running AutoMigrate: %w", err)
	}

	return nil
}
