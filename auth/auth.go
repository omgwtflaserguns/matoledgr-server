package auth

import (
	"context"
	"github.com/omgwtflaserguns/matomat-server/db"
	"github.com/omgwtflaserguns/matomat-server/model"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var logger = logging.MustGetLogger("log")

const allowedLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EnsureAuthentication(ctx context.Context) (model.Login, error) {

	md, _ := metadata.FromIncomingContext(ctx)

	authCookie := md["cookie"]

	if len(authCookie) == 0 {
		return model.Login{}, status.Error(codes.Unauthenticated, "No auth cookie found, please login")
	}

	parts := strings.Split(authCookie[0], "=")

	if len(parts) != 2 {
		return model.Login{}, status.Error(codes.Unauthenticated, "No auth cookie found, please login")
	}

	cookieValue := parts[1]

	maxAge := time.Now().AddDate(0, 0, -1)

	row := db.DbCon.QueryRow("SELECT a.id, a.username, a.hash, l.cookie, l.created "+
		"FROM Login l "+
		"INNER JOIN Account a "+
		"ON a.id = l.accountId "+
		"WHERE l.cookie = $1 AND l.created > $2", cookieValue, maxAge)

	usr := model.User{}
	login := model.Login{}
	login.User = usr
	err := row.Scan(&login.User.Id, &login.User.Username, &login.User.Hash, &login.Cookie, &login.Created)
	if err != nil {
		return model.Login{}, status.Error(codes.Unauthenticated, "No auth cookie found, please login")
	}

	logger.Debugf("Got login from db for account %s for cookie %s", login.User.Username, cookieValue)
	return login, nil
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
