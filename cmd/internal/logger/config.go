package logger

type Config struct {
	Level  string `default:"info"`
	Format string `default:"json"` // json | text
}
