package entity

import (
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/go-memdb"
)

// Entity struct
type Wallet struct {
    Id          uint    `form:"id" json:"id" bind:"required"`
    Balance     float64 `form:"balance" json:"balance"`
}

// Request and response objects
type BalanceMovementRequestObject struct {
    Id      int64   `form:"id" json:"id" binding:"required"`
    Amount  float64 `form:"amount" json:"amount" binding:"required"`
}

type BalanceParamParamObject struct {
    UserId  uint   `form:"id" json:"id" binding:"required"`
}

type WalletResponseObject struct {
    Wallet Wallet   `json:"wallet"`
    Errors []string `json:"errors"`
}

// Entity methods
func (w *Wallet) FindOneById(db *memdb.MemDB) *Wallet {
    txn := db.Txn(false)
    defer txn.Abort()

    raw, err := txn.First("wallet", "id", w.Id)
    if err != nil {
        panic(err)
    }
    w = raw.(*Wallet)

    return w
}

func (w *Wallet) FindOneByUserId(db *memdb.MemDB, userId uint) *Wallet {
    user := (&User{
        Id: userId,
    }).FindOneById(db)

    w.Id = user.WalletId
    w = w.FindOneById(db)

    return w
}

func (w *Wallet) Insert(db *memdb.MemDB) *Wallet {
    txn := db.Txn(true)
    if err := txn.Insert("wallet", w); err != nil {
		panic(err)
	}
    txn.Commit()

    return w
}

func (w *Wallet) Update(db *memdb.MemDB) *Wallet {
    txn := db.Txn(true)
    if err := txn.Insert("wallet", w); err != nil {
        panic(err)
    }
    txn.Commit()

    return w
}

func (w *Wallet) AugmentBalance(db *memdb.MemDB, bmr BalanceMovementRequestObject) (*Wallet, error) {
    w.Balance = w.Balance + bmr.Amount

    txn := db.Txn(true)
    if err := txn.Insert("wallet", w); err != nil {
        txn.Abort()
        return w, err
    }
    txn.Commit()

    return w, nil
}

func (w *Wallet) DeductBalance(db *memdb.MemDB, bmr BalanceMovementRequestObject) (*Wallet, error) {
    if w.Balance - bmr.Amount < 0 {
        return w, errors.New("Negative balance")
    }

    w.Balance = w.Balance - bmr.Amount
    txn := db.Txn(true)
    if err := txn.Insert("wallet", w); err != nil {
        txn.Abort()
        return w, err
    }
    txn.Commit()

    return w, nil
}

func (w *Wallet) BuildResponse(e []string) (response WalletResponseObject) {
    response.Wallet = *w
    for _, err := range e {
        log.Println(fmt.Sprintf("[Application Log] %s", err))
        response.Errors = append(response.Errors, err)
    }

    return
}

