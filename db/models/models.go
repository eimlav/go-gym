package models

import (
	"time"

	"gorm.io/gorm"
)

// Class represents a gym class.
type Class struct {
	gorm.Model
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  uint
}

// ClassEvent represents an individual event associated with a Class.
type ClassEvent struct {
	gorm.Model
	ClassID uint
	Class   Class
	Date    time.Time
}

// Booking represents a booking of a ClassEvent.
type Booking struct {
	gorm.Model
	ClassEventID uint
	ClassEvent   ClassEvent
	MemberName   string
}
