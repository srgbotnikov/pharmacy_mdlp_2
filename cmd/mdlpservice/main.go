package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	p "ci.drugs.main/okit/pharmacy_mdlp_2"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/cache"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/handler"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/mdlpapi"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/repository"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/service"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/sign"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configDir string
)

func init() {
	flag.StringVar(&configDir, "config-dir", "./configs", "path to config directory")
}

func main() {
	flag.Parse()

	if err := initConfig(configDir); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	api := mdlpapi.NewMdlpAPI(mdlpapi.Config{
		Url:          viper.GetString("api.Url"),          
		ClientID:     viper.GetString("api.ClientID"),     
		ClientSecret: viper.GetString("api.ClientSecret"), 
		UserID:       viper.GetString("api.UserID"),       
		AuthType:     viper.GetString("api.AuthType"),     //"SIGNED_CODE",
		UrlMdlp:      viper.GetString("api.UrlMdlp"),      //"https://api.sb.mdlp.crpt.ru:443",
	})
	logrus.Println(viper.GetString("port"))
	rdb, err := cache.NewRedisDB(cache.ConfigRedis{
		Host: viper.GetString("redis.Host"), //"localhost",
		Port: viper.GetString("redis.Port"), //"6379",
	})
	if err != nil {
		logrus.Fatalf("field to initialized Redis: %s", err.Error())
	}
	cache := cache.NewCache(rdb)

	sign := sign.NewSign(sign.Config{
		Thumbprint: viper.GetString("sign.Thumbprint"), 
		SignShPath: viper.GetString("sign.SignShPath"), 
	})
	services := service.NewService(repos, api, sign, cache)
	handlers := handler.NewHandler(services, viper.GetString("logstash"))

	srv := new(p.Server)
	// if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
	// 	log.Fatalf("error ocured while running http server: %s", err.Error())
	// }
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}

func initConfig(configDir string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
