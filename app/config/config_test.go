package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	os.Setenv("MODE", "test")
	os.Setenv("MONGO_USER", "test")
	os.Setenv("MONGO_PASSWORD", "test")
	os.Setenv("MONGO_DOMAIN", "test.cluster")

	c := GetConfig()

	if c.Mongo.Uri != "mongodb+srv://test:test@test.cluster/?retryWrites=true&w=majority" {
		t.Errorf("mongo.uri is invalid. %v\n", c.Mongo.Uri)
	}
}
