package schema

import (
	"time"

	"gorm.io/gorm"
)

// --- User and related entities ---

type Municipality struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(255);not null;unique"`
	Users []User `gorm:"foreignKey:MunicipalityID"`
}

type PaymentState struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"type:varchar(100);not null;unique"`
	Payments []Payment `gorm:"foreignKey:StateID"`
}

type TypeCylinder struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Disponible  bool    `gorm:"default:true"`

	OrderDetails    []OrderDetail    `gorm:"foreignKey:TypeCylinderID"`
	DeliveryDetails []DeliveryDetail `gorm:"foreignKey:TypeCylinderID"`
}

type OrderState struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"type:varchar(100);not null;unique"`
	Orders []Order `gorm:"foreignKey:StateOrderID"`
}

type ReportType struct {
	ID      uint     `gorm:"primaryKey"`
	Name    string   `gorm:"type:varchar(100);not null;unique"`
	Reports []Report `gorm:"foreignKey:TypeID"`
}

type ReportState struct {
	ID      uint     `gorm:"primaryKey"`
	Name    string   `gorm:"type:varchar(100);not null;unique"`
	Reports []Report `gorm:"foreignKey:ReportStateID"`
}

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(255);not null"`
	Email          string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password       string `gorm:"type:varchar(255);not null"`
	Active         bool   `gorm:"default:true"`
	MunicipalityID uint   `gorm:"not null"`

	Municipality Municipality `gorm:"foreignKey:MunicipalityID"`
	Commerce     *Commerce    `gorm:"foreignKey:UserID"`
	Disabled     *Disabled    `gorm:"foreignKey:UserID"`
	Council      *Council     `gorm:"foreignKey:UserID"`
	Orders       []Order      `gorm:"foreignKey:UserID"`
}

type Disabled struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"unique;not null"`
	Document   string `gorm:"type:varchar(255)"`
	Disability string `gorm:"type:text"`
	User       User   `gorm:"foreignKey:UserID"`
}

type Council struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `gorm:"unique;not null"`
	LeaderName     string `gorm:"type:varchar(255);not null"`
	LeaderDocument string `gorm:"type:varchar(20);not null;unique"`
	User           User   `gorm:"foreignKey:UserID"`
}

type Commerce struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"unique;not null"`
	Rif          string `gorm:"type:varchar(20);not null;unique"`
	BossName     string `gorm:"type:varchar(255);not null"`
	BossDocument string `gorm:"type:varchar(20);not null;unique"`
	User         User   `gorm:"foreignKey:UserID"`
}

type Order struct {
	gorm.Model
	UserID       uint    `gorm:"not null"`
	TotalPrice   float64 `gorm:"not null"`
	StateOrderID uint    `gorm:"not null"`

	User         User          `gorm:"foreignKey:UserID"`
	OrderState   OrderState    `gorm:"foreignKey:StateOrderID"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
	Delivery     *Delivery     `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID             uint    `gorm:"primaryKey"`
	OrderID        uint    `gorm:"not null"`
	TypeCylinderID uint    `gorm:"not null"`
	Quantity       int     `gorm:"not null"`
	Price          float64 `gorm:"not null"`

	Order        Order        `gorm:"foreignKey:OrderID"`
	TypeCylinder TypeCylinder `gorm:"foreignKey:TypeCylinderID"`
}

type Delivery struct {
	gorm.Model
	OrderID    uint    `gorm:"unique;not null"`
	PaymentID  uint    `gorm:"unique;not null"`
	TotalPrice float64 `gorm:"not null"`

	Order           Order            `gorm:"foreignKey:OrderID"`
	Payment         *Payment         `gorm:"foreignKey:PaymentID"`
	DeliveryDetails []DeliveryDetail `gorm:"foreignKey:DeliveryID"`
	Reports         []Report         `gorm:"foreignKey:DeliveryID"`
}

type DeliveryDetail struct {
	gorm.Model
	DeliveryID     uint         `gorm:"not null"`
	TypeCylinderID uint         `gorm:"not null"`
	Quantity       int          `gorm:"not null"`
	Delivery       Delivery     `gorm:"foreignKey:DeliveryID"`
	TypeCylinder   TypeCylinder `gorm:"foreignKey:TypeCylinderID"`
}

type Report struct {
	gorm.Model
	DeliveryID    uint      `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	Description   string    `gorm:"type:text;not null"`
	Date          time.Time `gorm:"not null"`
	TypeID        uint      `gorm:"not null"`
	ReportStateID uint      `gorm:"not null"`

	User        User        `gorm:"foreignKey:UserID"`
	Delivery    Delivery    `gorm:"foreignKey:DeliveryID"`
	ReportType  ReportType  `gorm:"foreignKey:TypeID"`
	ReportState ReportState `gorm:"foreignKey:ReportStateID"`
}

type Payment struct {
	gorm.Model
	UserID   uint    `gorm:"not null"`
	OrderID  uint    `gorm:"not null"`
	Quantity float64 `gorm:"not null"`
	StateID  uint    `gorm:"not null"`

	User         User         `gorm:"foreignKey:UserID"`
	Order        Order        `gorm:"foreignKey:OrderID"`
	PaymentState PaymentState `gorm:"foreignKey:StateID"`
	Delivery     *Delivery    `gorm:"foreignKey:PaymentID"`
}
