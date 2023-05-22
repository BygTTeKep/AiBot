package main

import (
	"context"
	"log"

	aibot "github.com/g91TeJl/AiBot"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	// cfg := repository.Config{
	// 	Host:    viper.GetString("db.host"),
	// 	Port:    viper.GetString("db.port"),
	// 	User:    viper.GetString("db.user"),
	// 	Pass:    viper.GetString("db.password"), //os.getEnv (godot)
	// 	DBName:  viper.GetString("db.dbname"),
	// 	SSLmode: viper.GetString("db.sslmode"),
	// }
	// db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Pass))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// repo := repository.NewRepository(db)
	// service := service.NewService(repo)
	api := viper.GetString("api")
	apiKey := viper.GetString("STABILITY_API_KEY")
	//time.Sleep(1 * time.Minute)
	//endpoint.AuthStabilityAi(apiKey)
	ctx := context.Background()
	aibot.TelegramBot(ctx, api, apiKey)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
