package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeUpdates(t *testing.T) {
	jsonFile, err := os.Open("updates.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		assert.Fail(t, err.Error())
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	updateContent, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Updates array
	var updates Updates

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(updateContent, &updates)

	assert.Equal(t, 79, updates.Total)
	assert.Equal(t, 4, len(updates.Updates))
}
