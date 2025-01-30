package applicationusers

import (
	domainusers "demob/src/users/domain_users"
)

type UpdateUserUseCase struct {
	repository domainusers.UserInterface
}

func NewUpdateUserUseCase(repository domainusers.UserInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{repository: repository}
}

func (uc *UpdateUserUseCase) Execute(user *domainusers.User) error {
	return uc.repository.UpdateUser(user)
}
