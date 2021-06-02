package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

func main() {
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
	key := "demo"
	resp, err := api.ReadWorkersKV(context.Background(), os.Getenv("CF_API_WORKER_NAMESPACE_ID"), key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", resp)
	fmt.Println(lsr)
	fmt.Println(u)
}
