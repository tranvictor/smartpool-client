package protocol

import (
	"github.com/SmartPool/smartpool-client"
)

// InMemClaimRepo implements ClaimRepo interface. It stores claims in one start
// up. However this claim repo doesn't persist verified claims and even the
// active claim. So if the client is shutdown, all past claims and current
// shares information will be lost.
// This shouldn't be used in production.
type InMemClaimRepo struct {
	claims       map[int]*Claim
	cClaimNumber uint64
}

func NewInMemClaimRepo() *InMemClaimRepo {
	return &InMemClaimRepo{
		map[int]*Claim{0: NewClaim()},
		0,
	}
}

func (cr *InMemClaimRepo) AddShare(s smartpool.Share) {
	cr.claims[int(cr.cClaimNumber)].AddShare(s)
}

func (cr *InMemClaimRepo) GetClaim(number uint64) smartpool.Claim {
	return cr.claims[int(number)]
}

// TODO: This needs lock to prevent concurrent writes
func (cr *InMemClaimRepo) GetCurrentClaim(threshold int) smartpool.Claim {
	c := cr.GetClaim(cr.cClaimNumber)
	if c.NumShares().Int64() < int64(threshold) {
		return nil
	}
	cr.cClaimNumber++
	cr.claims[int(cr.cClaimNumber)] = NewClaim()
	return c
}
