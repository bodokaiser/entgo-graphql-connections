package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"entgo-graphql-connections/ent"
)

var usernames = []string{"a8m", "bodo", "yoni"}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()

	// create users
	for _, username := range usernames {
		if _, err := CreateUser(ctx, client, username); err != nil {
			log.Fatalf("failed creating user: %s", username)
		}
	}

	// query users
	users, err := QueryUser(ctx, client)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	fmt.Println(printJson(users))
}

func CreateUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func printJson(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
