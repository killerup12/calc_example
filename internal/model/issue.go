package model

import (
	"time"

	"gorm.io/gorm"
)

type Issue struct {
	ID                    uint           `json:"id" gorm:"primaryKey"`
	FullName              string         `json:"full_name" gorm:"not null"`
	ContactInfo           string         `json:"contact_info" gorm:"not null"`
	PreferredContactMethod string        `json:"preferred_contact_method" gorm:"not null"`
	HasChinaExperience    bool           `json:"has_china_experience" gorm:"not null"`
	HasSupplierContacts   bool           `json:"has_supplier_contacts" gorm:"not null"`
	ProductDescription    string         `json:"product_description" gorm:"not null"`
	ExistingProductLinks  string         `json:"existing_product_links"`
	Volume                *float64       `json:"volume,omitempty"`
	Weight                *float64       `json:"weight,omitempty"`
	Density               *float64       `json:"density,omitempty"`
	PreviousInvoiceFile   string         `json:"previous_invoice_file,omitempty"`
	ExpectedDeliveryDate  string         `json:"expected_delivery_date" gorm:"not null"`
	Status                string         `json:"status" gorm:"default:'open'"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateIssueRequest struct {
	FullName              string   `json:"full_name" binding:"required"`
	ContactInfo           string   `json:"contact_info" binding:"required"`
	PreferredContactMethod string  `json:"preferred_contact_method" binding:"required"`
	HasChinaExperience    bool     `json:"has_china_experience"`
	HasSupplierContacts   bool     `json:"has_supplier_contacts"`
	ProductDescription    string   `json:"product_description" binding:"required"`
	ExistingProductLinks  string   `json:"existing_product_links"`
	Volume                *float64 `json:"volume,omitempty"`
	Weight                *float64 `json:"weight,omitempty"`
	Density               *float64 `json:"density,omitempty"`
	PreviousInvoiceFile   string   `json:"previous_invoice_file,omitempty"`
	ExpectedDeliveryDate  string   `json:"expected_delivery_date" binding:"required"`
}

type UpdateIssueRequest struct {
	Status string `json:"status" binding:"required,oneof=open closed"`
}

type IssueResponse struct {
	ID                    uint      `json:"id"`
	FullName              string    `json:"full_name"`
	ContactInfo           string    `json:"contact_info"`
	PreferredContactMethod string   `json:"preferred_contact_method"`
	HasChinaExperience    bool      `json:"has_china_experience"`
	HasSupplierContacts   bool      `json:"has_supplier_contacts"`
	ProductDescription    string    `json:"product_description"`
	ExistingProductLinks  string    `json:"existing_product_links"`
	Volume                *float64  `json:"volume,omitempty"`
	Weight                *float64  `json:"weight,omitempty"`
	Density               *float64  `json:"density,omitempty"`
	PreviousInvoiceFile   string    `json:"previous_invoice_file,omitempty"`
	ExpectedDeliveryDate  string    `json:"expected_delivery_date"`
	Status                string    `json:"status"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
} 