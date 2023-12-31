package plan

import (
	"context"
	"echo-me/domain/entity/user"
	"errors"
	"time"
)

type Plan struct {
	Candidate    *Candidate
	Reservations []*reservation
}

type Repository interface {
	GetByResearvationID(ctx context.Context, id ReservationID) (*Plan, error)
	Save(ctx context.Context, p *Plan) error
}

func NewPlan(hostID user.UserID, date time.Time) *Plan {
	return &Plan{
		Candidate: NewCandidate(hostID, date),
	}
}

func (p *Plan) Settle(rID ReservationID) error {
	if p.Candidate.status != Scheduled {
		return errors.New("candidate is not wainting for settlement")
	}

	p.Candidate.status = Completed

	return nil
}

func (p *Plan) Request(guestUserID guestUserID) error {
	if p.Candidate.status != Open {
		return errors.New("candidate is not open for reservation")
	}
	if p.isSchedulingInProgress() {
		return errors.New("a reservation already exists for this candidate")
	}

	p.Reservations = append(p.Reservations, NewReservation(p.Candidate.ID, guestUserID))
	p.Candidate.status = Scheduled

	return nil
}

func (p *Plan) isSchedulingInProgress() bool {
	for _, r := range p.Reservations {
		if r.status == Requesting {
			return true
		}
	}
	return false
}
