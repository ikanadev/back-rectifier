package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	assert := assert.New(t)
	t.Run("should read existing file", func(t *testing.T) {
		fileData, err := readFile("testfiles/invalid.yaml")
		if assert.Nil(err) {
			wanted := []byte("this file has no sense")
			assert.Equal(wanted, fileData, "file readed content should be equal")
		}
	})
	t.Run("error in bad file", func(t *testing.T) {
		_, err := readFile("bad/path.md")
		assert.NotNil(err, "Should return an error")
	})
}
func TestParseYaml(t *testing.T) {
	t.Run("should parse YAML file correctly", func(t *testing.T) {
		_, err := GetConfig("testfiles/valid.yaml")
		assert.Nil(t, err, "Should get nil error")
	})
}
