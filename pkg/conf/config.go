package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type Mysql struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var (
	sql        *Mysql
	PageSize   int
	ServerAddr string
)

func Init() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./pkg/conf")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			fmt.Println("config file was found but another error was produced")
		}
	}

	ServerAddr = viper.GetStringMapString("server")["addr"]

	PageSize, _ = strconv.Atoi(viper.GetStringMapString("page")["page-size"])

	sql = LoadMysql(viper.GetStringMapString("mysql"))

	path := strings.Join([]string{sql.user, ":", sql.password, "@tcp(", sql.host, ":", sql.port, ")/", sql.dbname, "?charset=utf8mb4&parseTime=True"}, "")
	return path
}

func LoadMysql(myConf map[string]string) *Mysql {
	return &Mysql{
		host:     myConf["host"],
		port:     myConf["port"],
		user:     myConf["user"],
		password: myConf["password"],
		dbname:   myConf["dbname"],
	}
}
