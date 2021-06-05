package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
	uuid "github.com/google/uuid"
)

type Metadata struct {
	RedirectUrl string `json:"redirect_url"`
}

func main() {
	key := uuid.NewString()
	log.Println(key)

	redirectUrl := fmt.Sprintf("https://twomenus.io/menu/%s/categories", key)
	value, err := json.Marshal(&Metadata{
		RedirectUrl: redirectUrl,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(value)

	// Construct a new API object
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"), cloudflare.UsingAccount(os.Getenv("CF_API_ACCOUNT_ID")))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// list KV
	lsr, err := api.ListWorkersKVs(context.Background(), os.Getenv("CF_API_WORKER_NAMESPACE_ID"))
	if err != nil {
		log.Fatal(err)
	}

	// read kv
	key = "demo"
	respRead, err := api.ReadWorkersKV(context.Background(), os.Getenv("CF_API_WORKER_NAMESPACE_ID"), key)
	if err != nil {
		log.Fatal(err)
	}

	payload := []byte("test")
	key = "test"

	respCreate, err := api.WriteWorkersKV(context.Background(), os.Getenv("CF_API_WORKER_NAMESPACE_ID"), key, payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", respRead)
	fmt.Println(respCreate)

	fmt.Println(lsr)
	fmt.Println(u)
}
