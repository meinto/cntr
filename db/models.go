package db

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	UUID uuid.UUID
}

type Stats struct {
	gorm.Model
	ClientUUID uuid.UUID
	Year       int
	YearDay    int
	Hour       int
	Keys       int
	Clicks     int
}
