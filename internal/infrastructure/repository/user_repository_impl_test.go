package repository_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/infrastructure/database"
	"dealls-technical-test-dating-service/internal/infrastructure/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	currentTime     = time.Now()
	birthDateFormat = "2006-01-02"
	birthDate       = currentTime.Format(birthDateFormat)
	userColumns     = []string{"id", "email", "password", "name", "birth_date", "gender", "location", "profile_picture_url", "created_at", "updated_at"}
)

func TestUserRepositoryImpl_Insert_Failed_Find_First_Error(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	user := entity.NewUser("user@email.com", "password", "User", birthDate, "MALE", "Indonesia", "", currentTime, currentTime)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."email" = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(user.Email, 1).
		WillReturnError(schema.ErrUnsupportedDataType)

	repo := repository.NewUserRepository(gormDB)
	_, err := repo.Insert(context.TODO(), user)
	require.Error(t, err)
	assert.Equal(t, schema.ErrUnsupportedDataType, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_Insert_Failed_Email_Already_Exists(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	user := entity.NewUser("user@email.com", "password", "User", birthDate, "MALE", "Indonesia", "", currentTime, currentTime)
	birthDate, _ := time.Parse(birthDateFormat, birthDate)
	userRows := sqlmock.NewRows(userColumns).AddRow(1, "user@email.com", "password", "User", birthDate, "MALE", "Indonesia", "", currentTime, currentTime)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."email" = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(user.Email, 1).
		WillReturnRows(userRows)

	repo := repository.NewUserRepository(gormDB)
	id, err := repo.Insert(context.TODO(), user)
	require.Error(t, err)
	assert.Equal(t, 0, id)
	assert.Equal(t, gorm.ErrDuplicatedKey, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_Insert_Success(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	user := entity.NewUser("user@email.com", "password", "User", birthDate, "MALE", "Indonesia", "", currentTime, currentTime)
	userRows := sqlmock.NewRows(userColumns)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."email" = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(user.Email, 1).
		WillReturnRows(userRows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"users\" (.+) VALUES (.+)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	repo := repository.NewUserRepository(gormDB)
	_, err := repo.Insert(context.TODO(), user)
	require.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_FindByEmail_Failed_Not_Found(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	expectedSQL := "SELECT (.+) FROM \"users\" WHERE email = (.+)"
	userRows := sqlmock.NewRows(userColumns)
	mock.ExpectQuery(expectedSQL).WillReturnRows(userRows)

	repo := repository.NewUserRepository(gormDB)
	_, err := repo.FindByEmail(context.TODO(), "user@email.com")
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepositoryImpl_FindByEmail_Success(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	birthDate, _ := time.Parse(birthDateFormat, birthDate)
	userRows := sqlmock.NewRows(userColumns).AddRow(1, "user@email.com", "password", "User", birthDate, "MALE", "Indonesia", "", currentTime, currentTime)

	expectedSQL := "SELECT (.+) FROM \"users\" WHERE email = (.+)"
	mock.ExpectQuery(expectedSQL).WillReturnRows(userRows)

	repo := repository.NewUserRepository(gormDB)
	_, err := repo.FindByEmail(context.TODO(), "user@email.com")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
