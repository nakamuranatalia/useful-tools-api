package repository

import (
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/nakamuranatalia/useful-tools-api/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setup() (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error ocurred while creating sql mock: %s", err)
	}

	dialect := postgres.New(postgres.Config{
		Conn:       sqlDb,
		DriverName: "postgres",
	})

	gormDb, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatalf("An error ocurred while creating gorm mock: %s", err)
	}

	return sqlDb, gormDb, mock
}

func TestToolsRepository_SaveTool(tt *testing.T) {
	tt.Run("when there is no tag with the same name, should insert a new tool and tag", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()
		defer sqlDb.Close()

		toolUuid := "d07d32f5-2880-4cac-97d3-0b2fd871b3fa"
		toolId := 1
		tagId := 1

		tool := &model.Tool{
			Title:       "GoLand",
			Link:        "https://www.jetbrains.com",
			Description: "The Go IDE that truly understands the language.",
			Tags: []model.Tag{
				{
					Name: "IDE",
				},
			},
		}

		mock.ExpectQuery(`SELECT (.+) FROM "tags"`).
			WillReturnRows(mock.NewRows([]string{"id", "title"}).AddRow(nil, nil))
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "tools"`).WithArgs(
			"GoLand",
			"https://www.jetbrains.com",
			"The Go IDE that truly understands the language.",
		).WillReturnRows(sqlmock.NewRows([]string{"uuid", "id"}).AddRow(toolUuid, toolId))
		mock.ExpectQuery(`INSERT INTO "tags"`).
			WithArgs("IDE").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tagId))
		mock.ExpectExec(`INSERT INTO "tool_tag"`).
			WithArgs(toolId, tagId).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.SaveTool(tool)

		//ASSERT
		assert.Equal(t, uuid.MustParse(toolUuid), result.Uuid)
		assert.Equal(t, tool.Title, result.Title)
		assert.Equal(t, tool.Link, result.Link)
		assert.Equal(t, tool.Description, result.Description)
		assert.Equal(t, tool.Tags[0].Name, result.Tags[0].Name)
		assert.Equal(t, uint(tagId), result.Tags[0].Id)
		assert.Nil(t, err)
	})

	tt.Run("when a tag already exists, should insert only a new tool", func(t *testing.T) {
		//ARRANGE
		var (
			sqlDb, gormDb, mock = setup()
			toolUuid            = "72a7e6a2-0a31-4eb1-ad5a-fc6f7e2b151e"
			toolId              = 1
			tagId               = 2
			tool                = &model.Tool{
				Title:       "VsCode",
				Link:        "https://www.code.visualstudio.com",
				Description: "The open source AI code editor.",
				Tags: []model.Tag{
					{
						Name: "IDE",
					},
				},
			}
		)

		mock.ExpectQuery(`SELECT (.+) FROM "tags"`).
			WillReturnRows(mock.NewRows([]string{"id", "title"}).AddRow(tagId, "IDE"))
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "tools"`).WithArgs(
			"VsCode",
			"https://www.code.visualstudio.com",
			"The open source AI code editor.",
		).WillReturnRows(sqlmock.NewRows([]string{"uuid", "id"}).AddRow(toolUuid, toolId))
		mock.ExpectQuery(`INSERT INTO "tags"`).
			WithArgs("IDE", tagId).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tagId))
		mock.ExpectExec(`INSERT INTO "tool_tag"`).
			WithArgs(toolId, tagId).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.SaveTool(tool)

		//ASSERT
		assert.Equal(t, uuid.MustParse(toolUuid), result.Uuid)
		assert.Equal(t, tool.Title, result.Title)
		assert.Equal(t, tool.Link, result.Link)
		assert.Equal(t, tool.Description, result.Description)
		assert.Equal(t, tool.Tags[0].Name, result.Tags[0].Name)
		assert.Equal(t, uint(tagId), result.Tags[0].Id)
		assert.Nil(t, err)
	})
}

func TestToolRepository_FindTools(tt *testing.T) {
	tt.Run("when there is no tools, should return empty", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()

		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "title", "link", "description", "tags"}))

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.FindTools()

		//ASSERT
		assert.Empty(t, result)
		assert.Nil(t, err)
	})
	tt.Run("when there is tools, should return tools that are persisted", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()

		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "title", "link"}).
				AddRow("1", "e978cbdc-4634-469f-8d34-67e0ee828a74", "Google", "google.com"))
		mock.ExpectQuery(`SELECT (.+) FROM "tool_tag"`).
			WillReturnRows(mock.NewRows([]string{"id_tool", "id_tag"}).
				AddRow("1", "1"))
		mock.ExpectQuery(`SELECT (.+) FROM "tags"`).
			WillReturnRows(mock.NewRows([]string{"id", "name"}).
				AddRow("1", "Search"))

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.FindTools()

		//ASSERT
		assert.Equal(t, 1, len(result))
		assert.Nil(t, err)
	})
}

