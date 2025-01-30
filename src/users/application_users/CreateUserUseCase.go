package applicationusers

import (
	domainusers "demob/src/users/domain_users"
)

type CreateUserUseCase struct {
	repository domainusers.UserInterface
}

func NewCreateUserUseCase(repository domainusers.UserInterface) *CreateUserUseCase {
	return &CreateUserUseCase{repository: repository}
}

func (uc *CreateUserUseCase) Execute(name string, email string, password string) error {
	user := domainusers.NewUser(name, email, password)
	return uc.repository.Create(user)
}
