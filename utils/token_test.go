package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func Test_CreateToken(t *testing.T) {
	testTable := []struct {
		name          string
		ttl           time.Duration
		payload       interface{}
		privateKey    string
		expectedError error
	}{
		{
			name:          "Ok",
			ttl:           time.Minute * 15,
			payload:       1,
			privateKey:    "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWFFJQkFBS0JnUUNVNjd1dmpQTXlNallGTnd6YlVVTEJMemxadHJuWlF1WXpQdEZOVnUwK1RqblBNekczCjFSUWkvYTlHenBUaTFHZDJHVmRhSkpLWU1JZXZIZURpU3IveVloWVA5UmNXOWtrdytrNjJ4Zis3VlkxbFlCTWUKNXVBbTAzYUpIL0NZRzlYSStOMUNxOWtvZ1JvdG5oUExranNmMTRUTS9MYXEwVDNMVTU0QzBnR0hnUUlEQVFBQgpBb0dBTGVvWWJlRzVRYXExZVJvbC9xQ3BRa0s3SGV2MmxRUEpVZGQyVkxBT2FYMVkyWWRoTnRxZFlNUnlmQlJKClZ6NUJ2K25FUXVpdndmaUVmUHRyVUpuWFgvdFpzM1pteVJJZVBtNjA0TEhsajg3ZGFrV081WUNNZDA0aDNWNWoKTjQ1a3EybkNRUi92ZXdLUEF6dFN3R2FuaHB6UlQ0QmtxV080cVRBcTdkNVdvNkVDUVFEclNoenc4MUxyc0tYbgpXUUJrQU55OENtRzdyRStCdUNLeFVOYnhaeU9ISkVQK3BkRmExZWRsK0NheXpycmpkWjlIeGhOdzllcG1uOGR0CktwMDJrREI5QWtFQW9nZHZXbmR4SHVhdWlXQ3ppUGEyejkrd25vaVMxNGw0eVY2d2xRMFhuZkIyeURvWGtETlYKOWFHQ25xdVFvSHN4TWdSYU12QUpkc2VDVGxiVzRjQ0dWUUpCQUtjS1M3ZW9GNE5xT3E0ZTJvOGtxWkQ2RWQ2SgorendOdk51RGw4VjBRcDNMMmxYcjVOQ0hNYXVMMi9Wdm5QQ2s3YnFuM2QrQlJyTXExZ3NqaU05VzJJVUNRRzY4CjlCNzVXU2ZNYzJkUzN3RngxTm5Yd1Fkb3dpdHJPbEV1VlROWmdsV2dmaDQwamR3eEtjTzZBZUxkMFBmTm1uN2IKdUtYdHBobzNHdGpkd3ZrQnN3MENRUUNhbVUrdWh6NlNNMWloSTJUaWhaZVRjTGJsVjMzZktKUjRYM1M0UFBvMQorazlMR0oyVW9uREtjSzMyNjYrNDYyWW1mTUFRU0gzdktDdmFNWmk1RUVnUAotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==",
			expectedError: nil,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			_, err := CreateToken(test.ttl, test.payload, test.privateKey)

			assert.Equal(t, err, test.expectedError, fmt.Sprintf("Incorrect result. Expect %v, got %s", test.expectedError, err))
		})
	}
}
