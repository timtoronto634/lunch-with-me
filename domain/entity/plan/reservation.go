package plan

import (
	"echo-me/domain/entity/user"

	"github.com/google/uuid"
)

type ReservationID uuid.UUID

type guestUserID user.UserID
type reservationStatus string

const (
	Initialized reservationStatus = "initialized"
	Requesting  reservationStatus = "requesting"
	Accepted    reservationStatus = "accepted"
	Declined    reservationStatus = "declined"
)

type reservation struct {
	ID      ReservationID
	candID  CandidateID
	guestID guestUserID
	status  reservationStatus
}

func NewReservation(candidateID CandidateID, guestID guestUserID) *reservation {
	return &reservation{
		ID:      generateReservationID(), // 予約ID生成の実装に依存
		candID:  candidateID,
		guestID: guestID,
		status:  Requesting,
	}
}

// 例: 予約IDを生成するためのダミー関数
// この関数の実装は、実際のシステムの要件や使用するストレージの特性に応じて変わる可能性があります。
func generateReservationID() ReservationID {
	// ランダムなIDや連番、UUIDなどを使用してIDを生成
	id, _ := uuid.NewUUID()
	return ReservationID(id)
}
