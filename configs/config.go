package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
	"log"
)

type conf struct {
	DBDriver      string `map_structure:"DB_DRIVER"`
	DBHost        string `map_structure:"DB_HOST"`
	DBPort        string `map_structure:"DB_PORT"`
	DBName        string `map_structure:"DB_NAME"`
	DBUser        string `map_structure:"DB_USER"`
	DBPassword    string `map_structure:"DB_PASSWORD"`
	WebServerPort string `map_structure:"WEB_SERVER_PORT"`
	JWTSecret     string `map_structure:"JWT_SECRET"`
	JWTExpiresIn  int    `map_structure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (cfg *conf, err error) {
	CfgType := "env"

	viper.SetConfigName("app_config")
	viper.SetConfigType(CfgType)
	viper.AddConfigPath(path)
	viper.SetConfigFile("." + CfgType)

	// reescreve o .env pelas variáveis de ambiente do servidor caso existam.
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error to Read Config: ", err)
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("error to Unmarshal: ", err)
		return nil, err
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return
}
