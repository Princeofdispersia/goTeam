package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"github.com/Princeofdispersia/goTeam"
	"github.com/Princeofdispersia/goTeam/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const (
	salt       = "YKbpnNej978wckXCSgh5q"
	signingKey = "sC1#OJ?}wC*T{yn$"
	tokenTTL   = 168 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user goTeam.User) (int, string, error) {
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, "", err
	}
	token, err := s.GenerateToken(id)
	return id, token, err
}

func (s *AuthService) GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		id,
	})

	return token.SignedString([]byte(signingKey))

}

func (s *AuthService) CheckSig(id int, sig string) bool {
	hash := sha1.New()
	hash.Write([]byte(strconv.Itoa(id)))
	hash.Write([]byte(salt))
	h := hex.EncodeToString(hash.Sum(nil))
	return h == sig
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	lambda := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	}

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, lambda)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	// ПРОВЕРКА НА ИСТЕЧЕНИЕ ДАТЫ
	return claims.Id, nil

}
