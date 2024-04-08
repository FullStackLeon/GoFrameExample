package utility

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"GoFrameExample/api/profile/v1"
)

func HashStr(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func EncodeStr(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func GenerateToken(record gdb.Record) (string, error) {
	// 创建一个新的JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":     record["username"],
		"email":        record["email"],
		"account_type": record["account_type"],
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为1天
	})

	// 使用签名密钥签名token
	secretKeyVar, err := g.Config().Get(context.Background(), "jwt.secretKey")
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(secretKeyVar.String()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secretKeyVar, err := g.Config().Get(context.Background(), "jwt.secretKey")
	if err != nil {
		return nil, err
	}
	// 解析JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回签名密钥
		return []byte(secretKeyVar.String()), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// 生成指定长度的字符串
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRandomUserProfile() v1.AuthUserInfo {
	return v1.AuthUserInfo{
		UserId: rand.Intn(10000000),
		Name:   GenerateRandomString(8),
		Age:    rand.Intn(100),
	}
}

func GetRandomProvider() interface{} {
	providers := []string{"Google", "Twitter", "Facebook"}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(providers))
	return providers[index]
}
