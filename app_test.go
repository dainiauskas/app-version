package version

import (
	"testing"
)

func Test_Application(t *testing.T) {
	_, err := Init("Test", "", "wrong version", "")
	if err == nil {
		t.Errorf("Wanted error, but got nil")
	}

	_, err = Init("", "", "v1.0.0", "6bbab8df7")
	if err == nil {
		t.Errorf("Wanted error, but got nil")
	}

	app, err := Init("Test", "Testing Application", "v1.0.0", "6bbab8df7")
	if err != nil {
		t.Errorf("Wanted error nil, but got: %s", err)
	}

	t.Log(app.String())

	_, err = app.JSON()
	if err != nil {
		t.Errorf("Got error: %s", err)
	}
}
