package model_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lanngoen1996/golang-basic/mock"
	model "github.com/lanngoen1996/golang-basic/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserModelTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	db   *sql.DB
}

func (suite *UserModelTestSuite) SetupTest() {
	mock, db := mock.MockConn()

	suite.mock = mock
	suite.db = db
}

func (suite *UserModelTestSuite) AfterTest() {
	defer suite.db.Close()
}

func (s *UserModelTestSuite) TestFindUser() {
	user := &model.User{}

	q := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1")
	row := sqlmock.NewRows([]string{
		"id", "username", "created_at", "updated_at",
	}).AddRow(1, "username", time.Now(), time.Now())

	s.mock.ExpectQuery(q).
		WithArgs("1").
		WillReturnRows(row)

	m := model.NewUserModel()
	m.Find("1", user)

	assert.Equal(s.T(), uint(1), user.ID)
	assert.Equal(s.T(), "username", user.Username)
}

func TestUserModelTestSuite(t *testing.T) {
	suite.Run(t, new(UserModelTestSuite))
}
