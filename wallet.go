package wallet_project

import (
    "errors"
    "sync"
)

type Wallet struct {
    balance float64
    mutex   sync.Mutex
}

func NewWallet() *Wallet {
    return &Wallet{}
}

func (w *Wallet) Deposit(amount float64) {
    w.mutex.Lock()
    defer w.mutex.Unlock()
    w.balance += amount
}

func (w *Wallet) Withdraw(amount float64) error {
    w.mutex.Lock()
    defer w.mutex.Unlock()
    if w.balance < amount {
        return errors.New("insufficient funds")
    }
    w.balance -= amount
    return nil
}

func (w *Wallet) Balance() float64 {
    w.mutex.Lock()
    defer w.mutex.Unlock()
    return w.balance
}
