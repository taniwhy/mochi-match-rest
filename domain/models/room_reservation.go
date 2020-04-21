package models

import (
	"time"
)

// RoomReservation : room_reservationテーブルモデル
type RoomReservation struct {
	ID           int
	Room         int
	RelasingTime time.Time
	CreatedAt    time.Time
}
