package plan

import (
	"errors"
	"time"
)

type Plan struct {
	Candidate    *Candidate
	Reservations []*reservation
}

func NewPlan(hostID hostUserID, date time.Time) *Plan {
	return &Plan{
		Candidate: NewCandidate(hostID, date),
	}
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
