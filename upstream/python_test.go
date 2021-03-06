package upstream

import (
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/simon04/aur-out-of-date/pkg"
)

func mockPython() *gock.Response {
	return gock.New("https://pypi.python.org/").
		Get("/pypi/httpie/json").
		Reply(http.StatusOK).
		SetHeader("Content-Type", "application/json").
		BodyString(`{
			"info": {
				"package_url": "http://pypi.python.org/pypi/httpie",
				"download_url": "https://github.com/jkbrzt/httpie",
				"platform": "UNKNOWN",
				"version": "0.9.9",
				"release_url": "http://pypi.python.org/pypi/httpie/0.9.9"
			}
		}`)
}

func TestPythonHttpieSource1(t *testing.T) {
	defer gock.Off()
	mockPython()

	p := pkg.New("httpie", "0", "", "https://pypi.python.org/packages/source/h/httpie/httpie-0.9.9.tar.gz")
	version, err := VersionForPkg(p)
	if err != nil {
		t.Error(err)
	}
	if version != "0.9.9" {
		t.Errorf("Expecting version 0.9.9, but got %v", version)
	}
}

func TestPythonHttpieSource2(t *testing.T) {
	defer gock.Off()
	mockPython()

	p := pkg.New("httpie", "0", "", "https://files.pythonhosted.org/packages/source/h/httpie/httpie-0.9.9.tar.gz")
	version, err := VersionForPkg(p)
	if err != nil {
		t.Error(err)
	}
	if version != "0.9.9" {
		t.Errorf("Expecting version 0.9.9, but got %v", version)
	}
}
