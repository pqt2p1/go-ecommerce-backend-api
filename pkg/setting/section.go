package setting

type Config struct{
	Mysql MySQLsetting `mapstructure:"mysql"`
}

type MySQLsetting struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBname string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`
}