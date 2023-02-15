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
	RunMode   string
	Port      int
	LogOutput string
}
type JWT struct {
	Secret string
}
