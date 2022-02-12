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

	first := 1
	for _, username := range usernames {
		MustCreateUser(ctx, client, username)
	}
	conn1 := MustQueryUser(ctx, client, &first, nil, nil, nil)
	fmt.Println(formatJson(conn1))
	conn2 := MustQueryUser(ctx, client, &first, nil, nil, conn1.PageInfo.EndCursor)
	fmt.Println(formatJson(conn2))
	conn3 := MustQueryUser(ctx, client, &first, nil, nil, conn2.PageInfo.EndCursor)
	fmt.Println(formatJson(conn3))
}

func MustCreateUser(ctx context.Context, client *ent.Client, name string) *ent.User {
	u, err := client.User.
		Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		panic(fmt.Errorf("failed creating user: %w", err))
	}
	return u
}

func MustQueryUser(ctx context.Context, client *ent.Client, first, last *int, before, after *ent.Cursor) *ent.UserConnection {
	c, err := client.User.
		Query().
		Paginate(ctx, after, first, before, last)
	if err != nil {
		panic(fmt.Errorf("failed querying user: %w", err))
	}
	return c
}

func formatJson(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
