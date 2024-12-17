package config

type Config struct {
	IP              string `yaml:"ip"`
	Port            int    `yaml:"port"`
	DB              db     `yaml:"db"`
	MailThemeMinLen int    `yaml:"mail_theme_min_len"`
	MailThemeMaxLen int    `yaml:"mail_theme_max_len"`
	MailTextMinLen  int    `yaml:"mail_text_min_len"`
	MailTextMaxLen  int    `yaml:"mail_text_max_len"`
}

type db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}
