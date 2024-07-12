package accounts

import (
	"awesomeProject/accounts/models"
	"awesomeProject/proto"

	"sync"
	"errors"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(req *proto.CreateAccountRequest) error {
	if req.Name == "" {
		return errors.New("name cannot be empty")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; ok {
		h.guard.Unlock()
		return errors.New("account already exists")
	}
	h.accounts[req.Name] = &models.Account{
		Name: req.Name,
		Amount: req.Amount,
	}
	h.guard.Unlock()
	return nil
}

func (h *Handler) DeleteAccount(req *proto.DeleteAccountRequest) error {
	if req.Name == "" {
		return errors.New("name cannot be empty")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return errors.New("account not found")
	}
	delete(h.accounts, req.Name)
	h.guard.Unlock()
	return nil
}

func (h *Handler) ChangeAccount(req *proto.ChangeAccountRequest) error {
	if req.Name == "" {
		return errors.New("name cannot be empty")
	}
	if req.NewName == "" {
		return errors.New("new name cannot be empty")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return errors.New("account not found")
	}
	if _, ok := h.accounts[req.NewName]; ok {
		h.guard.Unlock()
		return errors.New("account with this name already exists")
	}

	prevAmount := h.accounts[req.Name].Amount

	delete(h.accounts, req.Name)
	h.accounts[req.NewName] = &models.Account{
		Name: req.NewName,
		Amount: prevAmount,
	}

	h.guard.Unlock()
	return nil
}

func (h *Handler) PatchAccount(req *proto.PatchAccountRequest) error {
	if req.Name == "" {
		return errors.New("name cannot be empty")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return errors.New("account not found")
	}
	h.accounts[req.Name].Amount = req.NewAmount
	h.guard.Unlock()
	return nil
}

func (h *Handler) GetAccount(req *proto.GetAccountRequest) (*models.Account, error) {
	if req.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	h.guard.RLock()
	account, ok := h.accounts[req.Name]
	h.guard.RUnlock()
	if !ok {
		return nil, errors.New("account not found")
	}
	return account, nil
}
