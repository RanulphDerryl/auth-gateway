package auth_gateway

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// GenerateRSAKeyPair generates a new RSA key pair and writes it to a file.
func GenerateRSAKeyPair(ctx context.Context, privKeyPath, pubKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	privatePem := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	privatePEMBytes := pem.EncodeToMemory(privatePem)
	if err := os.WriteFile(privKeyPath, privatePEMBytes, 0600); err != nil {
		return err
	}

	pubKey := &rsa.PublicKey{N: privateKey.N, E: privateKey.E}
	publicPem := &pem.Block{Type: "PUBLIC KEY", Bytes: x509.MarshalPKIXPublicKey(pubKey)}
	publicPEMBytes := pem.EncodeToMemory(publicPem)
	if err := os.WriteFile(pubKeyPath, publicPEMBytes, 0644); err != nil {
		return err
	}

	return nil
}

// GetAvailablePort gets the first available port number on the local machine.
func GetAvailablePort() (int, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

// CreateRedisClient creates a Redis client instance with the given options.
func CreateRedisClient(options *redis.Options) *redis.Client {
	return redis.NewClient(options)
}

// CreateClient creates a new client instance for the Auth Gateway.
func CreateClient(ctx context.Context) (*RedisClient, error) {
	privateKeyPath := os.Getenv("AUTH_GATEWAY_PRIVATE_KEY_PATH")
	publicKeyPath := os.Getenv("AUTH_GATEWAY_PUBLIC_KEY_PATH")
	if privateKeyPath == "" || publicKeyPath == "" {
		return nil, fmt.Errorf("missing environment variable")
	}

	if err := GenerateRSAKeyPair(ctx, privateKeyPath, publicKeyPath); err != nil {
		return nil, err
	}

	privateKeyPath = strings.TrimPrefix(privateKeyPath, os.Getenv("HOME")+"/")
	publicKeyPath = strings.TrimPrefix(publicKeyPath, os.Getenv("HOME")+"/")

	options := &redis.Options{
		Addr:         "localhost:6379",
		Password:     "", // no password set
		DB:           0,  // use default DB
		ReadTimeout:  10 * time.Second,
	}

	return CreateRedisClient(options)
}

// RedisClient represents a Redis client instance for the Auth Gateway.
type RedisClient struct {
	*redis.Client
}

func (c *RedisClient) Get(ctx context.Context, key string) (*string, error) {
	return c.Client.Get(ctx, key).Result()
}

func (c *RedisClient) Expire(ctx context.Context, key string, seconds int64) error {
	return c.Client.Expire(ctx, key, time.Duration(seconds)*time.Second).Err()
}