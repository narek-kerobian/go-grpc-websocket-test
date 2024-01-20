package entity

import "github.com/hashicorp/go-memdb"

// Entity struct
type User struct {
    Id          uint `json:"id"`
    Name        string `json:"name"`
    WalletId    uint `jsno:"walletId"`
}

// Entity methods
func (u *User) FindOneById(db *memdb.MemDB) *User {
    txn := db.Txn(false)
    defer txn.Abort()

    raw, err := txn.First("user", "id", u.Id)
    if err != nil {
        panic(err)
    }
    u = raw.(*User)

    return u
}

func (u *User) Insert(db *memdb.MemDB) *User {
    txn := db.Txn(true)
    if err := txn.Insert("user", u); err != nil {
        panic(err)
    }
    txn.Commit()

    return u
}

func (u *User) Update(db *memdb.MemDB) *User {
    txn := db.Txn(true)
    if err := txn.Insert("user", u); err != nil {
        panic(err)
    }
    txn.Commit()

    return u
}
