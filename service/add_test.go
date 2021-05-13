package service_test

import (
	"testing"

	"github.com/phati/circleci-demo/service"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := service.Add(1, 2)
	assert.Equal(t, 3, result)
}
