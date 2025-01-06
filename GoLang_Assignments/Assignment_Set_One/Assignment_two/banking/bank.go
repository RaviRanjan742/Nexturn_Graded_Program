package banking

import (
    "fmt"
)

type Bank struct {
    accounts []*Account
}

func NewBank() *Bank {
    return &Bank{
        accounts: make([]*Account, 0),
    }
}

func (b *Bank) AddAccount(account *Account) error {
    
    for _, acc := range b.accounts {
        if acc.GetID() == account.GetID() {
            return fmt.Errorf("account with ID %d already exists", account.GetID())
        }
    }

    b.accounts = append(b.accounts, account)
    return nil
}

func (b *Bank) findAccount(id int) (*Account, error) {
    for _, account := range b.accounts {
        if account.GetID() == id {
            return account, nil
        }
    }
    return nil, fmt.Errorf("account with ID %d not found", id)
}

func (b *Bank) Deposit(id int, amount float64) error {
    account, err := b.findAccount(id)
    if err != nil {
        return err
    }
    return account.Deposit(amount)
}

func (b *Bank) Withdraw(id int, amount float64) error {
    account, err := b.findAccount(id)
    if err != nil {
        return err
    }
    return account.Withdraw(amount)
}

func (b *Bank) GetBalance(id int) (float64, error) {
    account, err := b.findAccount(id)
    if err != nil {
        return 0, err
    }
    return account.GetBalance(), nil
}

func (b *Bank) GetTransactionHistory(id int) ([]string, error) {
    account, err := b.findAccount(id)
    if err != nil {
        return nil, err
    }
    return account.GetTransactionHistory(), nil
}

func (b *Bank) ListAccounts() []*Account {
    return b.accounts
}