package user

type ModuleUser struct{}

func (u *ModuleUser) GetName() string {
	return "user"
}

func (u *ModuleUser) Init() {}
