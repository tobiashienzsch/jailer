package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
)

func TestNewFromFile(t *testing.T) {

	var tests = []struct {
		input       string
		expected    *jailer.Jailerfile
		expectError bool
	}{
		{"testdata/noexist/Jailerfile", &jailer.Jailerfile{}, true},
		{"testdata/label/Jailerfile", &jailer.Jailerfile{BaseImage: jailer.BaseImage{Name: "freebsd", Version: "latest"}}, false},
	}

	for _, tt := range tests {
		actual, err := jailer.NewFromFile(tt.input)

		// error
		if tt.expectError != (err != nil) {
			t.Errorf("Error %s", err)
		}

		// from
		if actual.BaseImage != tt.expected.BaseImage {
			t.Errorf("Expected: %q, Got: %q", tt.expected.BaseImage, actual.BaseImage)
		}
	}
}

func TestLabelParsing(t *testing.T) {

	jf, err := jailer.NewFromFile("testdata/label/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.Labels["maintainer"] != `"example@example.com"` {
		t.Errorf("Expected: \"%s\", got %s", "example@example.com", jf.Labels["maintainer"])
	}

	if jf.Labels["version"] != `"1.0"` {
		t.Errorf("Expected: \"%s\", got %s", "1.0", jf.Labels["version"])
	}

}

func TestFromWithImplicitLatestParsing(t *testing.T) {

	jf, err := jailer.NewFromFile("testdata/from/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.BaseImage.Name != "freebsd" {
		t.Errorf("Expected: %s, got %s", "freebsd", jf.BaseImage.Name)
	}

	if jf.BaseImage.Version != "latest" {
		t.Errorf("Expected: %s, got %s", "latest", jf.BaseImage.Version)
	}

}

func TestFromWithExplicitLatestParsing(t *testing.T) {

	jf, err := jailer.NewFromFile("testdata/from_with_latest/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.BaseImage.Name != "freebsd" {
		t.Errorf("Expected: %s, got %s", "freebsd", jf.BaseImage.Name)
	}

	if jf.BaseImage.Version != "latest" {
		t.Errorf("Expected: %s, got %s", "latest", jf.BaseImage.Version)
	}

}

func TestFromWithExplicitVersionParsing(t *testing.T) {

	jf, err := jailer.NewFromFile("testdata/from_with_version/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.BaseImage.Name != "freebsd" {
		t.Errorf("Expected: %s, got %s", "freebsd", jf.BaseImage.Name)
	}

	if jf.BaseImage.Version != "12.1" {
		t.Errorf("Expected: %s, got %s", "12.1", jf.BaseImage.Version)
	}

}

func TestRunParsing(t *testing.T) {

	jf, err := jailer.NewFromFile("testdata/run/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if len(jf.Instructions) != 2 {
		t.Errorf("Expected: %d, got %d", 2, len(jf.Instructions))
	}

	t.Run("first", func(t *testing.T) {
		if jf.Instructions[0].Name() != "RUN" {
			t.Errorf("Expected: %s, got %s", "RUN", jf.Instructions[0].Name())
		}

		val := jf.Instructions[0].(*jailer.RunInstruction)
		expected := "echo \"Hello Jailer!\""
		if val.Command != expected {
			t.Errorf("Expected: %s, got %s", expected, val.Command)
		}
	})

	t.Run("second", func(t *testing.T) {
		if jf.Instructions[1].Name() != "RUN" {
			t.Errorf("Expected: %s, got %s", "RUN", jf.Instructions[1].Name())
		}

		val := jf.Instructions[1].(*jailer.RunInstruction)
		expected := "pkg install -y nano"
		if val.Command != expected {
			t.Errorf("Expected: %s, got %s", expected, val.Command)
		}
	})

}

func TestWorkDirParsing(t *testing.T) {

	jf, err := jailer.NewFromFile("testdata/workdir/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if len(jf.Instructions) != 1 {
		t.Errorf("Expected: %d, got %d", 2, len(jf.Instructions))
	}

	if jf.Instructions[0].Name() != "WORKDIR" {
		t.Errorf("Expected: %s, got %s", "WORKDIR", jf.Instructions[0].Name())
	}

	val := jf.Instructions[0].(*jailer.WorkDirInstruction)
	expected := "/work"
	if val.Command != expected {
		t.Errorf("Expected: %s, got %s", expected, val.Command)
	}

}