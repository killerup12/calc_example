package model

import (
	"time"

	"gorm.io/gorm"
)

type Issue struct {
	ID                     uint           `json:"id" gorm:"primaryKey"`
	FullName               string         `json:"fullName" gorm:"not null"`
	ContactInfo            string         `json:"contactInfo" gorm:"not null"`
	PreferredContactMethod string         `json:"preferredContactMethod" gorm:"not null"`
	HasChinaExperience     bool           `json:"hasChinaExperience" gorm:"not null"`
	HasSupplierContacts    bool           `json:"hasSupplierContacts" gorm:"not null"`
	ProductDescription     string         `json:"productDescription" gorm:"not null"`
	ExistingProductLinks   string         `json:"existingProductLinks"`
	Volume                 *float64       `json:"volume,omitempty"`
	Weight                 *float64       `json:"weight,omitempty"`
	Density                *float64       `json:"density,omitempty"`
	PreviousInvoiceFile    string         `json:"previousInvoiceFile,omitempty"`
	ExpectedDeliveryDate   string         `json:"expectedDeliveryDate" gorm:"not null"`
	Status                 string         `json:"status" gorm:"default:'open'"`
	CreatedAt              time.Time      `json:"createdAt"`
	UpdatedAt              time.Time      `json:"updatedAt"`
	DeletedAt              gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateIssueRequest struct {
	FullName               string   `json:"fullName" binding:"required"`
	ContactInfo            string   `json:"contactInfo" binding:"required"`
	PreferredContactMethod string   `json:"preferredContactMethod" binding:"required"`
	HasChinaExperience     bool     `json:"hasChinaExperience"`
	HasSupplierContacts    bool     `json:"hasSupplierContacts"`
	ProductDescription     string   `json:"productDescription" binding:"required"`
	ExistingProductLinks   string   `json:"existingProductLinks"`
	Volume                 *float64 `json:"volume,omitempty"`
	Weight                 *float64 `json:"weight,omitempty"`
	Density                *float64 `json:"density,omitempty"`
	PreviousInvoiceFile    string   `json:"previousInvoiceFile,omitempty"`
	ExpectedDeliveryDate   string   `json:"expectedDeliveryDate" binding:"required"`
}

type UpdateIssueRequest struct {
	Status string `json:"status" binding:"required,oneof=open closed"`
}

type IssueResponse struct {
	ID                     uint      `json:"id"`
	FullName               string    `json:"fullName"`
	ContactInfo            string    `json:"contactInfo"`
	PreferredContactMethod string    `json:"preferredContactMethod"`
	HasChinaExperience     bool      `json:"hasChinaExperience"`
	HasSupplierContacts    bool      `json:"hasSupplierContacts"`
	ProductDescription     string    `json:"productDescription"`
	ExistingProductLinks   string    `json:"existingProductLinks"`
	Volume                 *float64  `json:"volume,omitempty"`
	Weight                 *float64  `json:"weight,omitempty"`
	Density                *float64  `json:"density,omitempty"`
	PreviousInvoiceFile    string    `json:"previousInvoiceFile,omitempty"`
	ExpectedDeliveryDate   string    `json:"expectedDeliveryDate"`
	Status                 string    `json:"status"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}
