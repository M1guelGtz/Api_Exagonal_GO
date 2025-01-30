package applicationusers

import domainusers "demob/src/users/domain_users"

type ViewAllUsersUseCase struct {
	repository domainusers.UserInterface
}

func NewViewAllusersUseCase(repository domainusers.UserInterface) *ViewAllUsersUseCase {
	return &ViewAllUsersUseCase{repository: repository}
}

func (uc *ViewAllUsersUseCase) Execute() ([]*domainusers.User, error) {
	return uc.repository.GetUsers()
}
