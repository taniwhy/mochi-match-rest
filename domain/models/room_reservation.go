package models

import (
	"time"
)

// RoomReservation : room_reservationテーブルモデル
type RoomReservation struct {
	ID           int64
	Room         int64
	RelasingTime time.Time
	CreatedAt    time.Time
}
