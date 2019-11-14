package implementation

import (
	"errors"
	. "github.com/halapastefan/go-kit-micro/awesomeservice"
	"strings"
)

type Service struct {
}

func (s *Service) New() StringService {
	return &Service{}
}

func (s *Service) Uppercase(str string) (string, error) {
	if str == "" {
		return "", errors.New("error")
	}

	return strings.ToUpper(str), nil
}

func (s *Service) Count(str string) int {
	return len(str)
}
