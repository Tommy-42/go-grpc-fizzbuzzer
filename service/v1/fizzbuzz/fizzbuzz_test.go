package fizzbuzz

import (
	"context"
	"testing"

	api "github.com/tommy-42/go-grpc-fizzbuzzer/api/v1/fizzbuzz"

	"github.com/stretchr/testify/assert"
)

var (
	svc *Service
)

func init() {
	svc = NewFizzBuzzService()
}

func TearDownTest(t *testing.T) {}

func TestCreateCompany(t *testing.T) {
	assert := assert.New(t)
	defer TearDownTest(t)

	var tests = []struct {
		name        string
		req         *api.GetFizzBuzzRequest
		expectedErr error
	}{
		// Should succeed
		{
			name: "Fizzbuzz req valid full fill",
			req: &api.GetFizzBuzzRequest{
				FirstWord:      "fazz",
				SecondWord:     "bozz",
				FirstMultiple:  3,
				SecondMultiple: 5,
				Limit:          100,
			},
		},
		{
			name: "Fizzbuzz req valid full fill with limit 0",
			req: &api.GetFizzBuzzRequest{
				FirstWord:      "fazz",
				SecondWord:     "bozz",
				FirstMultiple:  3,
				SecondMultiple: 5,
				Limit:          0,
			},
		},
		{
			name: "Fizzbuzz req valid with missing first/second_word",
			req: &api.GetFizzBuzzRequest{
				FirstMultiple:  3,
				SecondMultiple: 5,
				Limit:          100,
			},
		},
		{
			name: "Fizzbuzz req valid with missing first/second_word and limit 0",
			req: &api.GetFizzBuzzRequest{
				FirstMultiple:  3,
				SecondMultiple: 5,
				Limit:          0,
			},
		},
		// Should fail
		{name: "fizzbuzz req nil", req: nil, expectedErr: ErrRequestNil},
		{name: "fizzbuzz req empty", req: &api.GetFizzBuzzRequest{}, expectedErr: ErrDivisionByZero},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := svc.GetFizzBuzz(context.Background(), test.req)
			if test.expectedErr == nil {
				assert.Nil(err)
				if assert.NotNil(res) && assert.NotEmpty(res.Fizzbuzzed) {
					if test.req.FirstWord == "" {
						assert.Contains(res.Fizzbuzzed, defaultFirstWord)
					} else {
						assert.Contains(res.Fizzbuzzed, test.req.FirstWord)
					}
					if test.req.SecondWord == "" {
						assert.Contains(res.Fizzbuzzed, defaultSecondWord)
					} else {
						assert.Contains(res.Fizzbuzzed, test.req.SecondWord)
					}
				}
			}
			if test.expectedErr != nil {
				if assert.Error(err) {
					assert.Contains(err.Error(), test.expectedErr.Error())
				}
				assert.Nil(res)
			}
		})
	}
}
