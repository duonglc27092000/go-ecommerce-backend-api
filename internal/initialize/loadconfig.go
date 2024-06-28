package initialize

import (
	"fmt"

	"github.com/duonglc27092000/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath(("./config/")) //path to config
	viper.SetConfigName("local")       // ten file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	//read server configuration
	fmt.Println("Server port ::: ", viper.GetInt("server.port"))
	fmt.Println("Key ::: ", viper.GetString("security.jwt.key"))
	// configure  structure

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode configuration %w", err)
	}
}
