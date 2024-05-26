package repository_test

import (
	"context"
	"regexp"
	"testing"

	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/infrastructure/database"
	"dealls-technical-test-dating-service/internal/infrastructure/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var profileColumns = []string{"id", "user_id", "bio", "interests", "verified", "created_at", "updated_at"}

func TestProfileRepositoryImpl_Insert_Failed_Find_First_Error(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	profile := entity.NewProfile(1, "Bio", "Interests", false, currentTime, currentTime)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "profiles" WHERE "profiles"."user_id" = $1 ORDER BY "profiles"."id" LIMIT $2`)).
		WithArgs(profile.UserID, 1).
		WillReturnError(schema.ErrUnsupportedDataType)

	repo := repository.NewProfileRepository(gormDB)
	err := repo.Insert(context.TODO(), profile)
	require.Error(t, err)
	assert.Equal(t, schema.ErrUnsupportedDataType, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProfileRepositoryImpl_Insert_Failed_Email_Already_Exists(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	profile := entity.NewProfile(1, "Bio", "Interests", false, currentTime, currentTime)
	profileRows := sqlmock.NewRows(profileColumns).AddRow(1, 1, "Bio", "Interests", false, currentTime, currentTime)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "profiles" WHERE "profiles"."user_id" = $1 ORDER BY "profiles"."id" LIMIT $2`)).
		WithArgs(profile.UserID, 1).
		WillReturnRows(profileRows)

	repo := repository.NewProfileRepository(gormDB)
	err := repo.Insert(context.TODO(), profile)
	require.Error(t, err)
	assert.Equal(t, gorm.ErrDuplicatedKey, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProfileRepositoryImpl_Insert_Success(t *testing.T) {
	t.Parallel()

	db, gormDB, mock := database.DBMock(t)
	defer db.Close()

	profile := entity.NewProfile(1, "Bio", "Interests", false, currentTime, currentTime)
	profileRows := sqlmock.NewRows(profileColumns)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "profiles" WHERE "profiles"."user_id" = $1 ORDER BY "profiles"."id" LIMIT $2`)).
		WithArgs(profile.UserID, 1).
		WillReturnRows(profileRows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"profiles\" (.+) VALUES (.+)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	repo := repository.NewProfileRepository(gormDB)
	err := repo.Insert(context.TODO(), profile)
	require.NoError(t, err)
	assert.Equal(t, 1, profile.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
