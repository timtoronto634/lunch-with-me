package plan

import (
	"echo-me/domain/entity/user"
	"time"

	"github.com/google/uuid"
)

type CandidateID uuid.UUID
type hostUserID user.UserID
type Candidate struct {
	ID           CandidateID
	hostID       hostUserID
	internalTime time.Time
	status       CandidateStatus
}

type CandidateStatus string

const (
	Open      CandidateStatus = "open"
	Scheduled CandidateStatus = "scheduled"
	Completed CandidateStatus = "completed"
	Deleted   CandidateStatus = "deleted"
)

func NewCandidate(hostID hostUserID, date time.Time) *Candidate {
	return &Candidate{
		ID:           generateCandidateID(),
		hostID:       hostID,
		internalTime: date,
		status:       Open,
	}
}

func generateCandidateID() CandidateID {
	// ランダムなIDや連番、UUIDなどを使用してIDを生成
	id, _ := uuid.NewUUID()
	return CandidateID(id)
}
