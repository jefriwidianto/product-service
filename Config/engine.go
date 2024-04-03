package Config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"runtime"
)

func GetEnvironment(env string) (res ConfigSettingSql) {
	//process get file from environment
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), "../"+ENVIRONMENT_PATH+env+".yml")
	_, err := os.Stat(envPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//process read file Config environment "*.yml"
	content, err := os.ReadFile(envPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = yaml.Unmarshal(content, &res.Environment)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//load Config environment not error
	log.Println("Environment Config load successfully!")
	return
}
