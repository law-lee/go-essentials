package contextdemo

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Name       string
	IsLoggedIn bool
}

type userKeyType int

var userKey userKeyType

func contextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// returns nil if not set
func getUserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(userKey).(*User)
	if !ok {
		return nil
	}
	return user
}

// will panic if not set
func mustGetUserFromContext(ctx context.Context) *User {
	return ctx.Value(userKey).(*User)
}

func printUser(ctx context.Context) {
	user := getUserFromContext(ctx)
	fmt.Printf("User: %#v\n", user)
}

func longMathOp(ctx context.Context, n int) (int, error) {
	res := n
	for i := 0; i < 1000; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			res += i
			// simulate long operation by sleeping
			time.Sleep(time.Millisecond)
		}
	}
	return res, nil
}

func Run() {
	ctx := context.Background()
	user := &User{
		Name:       "John",
		IsLoggedIn: false,
	}
	ctx = contextWithUser(ctx, user)

	printUser(ctx)

	ctx1, _ := context.WithTimeout(context.Background(), time.Millisecond*2000)
	res, err := longMathOp(ctx1, 5)
	fmt.Printf("Called longMathOp() with 200ms timeout. res; %d, err: %v\n", res, err)

	ctx1, _ = context.WithTimeout(context.Background(), time.Millisecond*100)
	res, err = longMathOp(ctx1, 5)
	fmt.Printf("Called longMathOp() with 10ms timeout. res: %d, err: %v\n", res, err)
}
