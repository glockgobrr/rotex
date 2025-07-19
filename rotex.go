package rotex

import (
	"context"
	"crypto/rsa"
	"github.com/glockgobrr/rotex/fetch"
	"github.com/glockgobrr/rotex/store"
	"time"
	"github.com/redis/go-redis/v9"
)

func GenAndStoreToRedis(redCli *redis.Client, ctx context.Context, bits int, privateKeyId, publicKeyId string, ttls int) error {

	strPrivateKey, strPublicKey, err := store.EncodeKey(bits)
	if err != nil {
		return err
	}

	err = redCli.Set(ctx, privateKeyId, strPrivateKey, time.Duration(ttls)*time.Second).Err()
	if err != nil {
		return err
	}

	err = redCli.Set(ctx, publicKeyId, strPublicKey, time.Duration(ttls)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func FetchAndParsePrivateKeyFromRedis(redCli *redis.Client, ctx context.Context, privateKeyId string) (*rsa.PrivateKey, error) {

	strPrivateKey, err := redCli.Get(ctx, privateKeyId).Result()
	if err != nil {
		return nil, err
	}

	privateKey, err := fetch.DecodePrivateKey(strPrivateKey)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func FetchAndParsePublicKeyFromRedis(redCli *redis.Client, ctx context.Context, publicKeyId string) (*rsa.PublicKey, error) {

	strPublicKey, err := redCli.Get(ctx, publicKeyId).Result()
	if err != nil {
		return nil, err
	}

	publicKey, err := fetch.DecodePublicKey(strPublicKey)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

