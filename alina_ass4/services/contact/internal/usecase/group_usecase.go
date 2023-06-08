package usecase

import (

	"alina.net/services/contact/internal/domain"
	"alina.net/services/contact/internal/repository"
)

type GroupUseCase interface {
	CreateGroup(group *domain.Group) error
	ReadGroup(groupID int) (*domain.Group, error)
}

type groupUseCase struct {
	groupRepo repository.GroupRepository
	nextID    int
}

func NewGroupUseCase(groupRepo repository.GroupRepository) GroupUseCase {
	return &groupUseCase{
		groupRepo: groupRepo,
		nextID:    1,
	}
}

func (uc *groupUseCase) CreateGroup(group *domain.Group) error {
	err := uc.groupRepo.SaveGroup(group)
	if err != nil {
		return err
	}
	return nil
}

func (uc *groupUseCase) ReadGroup(groupID int) (*domain.Group, error) {
	group, err := uc.groupRepo.GetGroup(groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}
