package main

import (
	"chatbox-app/config"
	"chatbox-app/dao"
	"chatbox-app/lib/sloger"
	"chatbox-app/lib/token"
	"chatbox-app/lib/validate"
	"chatbox-app/models"
	"chatbox-app/routes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "chatbox-app/docs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const addr = ":4321"

type dbConfig struct {
	dbUser, dbPswd, dbHost, dbPort string
}

// @title           ChatBox API
// @version         0.01
// @description     这是chatbox api swagger文档.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4321
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	var cfg config.Settings
	err := config.LoadConfig("config.yaml", &cfg, ".")
	if err != nil {
		log.Fatal(err)
	}

	// 设置并启用全局 slog
	if cfg.ApiEnv == config.Pro {
		sloger.NewLogger(cfg.LogLevel, sloger.DefaultOutput(cfg.LogOutput))
	} else {
		sloger.NewLogger(cfg.LogLevel, os.Stdout)
	}

	jwtMaker, err := token.NewJWTMaker(cfg.JWTSecretKey, "ChatBox")
	if err != nil {
		log.Fatal(err)
	}

	trans := validate.NewValidation("zh")

	// 设置到全局使用
	config.SetupGlobalApp(cfg, jwtMaker, trans)

	// 从命令行中解析数连接据库参数
	conf := parseFlags()
	db := connMySQLDb(conf)
	dbMiragation(db)
	dao.DB = db // 设置到dao包全局使用
	router := routes.NewApiRoutes()
	srv := http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		MaxHeaderBytes: 2 << 20,
	}

	log.Println("start server on ", addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("srv.ListenAndServe err: ", err)
	}
}

func connMySQLDb(conf dbConfig) *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/db_chatbox?charset=utf8mb4&parseTime=True&loc=Local",
		conf.dbUser,
		conf.dbPswd,
		conf.dbHost,
		conf.dbPort,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(fmt.Sprintf("连接数据库失败：%v", err))
	}
	// 启用错误类型转换，将mysql错误转换成gorm error
	db.TranslateError = true
	return db
}

func parseFlags() dbConfig {
	var dbConf dbConfig

	flag.StringVar(&dbConf.dbUser, "dbUser", "", "数据库用户名")
	flag.StringVar(&dbConf.dbPswd, "dbPswd", "", "数据库密码")
	flag.StringVar(&dbConf.dbHost, "dbHost", "127.0.0.1", "数据库主机ip")
	flag.StringVar(&dbConf.dbPort, "dbPort", "3306", "数据库端口")
	flag.Parse()
	checkBaseVar(dbConf.dbUser, "数据库用户名不能为空")
	checkBaseVar(dbConf.dbPswd, "数据库密码不能为空")
	return dbConf
}

func checkBaseVar(v string, msg string) {
	if v == "" {
		panic(msg)
	}
}

func dbMiragation(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.VerifyEmail{})
	if err != nil {
		panic(err)
	}
}
