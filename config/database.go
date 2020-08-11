package config

// 数据库驱动
const DriverName = "mysql"

// 数据库配置
type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	Charset  string
	ShowLog  bool
	DebugLog bool
}

// 主库配置
var MasterDbConfig = DbConfig{
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "root",
	DbName:   "www_store_com",
	Charset:  "utf8mb4",
	ShowLog:  true,
	DebugLog: true,
}
