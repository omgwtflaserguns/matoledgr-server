package account

import (
	"context"
	"errors"
	"github.com/omgwtflaserguns/matomat-server/config"
	"github.com/omgwtflaserguns/matomat-server/db"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"github.com/op/go-logging"
	"golang.org/x/crypto/bcrypt"
)

var logger = logging.MustGetLogger("log")
var conf = config.GetConfig()

var (
	ErrUserNotFound = errors.New("account: User not found")
	ErrNoRows       = errors.New("account: No rows given")
)

type Service struct{}

func (s *Service) Register(ctx context.Context, in *pb.AccountRequest) (*pb.RegisterResponse, error) {

	if len(in.Password) < conf.Security.MinimalPasswordLength {
		logger.Debugf("Account: Register for user %s failed, password too short", in.Username)
		return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_PASSWORD_INVALID}, nil
	}

	userExists, err := doesUserExist(in.Username)

	if err != nil {
		logger.Fatalf("Account: Register for user %s failed with error: %v", in.Username, err)
		return nil, err
	}

	if userExists {
		logger.Debugf("Account: Register for user %s failed, user exists", in.Username)
		return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_NAME_ALREADY_IN_USE}, nil
	}

	err = createUser(in.Username, in.Password)

	if err != nil {
		logger.Fatalf("Account: Register for user %s failed with error: %v", in.Username, err)
		return nil, err
	}

	logger.Debugf("Account: Register for user %s successful", in.Username)
	return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_OK}, nil
}

func (s *Service) Login(ctx context.Context, in *pb.AccountRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func createUser(username string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), conf.Security.BcryptCost)

	if err != nil {
		return err
	}

	_, err = db.DbCon.Exec("INSERT INTO Account (username, hash) "+
		"VALUES ($1, $2)", username, hash)

	return err
}

func doesUserExist(username string) (bool, error) {
	rows, err := db.DbCon.Query("SELECT id "+
		"FROM Account "+
		"WHERE username = $1", username)

	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

/*
func getUserByUsername(username string) (model.User, error) {
	user := model.User{}
	rows, err := db.DbCon.Query("SELECT id, username, hash " +
		"FROM Account " +
		"WHERE username = $1", username)

	if err != nil {
		return user, err
	}

	user, err = getUserFromRows(rows)

	if err != nil {
		if err == ErrNoRows {
			return user, ErrUserNotFound
		}
		return user, err
	}

	return user, nil
}

func getUserFromRows(rows *sql.Rows) (model.User, error){
	if rows.Next() {

		var id int32
		var username string
		var hash string

		err := rows.Scan(&id, &username, &hash)
		if err != nil {
			return model.User{}, err
		}

		return model.User{Id: id, Username: username, Hash: hash}, nil
	}
	return model.User{}, ErrNoRows
}
*/
