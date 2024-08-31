package organization_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/lutfiandri/golang-clean-architecture/internal/bootstrap"
	"github.com/lutfiandri/golang-clean-architecture/internal/config"
	"github.com/lutfiandri/golang-clean-architecture/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

var app *fiber.App

func beforeAll() {
	viperConfig := infrastructure.NewViper("../../../.env.testing")
	config.LoadEnv(viperConfig)

	log := infrastructure.NewLogger()

	db := infrastructure.NewDatabase(log)
	validate := infrastructure.NewValidator()
	app = infrastructure.NewFiber()

	bootstrap.BootstrapApp(bootstrap.BootstrapAppConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
	})

	// time.Sleep(5 * time.Second)
}

func afterAll() {
	// db.Migrator().DropTable(&entity.Organization{})
	log.Println("After All")
}

func TestMain(m *testing.M) {
	beforeAll()
	m.Run()
	afterAll()
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
			name: "success: all field",
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
			name: "success: without description",
			body: map[string]any{
				"name": "test org 1",
			},
			expectedCode: http.StatusCreated,
			expectedResult: map[string]any{
				"name": "test org 1",
			},
		},
		{
			name:           "success: without name",
			body:           map[string]any{},
			expectedCode:   http.StatusBadRequest,
			expectedResult: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.body)
			assert.NoError(t, err)

			req := httptest.NewRequest("POST", "/organizations", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, -1)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			// test fields
			var responseBody map[string]any
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			if tt.expectedCode < 400 {
				// success scenario
				responseData := responseBody["data"].(map[string]any)

				log.Printf("response: %+v", responseBody)

				for key, expectedValue := range tt.expectedResult {
					assert.Equal(t, expectedValue, responseData[key])
				}
			} else {
				// error scenario
				// assert.Equal(t, tt.expectedResult["message"], responseBody["message"])
			}
		})
	}
}
