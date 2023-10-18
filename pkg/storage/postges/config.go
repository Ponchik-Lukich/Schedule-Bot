package postges

type Config struct {
	DSN string
}

func (c *Config) ReturnDatabase() string {
	return c.DSN
}
