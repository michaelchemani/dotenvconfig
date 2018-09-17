package dotenvconfig_test

import (
	"os"
	"testing"

	"github.com/lightspeedretail/dotenvconfig"
)

func TestOverrideDefaultValue(t *testing.T) {

	dotenvconfig.Load("testing.env")

	if os.Getenv("TEST_VALUE") != "my-awesome-test-value" {
		t.Errorf("Env variables not loaded")
	}

}

func TestCommentInDotEnv(t *testing.T) {
	dotenvconfig.Load("testing.env")

	if os.Getenv("COMMENT_VALUE") != "" {
		t.Errorf("Failure when reading commented line")
	}
}

func TestReadLineWithSpecialChars(t *testing.T) {
	dotenvconfig.Load("testing.env")

	if os.Getenv("TEST_SPECIAL_CHARS") != "ab=====cde12345!@#$%^++''{}=====" {
		t.Errorf("Failure when reading a line with special chars")
	}
}
