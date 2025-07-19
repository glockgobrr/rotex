
# Project Title

Rotex is a utility tool that generates rsa private and public key , stores it and fetches it from the user designated redis client.


## Features

- A key generator that creates RSA private and public keys, marshals them, and PEM-encodes them for secure storage.
- A key parser that accepts PEM-encoded keys, decodes them, and returns usable RSA key objects.
- A Redis client placeholder that accepts an active Redis client, making the tool easy to integrate and work with. 

 
## Installation

Import my project as a library

```go
  go get -u "github.com/glockgobrr/rotex"
```
## Usage/Examples

// For generating and storing PEM-encoded key in redis

```go
import (
    "log"
    "context"
    rotex "github.com/glockgobrr/rotex"
    "github.com/redis/go-redis/v9"

func RedisClient() *redis.Client {
  return &redis.Client(&redis.Options{
      Addr : "localhost:6379",
  })
}

func main() {

    ctx = context.Background()
    redCli = RedisClient()

    // Function to generate and store RSA keys in Redis
    err := rotex.GenAndStoreToRedis(
        redCli, ctx, bits, privateKeyId, publicKeyId, ttls                               
        ) 

    // redCli : Active Redis Client 
    // ctx : context
    // bits : size of key you want to generate int bits (2048 or 4096)
    // privateKeyId : key id you want to store your private key with (in string)
    // publicKeyId : key id you want to store your public key with (in string)
    // ttls : time to live in seconds , put 0 for no expiry 
    if err != nil {
        log.Printf("failed to generate and store rsa key pair : %v" , err)
    }


}
```
// For Fetching Private Key from redis and PEM-decoding it. 

```go
import (
    "log"
    "context"
    rotex "github.com/glockgobrr/rotex"
    "github.com/redis/go-redis/v9"

func RedisClient() *redis.Client {
  return &redis.Client(&redis.Options{
      Addr : "localhost:6379",
  })
}

func main() {

    ctx = context.Background()
    redCli = RedisClient()

    // Function to fetch and decode RSA Private key from Redis
    privateKey , err := FetchAndParsePrivateKeyFromRedis(
        redCli, ctx, privateKeyId string
        ) 
    }

    if err != nil {
        log.Printf("failed to fetch and decode rsa private key : %v" , err)
    }

}

```

// For Fetching Public Key from redis and PEM-decoding it.


```go
import (
    "log"
    "context"
    rotex "github.com/glockgobrr/rotex"
    "github.com/redis/go-redis/v9"

func RedisClient() *redis.Client {
  return &redis.Client(&redis.Options{
      Addr : "localhost:6379",
  })
}

func main() {

    ctx = context.Background()
    redCli = RedisClient()

    // Function to fetch and decode RSA Public key from Redis
    publicKey , err := FetchAndParsePublicKeyFromRedis(redCli, ctx, publicKeyId)
    }

    if err != nil {
        log.Printf("failed to fetch and decode rsa public key : %v" , err)
    }

}

```
## License

This project is licensed under the [MIT License](./LICENSE).
