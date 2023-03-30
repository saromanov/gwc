package core

// Config defines data for the config
type Config struct {
	Words bool
	Bytes bool
	Chars bool
	Lines bool
}

// isAllFalse returns  true if all options
// in config is false
func (c Config) isAllFalse() bool {
	return !c.Words && !c.Bytes && !c.Chars && !c.Lines
}
