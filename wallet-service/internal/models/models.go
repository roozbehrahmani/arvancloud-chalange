package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Phone     *string `gorm:"unique;not null"`
	Wallet    Wallet
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime    `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type Wallet struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	UserId            uint
	Inventory         float64 `gorm:"not null;DEFAULT:0"`
	Balance           float64 `gorm:"not null;DEFAULT:0"`
	BlockedAmount     float64 `gorm:"not null;DEFAULT:0"`
	WalletTransaction []WalletTransaction
	CreatedAt         time.Time       `gorm:"autoCreateTime"`
	UpdatedAt         sql.NullTime    `gorm:"autoUpdateTime"`
	DeletedAt         *gorm.DeletedAt `gorm:"index"`
}

type WalletTransaction struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	WalletId  uint
	Amount    float64
	Type      bool
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime    `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type ChargeCode struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Code         string `gorm:"unique;not null"`
	Amount       float64
	MaxUsage     int
	CurrentUsage int
	Status       bool
	Users        []User `gorm:"many2many:charge_code_users;association_foreignkey:id;foreignkey:id"`
	ExpireAt     sql.NullTime
	CreatedAt    time.Time       `gorm:"autoCreateTime"`
	UpdatedAt    sql.NullTime    `gorm:"autoUpdateTime"`
	DeletedAt    *gorm.DeletedAt `gorm:"index"`
}

type ChargeCodeUser struct {
	UserId       uint
	ChargeCodeId uint
}
