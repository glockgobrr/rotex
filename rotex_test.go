package rotex

import (
	"context"
	"github.com/redis/go-redis/v9"
	"reflect"
	"testing"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func TestEncode(t *testing.T) {

	var ctx = context.Background()

	var err = GenAndStoreToRedis(NewRedisClient(), ctx, 2048, "RSA:PRIVATEKEY", "RSA:PUBLICKEY", 3600)

	if err != nil {
		t.Fatalf("expected success , got : %v", err)
	}
}

func TestDecode(t *testing.T) {

	var ctx = context.Background()

	privateKey, err := FetchAndParsePrivateKeyFromRedis(NewRedisClient(), ctx, "RSA:PRIVATEKEY")

	if err != nil {
		t.Fatalf("expected success , got : %v", err)
	}

	publicKey, err := FetchAndParsePublicKeyFromRedis(NewRedisClient(), ctx, "RSA:PUBLICKEY")

	if err != nil {
		t.Fatalf("expected success , got : %v", err)
	}

	if !reflect.DeepEqual(publicKey, &privateKey.PublicKey) {
		t.Fatalf("private key and public key do not match")
	}
}