func TestToolsRepository_FindToolByUuid(tt *testing.T) {
	tt.Run("when there is no tool with the passed uuid, should return empty", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()

		toolUuid := "4d710234-eb36-4c40-ba93-0a829be8fbdf"

		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "title", "link"}).
				AddRow(nil, nil, nil, nil))
		mock.ExpectQuery(`SELECT (.+) FROM "tool_tag"`).
			WillReturnRows(mock.NewRows([]string{"id_tool", "id_tag"}).
				AddRow(nil, nil))
		mock.ExpectQuery(`SELECT (.+) FROM "tags"`).
			WillReturnRows(mock.NewRows([]string{"id", "name"}).
				AddRow(nil, nil))

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.FindToolByUuid(toolUuid)

		//ASSERT
		assert.Empty(t, result)
		assert.Nil(t, err)

	})
	tt.Run("when there is a tool with the passed uuid, should return the specific tool", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()

		toolUuid := "4d710234-eb36-4c40-ba93-0a829be8fbdf"

		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "title", "link"}).
				AddRow("1", toolUuid, "Postman", "www.postman.com"))
		mock.ExpectQuery(`SELECT (.+) FROM "tool_tag"`).
			WillReturnRows(mock.NewRows([]string{"id_tool", "id_tag"}).
				AddRow("1", "2"))
		mock.ExpectQuery(`SELECT (.+) FROM "tags"`).
			WillReturnRows(mock.NewRows([]string{"id", "name"}).
				AddRow("2", "API platform"))

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.FindToolByUuid(toolUuid)

		//ASSERT
		assert.Equal(t, uuid.MustParse(toolUuid), result.Uuid)
		assert.Nil(t, err)
	})
}

func TestToolsRepository_DeleteToolByUuid(tt *testing.T) {
	tt.Run("when there is no tool with given uuid, should return error", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()

		toolUuid := "aac72494-f4e6-4aba-af26-0701a5021d63"
		expectedError := errors.New("record not found")
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnError(errors.New("record not found"))

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		err := repository.DeleteToolByUuid(toolUuid)

		//ASSERT
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
	tt.Run("when there is a tool with the passed uuid, should delete this tool", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()

		toolUuid := "bc289273-a69a-44ad-8ed7-335636889030"

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "itle", "link"}).
				AddRow(3, toolUuid, "Udemy", "www.udemy.com"))
		mock.ExpectExec(`DELETE FROM "tool_tag" WHERE "tool_tag"\."tool_id" = \$1`).
			WithArgs(3).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec(`DELETE FROM "tools" WHERE "tools"\."id" = \$1`).
			WithArgs(3).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		err := repository.DeleteToolByUuid(toolUuid)

		//ASSERT
		assert.Nil(t, err)
	})
}

func TestToolsRepository_UpdateTool(tt *testing.T) {
	tt.Run("when there is no tool with given uuid, should return error when trying to delete", func(t *testing.T) {
		//ARRANGE
		sqlDb, gormDb, mock := setup()

		toolUuid := "4d710234-eb36-4c40-ba93-0a829be8fbdf"
		expectedError := errors.New("record not found")

		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "title", "link"}).
				AddRow(nil, nil, nil, nil))
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnError(errors.New("record not found"))

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.UpdateTool(&model.Tool{}, toolUuid)

		//ASSERT
		assert.Empty(t, result)
		assert.Equal(t, expectedError, err)
	})
	tt.Run("when there is a tool with given uuid, should update the tool", func(t *testing.T) {
		//ASSERT
		sqlDb, gormDb, mock := setup()

		toolUuid := "2d6e1351-d6b1-4f73-822e-a53b2af40820"
		toolToUpdate := &model.Tool{
			Title:       "Copilot",
			Link:        "www.github.com/features/copilot",
			Description: "Generative artificial intelligence chatbot developed by Microsoft",
			Tags: []model.Tag{
				{
					Name: "IA",
				},
			},
		}

		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "title", "link"}).
				AddRow("1", toolUuid, "Copilot", "www.github.com"))
		mock.ExpectQuery(`SELECT (.+) FROM "tool_tag"`).
			WillReturnRows(mock.NewRows([]string{"id_tool", "id_tag"}).
				AddRow("1", "1"))

		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT (.+) FROM "tools"`).
			WillReturnRows(mock.NewRows([]string{"id", "uuid", "itle", "link"}).
				AddRow(1, toolUuid, "Copilot", "www.github.com"))
		mock.ExpectExec(`DELETE FROM "tool_tag" WHERE "tool_tag"\."tool_id" = \$1`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec(`DELETE FROM "tools" WHERE "tools"\."id" = \$1`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(`SELECT (.+) FROM "tags"`).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "IA"))
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "tools"`).WithArgs(
			"Copilot",
			"www.github.com/features/copilot",
			"Generative artificial intelligence chatbot developed by Microsoft",
			toolUuid,
			1,
		).WillReturnRows(sqlmock.NewRows([]string{"uuid", "id"}).AddRow(toolUuid, 1))
		mock.ExpectQuery(`INSERT INTO "tags"`).
			WithArgs("IA", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectExec(`INSERT INTO "tool_tag"`).
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		defer sqlDb.Close()

		//ACT
		repository := NewRepository(gormDb)
		result, err := repository.UpdateTool(toolToUpdate, toolUuid)

		//ASSERT
		assert.Equal(t, toolToUpdate.Title, result.Title)
		assert.Equal(t, toolToUpdate.Link, result.Link)
		assert.Equal(t, toolToUpdate.Description, result.Description)
		assert.Equal(t, toolToUpdate.Tags[0].Name, result.Tags[0].Name)
		assert.Nil(t, err)
	})
}
