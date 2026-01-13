package model

import (
	"time"
)

type ExternalProduct struct {
	ID           int32
	Source       string
	ExternalId   string
	Name         *string
	Description  *string
	Brand        *string
	ScrapedAt    time.Time
	StoreProduct *StoreProduct
}

type StoreProduct struct {
	ID                int32
	StoreId           int32
	SpecificProductId int32
	Price             *float64
	LastUpdated       time.Time
}
