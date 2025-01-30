package applicationusers

import domainusers "demob/src/users/domain_users"

type DeleteUserUseCase struct {
	repository domainusers.UserInterface
}

func NewDeleteUserUseCase(repository domainusers.UserInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{repository: repository}
}

func (uc *DeleteUserUseCase) Execute(productId int32) error {
	return uc.repository.DeleteUser(productId)
}
