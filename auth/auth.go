package auth

import (
	"context"
	"github.com/omgwtflaserguns/matomat-server/db"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var logger = logging.MustGetLogger("log")

const allowedLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EnsureAuthentication(ctx context.Context) codes.Code {

	md, _ := metadata.FromIncomingContext(ctx)

	authCookie := md["cookie"]

	if len(authCookie) == 0 {
		return codes.Unauthenticated
	}

	parts := strings.Split(authCookie[0], "=")

	if len(parts) != 2 {
		return codes.Unauthenticated
	}

	cookieValue := parts[1]

	logger.Debugf("Found Cookie: %s", cookieValue)
	return codes.OK
}

func SetAuthCookie(ctx context.Context, accountId int32) error {
	value := createNewRandomCookieValue()

	_, err := db.DbCon.Exec("INSERT INTO Login (cookie, accountId, created) VALUES ($1, $2, $3)", value, accountId, time.Now())

	if err != nil {
		logger.Errorf("Could not write login to database: %v", err)
		return err
	}

	cookie := &http.Cookie{
		Name:  "matomat-auth",
		Value: value,
		Path:  "/",
	}
	err = grpc.SendHeader(ctx, metadata.New(map[string]string{"Set-Cookie": cookie.String()}))
	return err
}

func createNewRandomCookieValue() string {
	b := make([]byte, 128)
	for i := range b {
		b[i] = allowedLetters[rand.Intn(len(allowedLetters))]
	}
	return string(b)
}
