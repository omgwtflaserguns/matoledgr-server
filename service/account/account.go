package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/omgwtflaserguns/matomat-server/auth"
	"github.com/omgwtflaserguns/matomat-server/config"
	"github.com/omgwtflaserguns/matomat-server/db"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"github.com/omgwtflaserguns/matomat-server/model"
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
		return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_FAILED_PASSWORD_INVALID}, nil
	}

	userExists, err := doesUserExist(in.Username)

	if err != nil {
		logger.Errorf("Account: Register for user %s failed, doesUserExist returned error: %v", in.Username, err)
		return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_FAILED}, nil
	}

	if userExists {
		logger.Debugf("Account: Register for user %s failed, user exists", in.Username)
		return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_FAILED_NAME_ALREADY_IN_USE}, nil
	}

	err = createUser(in.Username, in.Password)

	if err != nil {
		logger.Errorf("Account: Register for user %s failed, create User returned error: %v", in.Username, err)
		return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_FAILED}, nil
	}

	logger.Debugf("Account: Register for user %s successful", in.Username)
	return &pb.RegisterResponse{Status: pb.RegisterStatus_REGISTER_OK}, nil
}

func (s *Service) Login(ctx context.Context, in *pb.AccountRequest) (*pb.LoginResponse, error) {
	usr, err := getUserByUsername(in.Username)

	if err != nil {
		if err == ErrUserNotFound {
			logger.Debugf("Account: Login for user %s failed, no user found", in.Username)
		} else {
			logger.Fatalf("Account: Login for user %s failed, get User returned error: %v", in.Username, err)
		}
		return &pb.LoginResponse{Status: pb.LoginStatus_LOGIN_FAILED}, nil
	}

	err = bcrypt.CompareHashAndPassword(usr.Hash, []byte(in.Password))

	if err != nil {
		logger.Debugf("Account: Login for user %s failed, bcrypt returned error: %v", in.Username, err)
		return &pb.LoginResponse{Status: pb.LoginStatus_LOGIN_FAILED}, nil
	}

	logger.Debugf("Account: Login for user %s successful", in.Username)

	err = auth.SetAuthCookie(ctx, usr.Id)
	if err != nil {
		logger.Errorf("Account: Setting cookie for user %s failed, error: %v", in.Username, err)
	}

	return &pb.LoginResponse{
		Status: pb.LoginStatus_LOGIN_OK,
		User:   model.GetProtoUserFromUser(usr),
	}, nil
}

func (s *Service) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {

	login, err := auth.EnsureAuthentication(ctx)

	if err != nil {
		return &pb.GetAccountResponse{Authenticated: false, User: nil}, nil
	}

	logger.Debugf("Account: GetAccount for user %s, found authentication", login.User.Username)
	return &pb.GetAccountResponse{
		Authenticated: true,
		User:          model.GetProtoUserFromUser(login.User),
	}, nil
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
	defer rows.Close()

	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

func getUserByUsername(username string) (model.User, error) {
	user := model.User{}
	rows, err := db.DbCon.Query("SELECT id, username, hash "+
		"FROM Account "+
		"WHERE username = $1", username)
	defer rows.Close()

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

func getUserFromRows(rows *sql.Rows) (model.User, error) {
	if rows.Next() {
		usr := model.User{}
		err := rows.Scan(&usr.Id, &usr.Username, &usr.Hash)
		if err != nil {
			return model.User{}, err
		}
		return usr, nil
	}
	return model.User{}, ErrNoRows
}
