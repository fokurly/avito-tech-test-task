package models

import "sync"

type Client struct {
	Id      int64
	Balance float64
	mu      sync.RWMutex
}
