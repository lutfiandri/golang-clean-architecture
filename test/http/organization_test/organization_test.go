package organization_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/lutfiandri/golang-clean-architecture/internal/bootstrap"
	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/infrastructure"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	app *fiber.App
	db  *gorm.DB
)

func bootstrapApp() {
	viperConfig := infrastructure.NewViper("../../../.env.testing")
	config.LoadEnv(viperConfig)

	log := infrastructure.NewLogger()

	validate := infrastructure.NewValidator()
	db = infrastructure.NewDatabase(log)
	app = infrastructure.NewFiber()

	bootstrap.BootstrapApp(bootstrap.BootstrapAppConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
	})
}

func migrateUp() {
	db.Migrator().DropTable(&entity.Organization{})
	db.AutoMigrate(&entity.Organization{})
}

func migrateDown() {
	db.Migrator().DropTable(&entity.Organization{})
}

func seed() {
	tx := db.Begin()
	seedOrganization(tx)
	tx.Commit()
}

func TestMain(m *testing.M) {
	// before test
	bootstrapApp()
	migrateUp()
	seed()

	// test
	m.Run()

	// after test
	migrateDown()
}

// TEST CASES

func TestCreateOrganization(t *testing.T) {
	tests := []struct {
		name           string
		body           map[string]any
		expectedCode   int
		expectedResult map[string]any
	}{
		{
			name: "positive all field",
			body: map[string]any{
				"name":        "test org 1",
				"description": "testing organization 1",
			},
			expectedCode: http.StatusCreated,
			expectedResult: map[string]any{
				"name":        "test org 1",
				"description": "testing organization 1",
			},
		},
		{
			name: "positive without description",
			body: map[string]any{
				"name": "test org 1",
			},
			expectedCode: http.StatusCreated,
			expectedResult: map[string]any{
				"name": "test org 1",
			},
		},
		{
			name:           "negative: without name",
			body:           map[string]any{},
			expectedCode:   http.StatusBadRequest,
			expectedResult: map[string]any{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(test.body)
			assert.NoError(t, err)

			req := httptest.NewRequest("POST", "/organizations", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, -1)

			assert.Equal(t, test.expectedCode, resp.StatusCode)

			if test.expectedCode < 400 {
				// success scenario

				var responseBody map[string]any
				err = json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.NoError(t, err)

				responseData := responseBody["data"].(map[string]any)

				for key, expectedValue := range test.expectedResult {
					assert.Equal(t, expectedValue, responseData[key])
				}
			}
		})
	}
}
