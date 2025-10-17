package main

import (
	"context"
	"encoding/json"
	"fmt"
)

var users []*User = []*User{
	{
		Username: "jarwo_kuwat",
		Password: hashSHA1("secret123"),
		Email:    "jarwo@gmail.com",
		RealName: "Jarwo Kuwat",
	},
	{
		Username: "sule_prikitiw",
		Password: hashSHA1("secret123"),
		Email:    "sule@gmail.com",
		RealName: "Sule Sutisna",
	},
}

func (a *application) seedRedis() {
	fmt.Println("Seeding...")
	m := make(map[string]any, 0)

	for _, u := range users {
		b, err := json.Marshal(u)
		if err != nil {
			panic(err)
		}

		k := "login_" + u.Username
		m[k] = string(b)
	}

	if _, err := a.rdb.MSet(context.Background(), m).Result(); err != nil {
		panic(err)
	}
	fmt.Println("Seeding Complete...")
}
