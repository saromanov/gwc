package core

// App defines main app
type App struct {
	cfg Config
}

func New(cfg Config) *App {
	return &App{
		cfg: cfg,
	}
}

// Run provides running of the app
func (a *App) Run(data []byte) error {
	r, err := countData(a.cfg, data)
	if err != nil {
		return err
	}
	output(r)
	return nil
}