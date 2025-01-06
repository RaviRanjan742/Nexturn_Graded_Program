
package banking

import (
    "fmt"
    "time"
)

type Account struct {
    id          int
    name        string
    balance     float64
    transactions []string
}

func NewAccount(id int, name string, initialBalance float64) *Account {
    account := &Account{
        id:           id,
        name:         name,
        balance:      initialBalance,
        transactions: make([]string, 0),
    }
    account.addTransaction(fmt.Sprintf("Account created with initial balance of $%.2f", initialBalance))
    return account
}

func (a *Account) GetID() int {
    return a.id
}

func (a *Account) GetName() string {
    return a.name
}

func (a *Account) GetBalance() float64 {
    return a.balance
}

func (a *Account) Deposit(amount float64) error {
    if amount < MinimumDeposit {
        return fmt.Errorf("minimum deposit amount is $%.2f", MinimumDeposit)
    }

    a.balance += amount
    a.addTransaction(fmt.Sprintf("Deposited $%.2f", amount))
    return nil
}

func (a *Account) Withdraw(amount float64) error {
    if amount < MinimumDeposit {
        return fmt.Errorf("minimum withdrawal amount is $%.2f", MinimumDeposit)
    }

    if a.balance-amount < MinimumBalance {
        return fmt.Errorf("insufficient balance")
    }

    a.balance -= amount
    a.addTransaction(fmt.Sprintf("Withdrawn $%.2f", amount))
    return nil
}

func (a *Account) GetTransactionHistory() []string {
    return a.transactions
}

func (a *Account) addTransaction(description string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    transaction := fmt.Sprintf("[%s] %s - Balance: $%.2f", timestamp, description, a.balance)
    a.transactions = append(a.transactions, transaction)
}