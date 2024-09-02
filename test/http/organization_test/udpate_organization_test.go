package organization_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCreateOrganization(t *testing.T) {
	tests := []struct {
		id             int
		name           string
		body           map[string]any
		expectedCode   int
		expectedResult map[string]any
	}{
		{
			id:   1,
			name: "positive all field",
			body: map[string]any{
				"name":        "org 1 - update",
				"description": "description 1 - update",
			},
			expectedCode: http.StatusOK,
			expectedResult: map[string]any{
				"name":        "org 1 - update",
				"description": "description 1 - update",
			},
		},
		{
			id:   2,
			name: "positive without description",
			body: map[string]any{
				"name": "org 2 - update",
			},
			expectedCode: http.StatusOK,
			expectedResult: map[string]any{
				"name": "org 2 - update",
			},
		},
		{
			id:             3,
			name:           "negative without name",
			body:           map[string]any{},
			expectedCode:   http.StatusBadRequest,
			expectedResult: map[string]any{},
		},
		{
			id:   100,
			name: "negative not found",
			body: map[string]any{
				"name":        "org 100 - update",
				"description": "description 100 - update",
			},
			expectedCode:   http.StatusNotFound,
			expectedResult: map[string]any{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(test.body)
			assert.NoError(t, err)

			req := httptest.NewRequest("PUT", fmt.Sprintf("/organizations/%d", test.id), bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, -1)

			assert.Equal(t, test.expectedCode, resp.StatusCode)

			if test.expectedCode >= 400 {
				// error scenario
				return
			}

			// success scenario
			var responseBody map[string]any
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			responseData := responseBody["data"].(map[string]any)

			organizationID, idExists := responseData["id"]
			assert.Truef(t, idExists, "response `id` not found")

			for key, expectedValue := range test.expectedResult {
				assert.Equal(t, expectedValue, responseData[key])
			}

			id := uint(organizationID.(float64))

			organizationInDB := new(entity.Organization)
			if err := db.Where("id = ?", id).First(organizationInDB).Error; err != nil {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, test.expectedResult["name"], organizationInDB.Name)
			if _, exists := test.expectedResult["description"]; exists {
				assert.EqualValues(t, test.expectedResult["description"], *organizationInDB.Description)
			}
			assert.NotNil(t, organizationInDB.CreatedAt)
			assert.NotNil(t, organizationInDB.UpdatedAt)
		})
	}
}
