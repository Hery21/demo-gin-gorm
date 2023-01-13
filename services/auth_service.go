package services

import (
	"fmt"
	" hery-ciaputra/demo-gin/config"
	" hery-ciaputra/demo-gin/dto"
	" hery-ciaputra/demo-gin/httperror"
	" hery-ciaputra/demo-gin/models"
	repositories " hery-ciaputra/demo-gin/repository"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type AuthService interface {
	SignIn(*dto.SignInReq) (*dto.TokenResponse, error)
}

type authService struct {
	userRepository repositories.UserRepository
	appConfig      config.AppConfig
}

type AuthSConfig struct {
	UserRepository repositories.UserRepository
	AppConfig      config.AppConfig
}

func NewAuthService(c *AuthSConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *models.User `json:"user"`
}

// ref: https://pkg.go.dev/github.com/golang-jwt/jwt#example-NewWithClaims-CustomClaimsType
func (a *authService) generateJWTToken(user *models.User) (*dto.TokenResponse, error) {
	// todo: create custom claims that will be added to jwt payload

	var idExp = a.appConfig.JWTExpireInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	timeExpire := jwt.NumericDate{Time: time.Unix(tokenExp, 0)}
	timeNow := jwt.NumericDate{Time: time.Now()}

	claims := &idTokenClaims{
		// todo: create custom claims
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.appConfig.AppName,
			IssuedAt:  &timeNow,
			ExpiresAt: &timeExpire,
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.appConfig.JWTSecret)

	// todo: resp error when error occurred
	if err != nil {
		return new(dto.TokenResponse), httperror.BadRequestError("BAD_REQUEST", "")
	}
	return &dto.TokenResponse{IDToken: tokenString}, nil

	// todo: return token resp
	// { idToken: "jwt-token" }
}

func (a *authService) SignIn(req *dto.SignInReq) (*dto.TokenResponse, error) {
	user, err := a.userRepository.MatchingCredential(req.Email, req.Password)
	if err != nil || user == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED",
			Message:    "Unauthorized",
		}
	}
	token, err := a.generateJWTToken(user)
	fmt.Println("lewatttttt")
	return token, err
}
