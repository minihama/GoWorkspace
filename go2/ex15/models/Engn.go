package models

import (
	"github.com/dgraph-io/ristretto"
	"xorm.io/xorm"
)

type Engn struct {
	db    *xorm.Engine
	cache *ristretto.Cache
}

func NewEngn() (*Engn, error) {
	engn, err := xorm.NewEngine("mysql", "go1:goldMa$k46@tcp(192.168.0.16:3306)/go1")
	if err != nil {
		return nil, err
	}
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	e := Engn{
		db:    engn,
		cache: cache,
	}
	return &e, nil
}

func (engn *Engn) CloseEngn() error {
	engn.cache.Close()
	return engn.db.Close()
}
