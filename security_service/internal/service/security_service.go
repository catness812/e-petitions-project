package security_service

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/gookit/slog"
	"github.com/streadway/amqp"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/catness812/e-petitions-project/security_service/internal/config"
	models "github.com/catness812/e-petitions-project/security_service/internal/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	CheckIfEmailExists(mail string) bool
	GetUserByEmail(email string) (models.User, error)
}

type IRedisRepository interface {
	ReplaceToken(currentToken, newToken string, expires time.Duration) error
	InsertUserToken(key string, value string, expires time.Duration) error
	InsertOTP(otp string, mail string, expires time.Duration) error
	GetOTP(mail string) (string, error)
	DeleteOTP(mail string) error
}

type SecurityService struct {
	userRepo  IUserRepository
	redisRepo IRedisRepository
}

func NewSecurityService(userRepo IUserRepository, redisRepo IRedisRepository) *SecurityService {
	return &SecurityService{
		userRepo:  userRepo,
		redisRepo: redisRepo,
	}
}

func (svc *SecurityService) Login(userLogin *models.UserCredentialsModel) (map[string]string, string, error) {
	user, err := svc.userRepo.GetUserByEmail(userLogin.Email)
	if err != nil {
		slog.Errorf("invalid credentials: %v", err)
		return nil, "", err
	}
	if err = svc.comparePasswordHash(user.Password, userLogin.Password); err != nil {
		return nil, "", errors.New("invalid credentials")
	}
	token, err := generateTokenPair(user.Email)
	if err != nil {
		slog.Errorf("Could not generate token pair %v", err)
		return nil, "", err
	}
	if err = svc.redisRepo.InsertUserToken(token["refresh_token"], user.Email, time.Hour*5); err != nil {
		return nil, "", err
	}
	return token, strconv.FormatUint(uint64(user.Id), 10), nil
}

func (svc *SecurityService) RefreshUserToken(token string, email string) (map[string]string, error) {
	tokenMap, err := generateTokenPair(email)
	if err != nil {
		return nil, err
	}
	if err := svc.redisRepo.ReplaceToken(token, tokenMap["refresh_token"], time.Hour*5); err != nil {
		return nil, err
	}
	return tokenMap, nil
}

func (svc *SecurityService) SendOTP(mail string, ch *amqp.Channel, cfg *config.Config) (string, error) {
	otp := svc.generateOTP()
	verifyLink := fmt.Sprintf(cfg.Gateway.Host + cfg.Gateway.Port + "/validate-otp?otp=" + otp + "&email=" + mail)
	message := fmt.Sprintf(mail + " " + verifyLink)
	if err := svc.redisRepo.InsertOTP(otp, mail, time.Minute*5); err != nil {
		return "", err
	}
	if err := svc.publishMessage(ch, message); err != nil {
		return "", err
	}
	return otp, nil
}

func (svc *SecurityService) VerifyOTP(mail string, userOTP string) error {
	otp, err := svc.redisRepo.GetOTP(mail)
	if err != nil {
		return errors.New("failed to get otp from redis")
	}
	if subtle.ConstantTimeCompare([]byte(otp), []byte(userOTP)) != 1 {
		return errors.New("OTP does not match")
	}
	if err := svc.redisRepo.DeleteOTP(mail); err != nil {
		return errors.New("failed to delete otp from redis: otp not found")
	}
	return nil
}

func (svc *SecurityService) generateOTP() string {
	numberSet := "0123456789"
	var password strings.Builder
	for i := 0; i < 5; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func (svc *SecurityService) publishMessage(ch *amqp.Channel, message string) error {
	err := ch.Publish("", "verify", false, false, amqp.Publishing{ContentType: "application/json", Body: []byte(message)})
	if err != nil {
		return err
	}
	return nil
}

func (svc *SecurityService) generatePasswordHash(pass string) (string, error) {
	const salt = 14
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	if err != nil {
		slog.Errorf("Failed to generate password hash: %v\n", err)
		return "", err
	}
	return string(hash), nil
}

func (svc *SecurityService) comparePasswordHash(hash, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		slog.Errorf("Failed to compare password and hash: %v\n", err)
		return err
	}
	return nil
}

func generateTokenPair(email string) (map[string]string, error) {
	keyConfig := config.LoadConfig()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userEmail"] = email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte(keyConfig.Token.TKey))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	if err != nil {
		return nil, err
	}

	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["userEmail"] = email
	rtClaims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	rt, err := refreshToken.SignedString([]byte(keyConfig.Token.RTKey))

	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
