package auth

import (
	"context"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

var logger = logging.MustGetLogger("log")

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

func SetAuthCookie(ctx context.Context, userid int32) error {

	cookie := &http.Cookie{
		Name:  "matomat-auth",
		Value: "1927398157987349162397162987192873918732",
		Path:  "/",
	}

	err := grpc.SendHeader(ctx, metadata.New(map[string]string{"Set-Cookie": cookie.String()}))
	return err
}
