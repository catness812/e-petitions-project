package security_service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	models "github.com/catness812/e-petitions-project/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepostory interface {
	Register(user *models.UserModel) error
	CheckIfEmailExists(mail string) bool
	GetUserByEmail(email string) (models.UserModel, error)
}

type IRedisRepository interface {
	ReplaceToken(currentToken, newToken string, expires time.Duration) error
	InsertUserToken(key string, value string, expires time.Duration) error
}

type SecurityService struct {
	userRepo  IUserRepostory
	redisRepo IRedisRepository
}

func NewSecurityService(userRepo IUserRepostory, redisRepo IRedisRepository) *SecurityService {
	return &SecurityService{
		userRepo:  userRepo,
		redisRepo: redisRepo,
	}
}

func (svc *SecurityService) Register(user *models.UserModel) error {
	if svc.userRepo.CheckIfEmailExists(user.Email) {
		return errors.New("user already exists")
	}

	hash, err := svc.generatePasswordHash(user.Password)
	if err != nil {
		return errors.New("err can't register user")
	}
	user.Password = hash

	return svc.userRepo.Register(user)
}

func (svc *SecurityService) Login(userLogin *models.UserCredentialsModel) (map[string]string, error) {

	user, err := svc.userRepo.GetUserByEmail(userLogin.Email)

	if err != nil {
		log.Printf("invalid credentials: %v", err)
		return nil, err
	}

	err = svc.comparePasswordHash(user.Password, userLogin.Password)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := generateTokenPair(user.ID)

	if err != nil {
		return nil, errors.New("can't generate token")
	}

	//token_lifespan, err := strconv.Atoi(os.Getenv("RTOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return nil, nil
	}

	err = svc.redisRepo.InsertUserToken(token["refresh_token"], fmt.Sprint(user.ID), time.Hour*5)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (svc *SecurityService) generatePasswordHash(pass string) (string, error) {
	const salt = 14
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	if err != nil {
		log.Printf("ERR: %v\n", err)
		return "", err
	}

	return string(hash), nil
}

func (svc *SecurityService) comparePasswordHash(hash, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		log.Printf("ERR: %v\n", err)
		return err
	}

	return nil
}

func (svc *SecurityService) RefreshUserToken(token string, id uint) (map[string]string, error) {
	tokenMap, err := generateTokenPair(id)
	if err != nil {
		return nil, err
	}
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_MINUTE_LIFESPAN"))
	if err != nil {
		return nil, err
	}
	if err := svc.redisRepo.ReplaceToken(token, tokenMap["rt"], time.Duration(token_lifespan)); err != nil {
		return nil, err
	}
	return tokenMap, nil
}

func generateTokenPair(userId uint) (map[string]string, error) {
	err := godotenv.Load("internal/security/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_MINUTE_LIFESPAN"))
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()

	t, err := token.SignedString([]byte(os.Getenv("T_KEY")))

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rtoken_lifespan, err := strconv.Atoi(os.Getenv("RTOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return nil, err
	}

	rtClaims := refreshToken.Claims.(jwt.MapClaims)

	rtClaims["userID"] = userId
	rtClaims["exp"] = time.Now().Add(time.Hour * time.Duration(rtoken_lifespan)).Unix()

	rt, err := refreshToken.SignedString([]byte(os.Getenv("RT_KEY")))

	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
