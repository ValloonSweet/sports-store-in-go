package repo

func (repo *SqlRepository) Init() {
	if _, err := repo.Commands.Init.ExecContext(repo.Context); err != nil {
		repo.Logger.Panic("cannot exec init command")
	}
}

func (repo *SqlRepository) Seed() {
	if _, err := repo.Commands.Seed.ExecContext(repo.Context); err != nil {
		repo.Logger.Panic("cannot exec seed command")
	}
}
