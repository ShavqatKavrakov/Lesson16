package wallet

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/ShavqatKavrakov/Lesson16_Export/pkg/types"
)

var ErrPhoneRegistered = errors.New("phone already registered")
var ErrAmountMustBePositive = errors.New("ammount must be greater than zero")
var ErrAccountNotFound = errors.New("account not found")

type Service struct {
	nextAccountId int64
	accounts      []*types.Account
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, acc := range s.accounts {
		if acc.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountId++
	account := &types.Account{
		ID:      s.nextAccountId,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}
func (s *Service) FindAccountById(accountID int64) (*types.Account, error) {
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			return acc, nil
		}
	}
	return nil, ErrAccountNotFound
}
func (s *Service) Deposit(accountID int64, amount types.Money) (*types.Account, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}
	account, err := s.FindAccountById(accountID)
	if err != nil {
		return nil, err
	}
	account.Balance += amount
	return account, nil
}
func (s *Service) ExportToFile(path string) error {
	var result string
	for _, acc := range s.accounts {
		result += strconv.Itoa(int(acc.ID)) + " " + string(acc.Phone) + " " + strconv.Itoa(int(acc.Balance)) + ";\n"
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if cel := file.Close(); cel != nil {
			log.Print(cel)
		}
	}()
	_, err = file.Write([]byte(result))
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) ImportToFile(path string)  error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if cel := file.Close(); cel != nil {
			log.Print(cel)
		}
	}()
	buf := make([]byte, 4096)
	read, err := file.Read(buf)
	if err != nil {
		return err
	}
	log.Print(string(buf[:read]))
	return nil
}
