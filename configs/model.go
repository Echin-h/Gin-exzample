package configs

import "github.com/spf13/viper"

var (
	DbSettings  *DatabaseSettings
	JwtSettings *JWTSettings
)

type Setting struct {
	vp *viper.Viper
}

type DatabaseSettings struct {
	Root      string
	Password  string
	Host      string
	Port      int
	Dbname    string
	Charset   string
	ParseTime string
	Loc       string
}

type JWTSettings struct {
	Issuer    string
	Subject   string
	SecretKey string
}
