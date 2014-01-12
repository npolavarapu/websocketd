
package libwebsocketd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestParsePathWithScriptDir(t *testing.T) {
	baseDir, _ := ioutil.TempDir("", "websockets")
	scriptDir := filepath.Join(baseDir, "foo", "bar")
	scriptPath := filepath.Join(scriptDir, "baz.sh")

	defer os.RemoveAll(baseDir)

	if err := os.MkdirAll(scriptDir, os.ModePerm); err != nil {
		t.Error("could not create ", scriptDir)
	}
	if _, err := os.Create(scriptPath); err != nil {
		t.Error("could not create ", scriptPath)
	}

	config := new(Config)
	config.UsingScriptDir = true
	config.ScriptDir = baseDir

	var res *URLInfo
	var err error

	// simple url
	res, err = parsePath("/foo/bar/baz.sh", config)
	if err != nil {
		t.Error(err)
	}
	if res.ScriptPath != "/foo/bar/baz.sh" {
		t.Error("scriptPath")
	}
	if res.PathInfo != "" {
		t.Error("pathInfo")
	}
	if res.FilePath != scriptPath {
		t.Error("filePath")
	}

	// url with extra path info
	res, err = parsePath("/foo/bar/baz.sh/some/extra/stuff", config)
	if err != nil {
		t.Error(err)
	}
	if res.ScriptPath != "/foo/bar/baz.sh" {
		t.Error("scriptPath")
	}
	if res.PathInfo != "/some/extra/stuff" {
		t.Error("pathInfo")
	}
	if res.FilePath != scriptPath {
		t.Error("filePath")
	}

	// non-existing file
	_, err = parsePath("/foo/bar/bang.sh", config)
	if err == nil {
		t.Error("non-existing file should fail")
	}
	if err != ScriptNotFoundError {
		t.Error("should fail with script not found")
	}

	// non-existing dir
	_, err = parsePath("/hoohar/bang.sh", config)
	if err == nil {
		t.Error("non-existing dir should fail")
	}
	if err != ScriptNotFoundError {
		t.Error("should fail with script not found")
	}
}

func TestParsePathExplicitScript(t *testing.T) {
	config := new(Config)
	config.UsingScriptDir = false

	res, err := parsePath("/some/path", config)
	if err != nil {
		t.Error(err)
	}
	if res.ScriptPath != "/" {
		t.Error("scriptPath")
	}
	if res.PathInfo != "/some/path" {
		t.Error("pathInfo")
	}
	if res.FilePath != "" {
		t.Error("filePath")
	}
}
