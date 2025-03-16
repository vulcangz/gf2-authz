package compile

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
)

func TestNewCompiler(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	clock := ctime.NewMockClock(ctrl)
	// When
	compilerInstance := NewCompiler(
		clock,
	)

	// Then
	assert := assert.New(t)

	assert.IsType(new(compiler), compilerInstance)

	assert.Equal(clock, compilerInstance.clock)
}
