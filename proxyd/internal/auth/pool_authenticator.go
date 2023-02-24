package auth

import (
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/interplanetary-network/proxyd/internal/exchange"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrInvalidSignature = errors.New("invalid signature")
	ErrExpired          = errors.New("expired")
	ErrNotStarted       = errors.New("not started")
)

type Cache interface {
	Set(key any, value any)
	Get(key any) (any, bool)
	Del(key any)
}

type poolAuthenticator struct {
	poolID             string
	contractAddress    string
	ethereumGatewayURL string
	cache              Cache
}

// NewPoolAuthenticator creates a new authenticator for user which can access to a pool.
func NewPoolAuthenticator(poolID string, contractAddress string, ethereumGatewayURL string, cache Cache) Authenticator {
	return &poolAuthenticator{
		poolID:             poolID,
		cache:              cache,
		contractAddress:    contractAddress,
		ethereumGatewayURL: ethereumGatewayURL,
	}
}

// Authenticate checks if the username and password are valid.
// The username is the order id
// The password is signature string
func (a *poolAuthenticator) Authenticate(username, password string) error {
	sig, err := hexutil.Decode(password)
	if err != nil {
		return err
	}

	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	pubKey, err := crypto.SigToPub(accounts.TextHash([]byte(username)), sig)
	if err != nil {
		return err
	}

	orderID := new(big.Int)
	orderID, ok := orderID.SetString(username, 10)
	if !ok {
		return errors.New("Invalid order id")
	}

	// check from cache
	var order exchange.ProxyExchangeOrder

	o, ok := a.cache.Get(orderID.String())
	if ok {
		order, ok = o.(exchange.ProxyExchangeOrder)
		if !ok {
			return errors.New("Invalid order data in cache")
		}
	} else {
		order, err = a.orderOf(orderID)
		if err != nil {
			return err
		}
	}

	// compare customer address
	if order.Customer.Hex() != crypto.PubkeyToAddress(*pubKey).Hex() {
		return ErrInvalidSignature
	}

	// compare pool ID
	poolID := new(big.Int)
	poolID, ok = poolID.SetString(a.poolID, 10)
	if !ok {
		return errors.New("Invalid pool id")
	}
	if order.PoolID.Cmp(poolID) != 0 {
		return ErrInvalidSignature
	}

	now := time.Now().UTC().Unix()
	// compare order start time
	if uint64(now) < order.StartAt {
		return ErrNotStarted
	}

	// compare order expires
	// order.Duration is in minutes
	orderExp := order.StartAt + uint64(order.Duration*60)
	if uint64(now) >= orderExp {
		a.cache.Del(orderID.String())
		return ErrExpired
	}

	// valid order
	a.cache.Set(orderID.String(), order)

	return nil
}

func (a *poolAuthenticator) orderOf(orderID *big.Int) (exchange.ProxyExchangeOrder, error) {
	client, err := ethclient.Dial(a.ethereumGatewayURL)
	if err != nil {
		return *new(exchange.ProxyExchangeOrder), err
	}

	ex, err := exchange.NewExchange(common.HexToAddress(a.contractAddress), client)
	if err != nil {
		return *new(exchange.ProxyExchangeOrder), err
	}

	return ex.OrderOf(nil, orderID)
}

type poolOrderCache struct {
	data sync.Map
}

func NewPoolOrderCache() Cache {
	return &poolOrderCache{data: sync.Map{}}
}

func (c *poolOrderCache) Get(key any) (any, bool) {
	return c.data.Load(key)
}

func (c *poolOrderCache) Set(key, value any) {
	c.data.Store(key, value)
}

func (c *poolOrderCache) Del(key any) {
	c.data.Delete(key)
}
