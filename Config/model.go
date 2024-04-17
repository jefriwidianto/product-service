package Config

const (
	ENVIRONMENT_PATH = "/Environment/"
	Localhost        = "Local"
	PathMigration    = "/Services/Migration/"
)

type DbSqlConfigName string

const (
	// Database Connection Constant name
	DATABASE_MAIN DbSqlConfigName = "DBmain"
)

// ftroct for collect data object Config environment ".yml"
type Environment struct {
	App       app      `yaml:"apps"`
	Databases database `yaml:"databases"`
	Jwt       Jwt      `yaml:"jwt"`
	Grpc      Grpc     `yaml:"grpc"`
}

type database struct {
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Port               string `yaml:"port"`
	Engine             string `yaml:"engine"`
	Host               string `yaml:"host"`
	Maximum_connection int    `yaml:"maximum_connection"`
}

type app struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Jwt struct {
	SecretKey string `yaml:"secret_key"`
	Encrypt   string `yaml:"encrypt"`
}

type Grpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
