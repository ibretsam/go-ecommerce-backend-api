package settings

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	UserName        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
