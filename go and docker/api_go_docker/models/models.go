package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	id int	`json:"id" gorm:"text;not null;default:null`
	eid string	`json:"eid" gorm:"text;not null;default:nul`
	ename string	`json:"ename" gorm:"text;not null;default:nul`
	eemail string	`json:"eemail" gorm:"text;not null;default:nul`
	econtact string	`json:"econtact" gorm:"text;not null;default:nul`
}