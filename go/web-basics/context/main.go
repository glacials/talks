package main

import (
	"context"
	"errors"
	"log"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", "123")
	doAThingThatMightBreak(ctx)
}

func doAThingThatMightBreak(ctx context.Context) {
	err := doAnotherThing()
	if err != nil {
		log.Printf("(user_id=%s) The thing broke: %s", ctx.Value("user_id"), err)
	}
}

func doAnotherThing() error {
	return errors.New("beep")
}
