package usecase

import (

	"alina.net/services/contact/internal/domain"
	"alina.net/services/contact/internal/repository"
)

type ContactUseCase interface {
	CreateContact(contact *domain.Contact) error
	ReadContact(contactID int) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) error
	DeleteContact(contactID int) error

	CreateGroup(group *domain.Group) error
	ReadGroup(groupID int) (*domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}

type contactUseCase struct {
	contactRepo repository.ContactRepository
	groupRepo   repository.GroupRepository
	nextID    int
}

func NewContactUseCase(contactRepo repository.ContactRepository, groupRepo repository.GroupRepository) ContactUseCase {
	return &contactUseCase{
		contactRepo: contactRepo,
		groupRepo:   groupRepo,
		nextID:    1,
	}
}

func (uc *contactUseCase) CreateContact(contact *domain.Contact) error {
	contact.ID = uc.nextID
	err := uc.contactRepo.SaveContact(contact)
	if err != nil {
		return err
	}
	uc.nextID++
	return nil
}

func (uc *contactUseCase) ReadContact(contactID int) (*domain.Contact, error) {
	contact, err := uc.contactRepo.GetContact(contactID)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (uc *contactUseCase) UpdateContact(contact *domain.Contact) error {
	err := uc.contactRepo.SaveContact(contact)
	if err != nil {
		return err
	}
	return nil
}

func (uc *contactUseCase) DeleteContact(contactID int) error {
	err := uc.contactRepo.DeleteContact(contactID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *contactUseCase) CreateGroup(group *domain.Group) error {
	group.ID = uc.nextID
	err := uc.groupRepo.SaveGroup(group)
	if err != nil {
		return err
	}
	uc.nextID++
	return nil
}

func (uc *contactUseCase) ReadGroup(groupID int) (*domain.Group, error) {
	group, err := uc.groupRepo.GetGroup(groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *contactUseCase) AddContactToGroup(contactID, groupID int) error {
	contact, err := uc.contactRepo.GetContact(contactID)
	if err != nil {
		return err
	}

	err = uc.groupRepo.AddContactToGroup(contact, groupID)
	if err != nil {
		return err
	}

	return nil
}
