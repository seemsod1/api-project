package chain

import (
	"errors"
)

type RateService interface {
	GetRate(base, target string) (float64, error)
}

type Chain interface {
	RateService
	SetNext(chainInterface Chain)
}

var ErrNoRateProviders = errors.New("no available providers to get rate")

type BaseChain struct {
	rateService RateService
	next        Chain
}

func NewBaseChain(fetcher RateService) *BaseChain {
	return &BaseChain{rateService: fetcher}
}

func (b *BaseChain) SetNext(chainInterface Chain) {
	b.next = chainInterface
}

func (b *BaseChain) GetRate(base, target string) (float64, error) {
	rate, err := b.rateService.GetRate(base, target)
	if err != nil {
		next := b.next
		if next == nil {
			return -1, ErrNoRateProviders
		}

		return next.GetRate(base, target)
	}
	return rate, nil
}
