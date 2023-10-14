package main

type Config struct {
	Token        string `config:"TELEGRAM_BOT_TOKEN"`
	DatabaseUrl  string `config:"DB_URL"`
	DatabaseKey  string `config:"YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS"`
	DatabaseType string `config:"DB_TYPE"`
}
