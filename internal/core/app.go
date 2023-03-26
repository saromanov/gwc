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

func (a *App) Run() error {
	return nil
}