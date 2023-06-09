package db

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Link        string
	Description string
	Language    string
	IsDeleted   bool
	Jobs        []Job `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
}

type Job struct {
	gorm.Model
	Title       string
	Slug        string
	Region      string
	Type        string
	PubDate     string
	Description string
	ApplyUrl    string
	Salary      string
	IsDeleted   bool
	SourceID    uint
	Source      Source
	CompanyID   uint     `gorm:"index"` // Add index to tables
	CategoryID  uint     `gorm:"index"`
	Company     Company  `gorm:"foreignKey:CompanyID"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Keyword     string
}

type Company struct {
	gorm.Model
	Name        string
	Headquarter string
	WebSite     string
	Logo        string
	IsDeleted   bool
	Jobs        []Job `gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE"`
}

type Source struct {
	gorm.Model
	Type string
	Url  string
}

type Applicant struct {
	gorm.Model
	Slug        string
	Application int
}
