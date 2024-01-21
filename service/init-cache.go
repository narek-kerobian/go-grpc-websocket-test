package service

import (
	"log"

	"github.com/hashicorp/go-memdb"
)

func InitCache() *memdb.MemDB {
    log.Println("[Application Log] Initializing the memdb data store with schema mappings")

    // Create the DB schema
    schema := &memdb.DBSchema{
        Tables: map[string]*memdb.TableSchema{
            "user": &memdb.TableSchema{
                Name: "user",
                Indexes: map[string]*memdb.IndexSchema{
                    "id": &memdb.IndexSchema{
                        Name:    "id",
                        Unique:  true,
                        Indexer: &memdb.UintFieldIndex{Field: "Id"},
                    },
                    "name": &memdb.IndexSchema{
                        Name:    "name",
                        Unique:  false,
                        Indexer: &memdb.StringFieldIndex{Field: "Name"},
                    },
                    "walletId": &memdb.IndexSchema{
                        Name:    "walletId",
                        Unique:  false,
                        Indexer: &memdb.UintFieldIndex{Field: "WalletId"},
                    },
                },
            },
            "wallet": &memdb.TableSchema{
                Name: "wallet",
                Indexes: map[string]*memdb.IndexSchema{
                    "id": &memdb.IndexSchema{
                        Name:    "id",
                        Unique:  true,
                        Indexer: &memdb.UintFieldIndex{Field: "Id"},
                    },
                    "balance": &memdb.IndexSchema{
                        Name:    "balance",
                        Unique:  false,
                        Indexer: &memdb.StringFieldIndex{Field: "Balance"},
                    },
                },
            },
        },
    }

    // Create a new data base
    db, err := memdb.NewMemDB(schema)
    if err != nil {
        panic(err)
    }

    return db
}

