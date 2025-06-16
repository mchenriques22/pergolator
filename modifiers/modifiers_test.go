package modifiers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mchenriques22/pergolator/tree/defaultparser"
)

func TestFormatKeys(t *testing.T) {
	query := "KEY:value AND Key:value"
	expected := "key:value AND key:value"

	parsed, err := defaultparser.Parse(query)
	require.NoError(t, err)

	formatted := formatKeys(parsed, strings.ToLower)
	assert.Equal(t, expected, formatted.String())
}

func TestFormatValues(t *testing.T) {
	query := "key:VALUE AND -key:VALUE"
	expected := "key:value AND (NOT(key:value))"

	parsed, err := defaultparser.Parse(query)
	require.NoError(t, err)

	formatted := formatValues(parsed, strings.ToLower)
	assert.Equal(t, expected, formatted.String())
}

func TestIgnoreSomeKeys(t *testing.T) {
	testCases := []struct {
		query    string
		expected string
	}{
		{"ignore:value", "Empty struct, should be ignored"},
		{"ignore:value OR other:value", "other:value"},
		{"-ignore:value OR other:value", "other:value"},
		{"ignore:value OR ignore2:value AND other:value", "other:value"},
	}

	for _, testCase := range testCases {
		parsed, err := defaultparser.Parse(testCase.query)
		require.NoError(t, err)

		formatted := ignoreSomeKeys(parsed, []string{"ignore", "ignore2"})
		assert.Equal(t, testCase.expected, formatted.String())
	}
}

func TestFormatKeysToSnakeCase(t *testing.T) {
	query := "AnID:1234 AND myService:my-service"
	expected := "an_id:1234 AND my_service:my-service"

	parsed, err := defaultparser.Parse(query)
	require.NoError(t, err)

	formatted := FormatKeysToSnakeCase(parsed)
	assert.Equal(t, expected, formatted.String())
}

func TestFormatKeysToCamelCase(t *testing.T) {
	query := "an_id:1234 AND my_service:my-service"
	expected := "AnId:1234 AND MyService:my-service"

	parsed, err := defaultparser.Parse(query)
	require.NoError(t, err)

	formatted := FormatKeysToCamelCase(parsed)
	assert.Equal(t, expected, formatted.String())
}
