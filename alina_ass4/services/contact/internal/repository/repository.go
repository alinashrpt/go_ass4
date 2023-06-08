package repository

import (
	"fmt"
	"alina.net/services/contact/internal/domain"
)

type ContactRepository interface {
	SaveContact(contact *domain.Contact) error
	GetContact(contactID int) (*domain.Contact, error)
	DeleteContact(contactID int) error
}

type GroupRepository interface {
	SaveGroup(group *domain.Group) error
	GetGroup(groupID int) (*domain.Group, error)
	AddContactToGroup(contact *domain.Contact, groupID int) error
}

type contactRepository struct {
	contacts map[int]*domain.Contact
}

type groupRepository struct {
	groups map[int]*domain.Group
	
}

func NewContactRepository() ContactRepository {
	return &contactRepository{
		contacts: make(map[int]*domain.Contact),
	}
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{
		groups: make(map[int]*domain.Group),
	}
}

func (r *contactRepository) SaveContact(contact *domain.Contact) error {
	r.contacts[contact.ID] = contact
	return nil
}

func (r *contactRepository) GetContact(contactID int) (*domain.Contact, error) {
	contact, ok := r.contacts[contactID]
	if !ok {
		return nil, fmt.Errorf("contact not found")
	}
	return contact, nil
}

func (r *contactRepository) DeleteContact(contactID int) error {
	_, ok := r.contacts[contactID]
	if !ok {
		return fmt.Errorf("contact not found")
	}
	delete(r.contacts, contactID)
	return nil
}

func (r *groupRepository) SaveGroup(group *domain.Group) error {
	r.groups[group.ID] = group
	return nil
}

func (r *groupRepository) GetGroup(groupID int) (*domain.Group, error) {
	group, ok := r.groups[groupID]
	if !ok {
		return nil, fmt.Errorf("group not found")
	}
	return group, nil
}

func (r *groupRepository) AddContactToGroup(contact *domain.Contact, groupID int) error {
	group, ok := r.groups[groupID]
	if !ok {
		return fmt.Errorf("group not found")
	}

	group.Contacts = append(group.Contacts, contact)
	return nil
}


