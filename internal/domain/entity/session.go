package entity

import "time"

type Session struct {
	ID                string
	UserID            string
	Device            string
	IP                string
	CurrentJTI        string
	CreatedAt         time.Time
	ExpiresAt         time.Time
	UpdatedAt         time.Time
	JTITTLSeconds     int64
	SessionTTLSeconds int64
}
