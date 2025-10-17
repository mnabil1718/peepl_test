package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func (a *application) loginHandler(c fiber.Ctx) error {
	r := new(LoginRequest)

	if err := c.Bind().Body(r); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dummyHash := hashSHA1("dummy_pass")
	u := &User{
		Username: r.Username,
		Password: fmt.Sprintf("%x", dummyHash), // dummy, prevent timing attack
	}

	val, err := a.rdb.Get(c.Context(), fmt.Sprintf("login_%s", u.Username)).Result()

	if err != nil && err != redis.Nil {
		return err
	}

	if err == nil {
		if err = json.Unmarshal([]byte(val), u); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	fmt.Println(u)

	// constant time comparison
	if subtle.ConstantTimeCompare(
		[]byte(hashSHA1(r.Password)),
		[]byte(u.Password),
	) != 1 {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	return c.Status(fiber.StatusOK).JSON(HandlerResponse{
		Success: true,
		Message: "login successful",
	})
}
