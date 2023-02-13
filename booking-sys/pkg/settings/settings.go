package settings

type AppConfig struct {
	Database
	Server
	JWT
}

type Database struct {
	DBDriver string `mapstructure:"DBDriver"`
	DBSource string `mapstructure:"DBSource"`
}
type Server struct {
	Port int
}
type JWT struct {
	Secret string
}
