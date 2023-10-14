package ydb

type Config struct {
	Database string
}

func (c *Config) ReturnDatabase() string {
	return c.Database
}
