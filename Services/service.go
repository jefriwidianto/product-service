package Services

import (
	"flag"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"product-service/Config"
	logger "product-service/Logger"
	"product-service/Routes"
	"product-service/Services/Jwt"
)

var AppEnv = flag.String("env", "", "define environment")

func init() {
	flag.Parse()
	if *AppEnv == "" {
		*AppEnv = Config.Localhost
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func AppInitialization() {
	config := Config.GetEnvironment(*AppEnv)
	config.InitDB()

	var validate = echo.New()
	validate.Validator = &CustomValidator{validator: validator.New()}

	service := serviceInit(&config)
	Jwt.JwtConfigValue = service.Jwt

	log, err := logger.New("info", "json", *AppEnv)
	if err != nil {
		panic(err)
	}

	//Collect Routes
	var routes Routes.Routes
	routes.Log = log
	routes.Controller.Log = log
	routes.CollectRoutes(validate, config.App.Host, config.App.Port)
}

func serviceInit(Env *Config.ConfigSettingSql) service {
	serv := service{
		Jwt: Jwt.JwtConfig{Config: &Env.Jwt},
	}

	return serv
}
