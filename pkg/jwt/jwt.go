package jwt

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type JWTClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Config struct {
	AccessSecret  string
	RefreshSecret string
	AccessExpire  time.Duration
	RefreshExpire time.Duration
	Issuer        string
}

type JWT struct {
	config Config
	redis  *redis.Client
	// 添加互斥锁，用于并发控制
	renewLock sync.Mutex
}

func New(config Config, redis *redis.Client) *JWT {
	return &JWT{
		config: config,
		redis:  redis,
	}
}

func (j *JWT) GenerateToken(userID uint64, username string) (accessToken, refreshToken string, err error) {
	// 生成 Access Token
	accessClaims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.AccessExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(j.config.AccessSecret))
	if err != nil {
		return "", "", err
	}

	// 生成 Refresh Token
	refreshClaims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.RefreshExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(j.config.RefreshSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (j *JWT) ParseToken(ctx context.Context, tokenString string, isRefreshToken bool) (*JWTClaims, error) {
	// 检查是否在黑名单中
	if !isRefreshToken {
		inBlacklist, err := j.IsInBlacklist(ctx, tokenString)
		if err != nil {
			return nil, err
		}
		if inBlacklist {
			return nil, errors.New("token is in blacklist")
		}
	}

	secret := j.config.AccessSecret
	if isRefreshToken {
		secret = j.config.RefreshSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (j *JWT) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	claims, err := j.ParseToken(ctx, refreshToken, true)
	if err != nil {
		return "", "", err
	}

	return j.GenerateToken(claims.UserID, claims.Username)
}

// 生成黑名单的 key
func (j *JWT) getBlacklistKey(tokenStr string) string {
	return "token:blacklist:" + tokenStr
}

// 生成续期记录的 key
func (j *JWT) getRenewalKey(userID uint64) string {
	return "token:renewal:" + string(userID)
}

// 将 token 加入黑名单
func (j *JWT) AddToBlacklist(ctx context.Context, tokenStr string, claims *JWTClaims) error {
	expiration := time.Until(claims.ExpiresAt.Time)
	if expiration < 0 {
		log.Printf("token already expired for user %d", claims.UserID)
		return nil
	}

	if err := j.redis.Set(ctx, j.getBlacklistKey(tokenStr), "1", expiration).Err(); err != nil {
		log.Printf("failed to add token to blacklist for user %d: %v", claims.UserID, err)
		return err
	}

	log.Printf("token added to blacklist for user %d", claims.UserID)
	return nil
}

// 检查 token 是否在黑名单中
func (j *JWT) IsInBlacklist(ctx context.Context, tokenStr string) (bool, error) {
	i, err := j.redis.Exists(ctx, j.getBlacklistKey(tokenStr)).Result()
	return i != 0, err
}

// CheckAndRenewToken 检查并自动续期 token
func (j *JWT) CheckAndRenewToken(ctx context.Context, tokenStr string, claims *JWTClaims) (newAccessToken string, needRenew bool, err error) {
	j.renewLock.Lock()
	defer j.renewLock.Unlock()

	remainingTime := time.Until(claims.ExpiresAt.Time)

	if remainingTime < j.config.AccessExpire/4 {
		renewalKey := j.getRenewalKey(claims.UserID)
		exists, err := j.redis.Exists(ctx, renewalKey).Result()
		if err != nil {
			return "", false, fmt.Errorf("check renewal status failed: %w", err)
		}

		if exists != 0 {
			newAccessToken, err = j.generateAccessToken(claims.UserID, claims.Username)
			if err != nil {
				return "", false, fmt.Errorf("generate new token failed: %w", err)
			}

			// 使用 lua 脚本确保原子性
			script := `
				if redis.call("exists", KEYS[1]) == 0 then
					redis.call("setex", KEYS[1], ARGV[1], "1")
					return 1
				end
				return 0
			`

			success, err := j.redis.Eval(ctx, script, []string{renewalKey}, int(remainingTime.Seconds())).Int64()
			if err != nil {
				return "", false, fmt.Errorf("set renewal status failed: %w", err)
			}

			if success == 1 {
				return newAccessToken, true, nil
			}
		}
	}

	return "", false, nil
}

// 生成访问令牌
func (j *JWT) generateAccessToken(userID uint64, username string) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.AccessExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.AccessSecret))
}
