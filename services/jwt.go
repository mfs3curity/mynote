package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mfs3curity/mynote/api/dto"
)

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

type tokenDto struct {
	UserId   int
	Username string
	Roles    []string
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.Expire = time.Now().Add(8760 * time.Hour).Unix()

	atc := jwt.MapClaims{}

	atc["UserId"] = token.UserId
	atc["username"] = token.Username
	atc["roles"] = token.Roles
	atc["exp"] = td.Expire

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.Token, err = at.SignedString([]byte("passw@rd!2#"))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unExpected error")
		}
		return []byte("passw@rd!2#"), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, errors.New("claims not found")
}
