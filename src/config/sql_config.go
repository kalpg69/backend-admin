package config

type SQLConfig struct {
	UserName     string
	Password     string
	ServerName   string
	DatabaseName string
	Timeout      string
}

func GetSQLConfig() *SQLConfig {
	return &SQLConfig{
		UserName:     "sa",
		Password:     "itas@123",
		ServerName:   "localhost",
		DatabaseName: "blackduck",
		Timeout:      "30",
	}
}
