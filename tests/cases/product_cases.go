package cases

import "net/http"

// GetProductsTestCase defines the structure for GetProducts test cases
type GetProductsTestCase struct {
	Name          string // Name of the test case
	Page          string // Page number for pagination
	Limit         string // Number of items per page
	ExpectedCode  int    // Expected HTTP status code
	ExpectedCount int    // Expected number of products in response
}

// GetProductsTestCases returns test cases for testing GetProducts endpoint
func GetProductsTestCases() []GetProductsTestCase {
	return []GetProductsTestCase{
		{
			// Test case 1: Successful request with valid pagination
			Name:          "Success",
			Page:          "1",           // First page
			Limit:         "10",          // 10 items per page
			ExpectedCode:  http.StatusOK, // Expect 200 OK
			ExpectedCount: 1,             // Expect 1 product in response
		},
		{
			// Test case 2: Invalid pagination parameters
			Name:         "Invalid Page",
			Page:         "0", // Invalid page number (must be > 0)
			Limit:        "10",
			ExpectedCode: http.StatusBadRequest, // Expect 400 Bad Request
		},
		{
			// Test case 3: Invalid pagination parameters
			Name:         "Invalid Limit",
			Page:         "1",
			Limit:        "0",                   // Invalid limit (must be > 0)
			ExpectedCode: http.StatusBadRequest, // Expect 400 Bad Request
		},
		{
			// Test case 4: Invalid pagination parameters
			Name:         "Invalid Page and Limit",
			Page:         "0",                   // Invalid page number (must be > 0)
			Limit:        "0",                   // Invalid limit (must be > 0)
			ExpectedCode: http.StatusBadRequest, // Expect 400 Bad Request
		},
	}
}
