package model

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type BaseModel struct {

}

var db *gorm.DB

type MysqlConfig struct {
	Addr     string
	User string
	Passwd string
	DBName string
}

func init() {

	//读取配置文件的mysql
	viper.SetConfigName("config") // 配置文件的文件名，没有扩展名，如 .yaml, .toml 这样的扩展名
	viper.SetConfigType("yaml")  // 设置扩展名。在这里设置文件的扩展名。另外，如果配置文件的名称没有扩展名，则需要配置这个选项
	viper.AddConfigPath("./")             // 还可以在工作目录中搜索配置文件
	// 搜索并读取配置文件
	if err1:= viper.ReadInConfig(); err1!= nil { // 处理错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err1))
	}
	//fmt.Println(viper.GetInt(`mysql.port`))
	//fmt.Println(viper.Get("mysql.username"))
	//fmt.Println(viper.Get(`mysql.password`))
	//fmt.Println(viper.Get(`mysql.dbname`))
	//fmt.Println(viper.GetString(`mysql.code`))
	//fmt.Println(viper.GetString("mysql.code"))
	//fmt.Println(viper.GetString("mysql.code"))
	//fmt.Println(viper.Get("mysql.code.selected"))
	//fmt.Println(viper.Get(`mysql.code.companies`))
	//fmt.Println(viper.Get("switch"))

	//初始化数据库
	var dbConfig = mysql.Config{
		User: viper.GetString("mysql.username"),
		Passwd: viper.GetString("mysql.password"),
		DBName: viper.GetString("mysql.dbname"),
		Addr: viper.GetString("mysql.host"),
		Net: viper.GetString("mysql.net"),
		AllowNativePasswords: true,
	}
	dsn := dbConfig.FormatDSN()
	var err error
	db, err = gorm.Open("mysql",dsn)
	if err != nil{
		fmt.Println(err)
	}
}

func Close() {
	defer db.Close()
}
