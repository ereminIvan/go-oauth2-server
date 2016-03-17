package go_oauth2_server

import (
	"errors"

	"github.com/ereminIvan/go-oauth2-server/api/grants"
)

var (
	ErrorInvalidGrantType = errors.New("Invalid grant type")
)

type IService interface {
	SetGrantType(t grants.IGrant)
}

func NewService(t grants.IGrant) *Service{
	return &Service{}
}

type Service struct {
	grantType []grants.IGrant
}

func (s *Service) SetGrantType(t grants.IGrant) {
	if t == nil {
		return ErrorInvalidGrantType
	}
	if len(s.grantType) {
		s.grantType = []grants.IGrant{}
	}
	s.grantType = append(s.grantType, t)
}