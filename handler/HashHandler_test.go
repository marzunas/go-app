package handler

import (
	"github.com/marzunas/go-app/conf"
	"github.com/marzunas/go-app/store"
	"math/rand"
	"testing"
	"time"
)

//Test class for HashHandler
//TODO: Additional test cases
func Test_computeHash(t *testing.T) {
	var cfg = conf.GetConfig()
	var password = "super-secret-password"
	var id = rand.Int()
	computeHash(id, password)
	var hash = store.HashStore[id]
	if cfg.HashWaitPeriod != 0 &&  hash!= "" {
		t.Error("hash calculation did not wait")
	}
	time.Sleep(time.Duration(cfg.HashWaitPeriod + 2)* time.Second)
	hash = store.HashStore[id]
	if hash == "" {
		t.Error("hash should not be empty")
	}
}

func Test_getNextId(t *testing.T) {
	var idOne = getNextId()
	var idTwo = getNextId()
	if idTwo <= idOne {
		t.Error("generated id may not be unique or incrementing")
	}
}
