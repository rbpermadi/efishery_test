package entity_test

import (
	"efishery_test/user_api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		task    entity.User
		isError bool
	}{
		{
			name:    "empty struct",
			task:    entity.User{},
			isError: true,
		},
		{
			name:    "name included non-alphanumeric char",
			task:    entity.User{Phone: "6281321678386", Name: "rojali###"},
			isError: true,
		},
		{
			name:    "role not admin or user",
			task:    entity.User{Phone: "6281321678386", Name: "rojali", Role: "test"},
			isError: true,
		},
		{
			name:    "perfecto",
			task:    entity.User{Name: "rojali", Role: "admin", Phone: "6281321678386"},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isError, tt.task.Validate() != nil)
		})
	}
}
