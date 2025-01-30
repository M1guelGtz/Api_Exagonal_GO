package applicationusers

import domainusers "demob/src/users/domain_users"

type ViewUserByIdUseCase struct {
	repository domainusers.UserInterface
}

func NewViewUserUseCase(repository domainusers.UserInterface) *ViewUserByIdUseCase {
	return &ViewUserByIdUseCase{repository: repository}
}

func (uc *ViewUserByIdUseCase) Execute(userId int32) (*domainusers.User, error) {
	return uc.repository.GetUserById(userId)
}
