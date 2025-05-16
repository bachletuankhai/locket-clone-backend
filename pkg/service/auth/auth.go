package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	claims := JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  jwt.ClaimStrings{},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	key, err := os.ReadFile("./private.key")
	if err != nil {
		return "", err
	}
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GenerateRefreshToken(username string) (string, error) {
	claims := JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  jwt.ClaimStrings{},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	key, err := os.ReadFile("./private.key")
	if err != nil {
		return "", err
	}
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

type Token struct {
	Auth    string `json:"auth"`
	Refresh string `json:"refresh"`
}

type AuthService interface {
	Login(username, password string) (Token, error)
	// RefreshToken(token string) (Token, error)  // TODO: Implement refresh token
	Logout(token string) error
	ParseToken(token string) (JwtClaims, error)
}

type TokenRepo interface {
	SaveToken(token string, exp time.Time) error
	CheckTokenExists(token string) (bool, error)
}

type UserRepo interface {
	GetUserPasswordHashByUsername(username string) (string, error)
}

type authService struct {
	TokenBlacklist TokenRepo
	UserRepo       UserRepo
}

func (s *authService) ParseToken(token string) (JwtClaims, error) {
	if token == "" {
		return JwtClaims{}, errors.New("token is empty")
	}
	if isBlacklisted, err := s.TokenBlacklist.CheckTokenExists(token); err != nil {
		return JwtClaims{}, errors.New("cannot validate token")
	} else if isBlacklisted {
		return JwtClaims{}, errors.New("token is blacklisted")
	}

	claims := JwtClaims{}

	res, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return os.ReadFile("./public.key")
	}, jwt.WithValidMethods([]string{"RS256", "RS384", "RS512"}))
	if err != nil {
		return claims, err
	}

	if res.Valid {
		return claims, nil
	}
	return claims, errors.New("Unauthorized")
}

func (s *authService) Login(username string, password string) (Token, error) {
	hash, err := s.UserRepo.GetUserPasswordHashByUsername(username)
	if err != nil {
		return Token{}, err
	}
	if !CompareHashPassword(password, hash) {
		return Token{}, errors.New("invalid password")
	}

	token, err := GenerateToken(username)
	if err != nil {
		return Token{}, err
	}

	refreshToken, err := GenerateRefreshToken(username)
	if err != nil {
		return Token{}, err
	}

	return Token{
		Auth:    token,
		Refresh: refreshToken,
	}, nil

}

// RefreshToken(token string) (Token, error)  // TODO: Implement refresh token
func (s *authService) Logout(token string) error {
	claims, err := s.ParseToken(token)
	if err != nil {
		return err
	}

	if err := s.TokenBlacklist.SaveToken(token, claims.ExpiresAt.Time); err != nil {
		return err
	}
	return nil
}

func NewAuthService(tokenBlacklist TokenRepo, userRepo UserRepo) AuthService {
	return &authService{
		TokenBlacklist: tokenBlacklist,
		UserRepo:       userRepo,
	}
}
