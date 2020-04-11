package models

import (
	"time"
)

// RoomReservation : room_reservationテーブルモデル
type RoomReservation struct {
	ID           int
	Room         int
	ReeasingTime time.Time
	CreatedAt    time.Time
}
