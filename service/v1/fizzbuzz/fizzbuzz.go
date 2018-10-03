package fizzbuzz

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	apiV1 "github.com/tommy-42/go-grpc-fizzbuzzer/api/v1/fizzbuzz"
)

const (
	defaultFirstWord  = "fizz"
	defaultSecondWord = "buzz"
	defaultLimit      = 100
)

type Service struct {
	// might contains a DB or whatever the service needs
}

func NewFizzBuzzService() *Service {
	return &Service{}
}

// Accepts five parameters : two strings and three integers and returns a JSON
// It must return a list of strings with numbers from 1 to Limit, where:
// all multiples of firstMultiple are replaced by firstWord,
// all multiples of secondMultiple are replaced by secondWord,
// all multiples of firstMultiple and secondMultiple are replaced by firstWordsecondWord
func (s *Service) GetFizzBuzz(ctx context.Context, req *apiV1.GetFizzBuzzRequest) (*apiV1.GetFizzBuzzResponse, error) {

	if req.FirstMultiple == 0 || req.SecondMultiple == 0 {
		return nil, errors.Errorf("error: division by 0 is undefined, please choose another number")
	}

	if strings.TrimSpace(req.FirstWord) == "" {
		req.FirstWord = defaultFirstWord
	}
	if strings.TrimSpace(req.SecondWord) == "" {
		req.SecondWord = defaultSecondWord
	}
	if req.Limit < 1 {
		req.Limit = defaultLimit
	}

	var fizzbuzzed string
	for i := 1; i <= int(req.Limit); i++ {
		switch 0 {
		case i % int(req.FirstMultiple*req.SecondMultiple):
			fizzbuzzed = fmt.Sprintf("%s%s%s", fizzbuzzed, req.FirstWord, req.SecondWord)
		case i % int(req.FirstMultiple):
			fizzbuzzed = fmt.Sprintf("%s%s", fizzbuzzed, req.FirstWord)
		case i % int(req.SecondMultiple):
			fizzbuzzed = fmt.Sprintf("%s%s", fizzbuzzed, req.SecondWord)
		default:
			fizzbuzzed = fmt.Sprintf("%s%d", fizzbuzzed, i)
		}
		if i < int(req.Limit) {
			fizzbuzzed = fmt.Sprintf("%s,", fizzbuzzed)
		}
	}

	return &apiV1.GetFizzBuzzResponse{Fizzbuzzed: fizzbuzzed}, nil
}
