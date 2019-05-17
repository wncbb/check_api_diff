package conf

import (
	"path"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log(path.Split("static/myfile.css"))
	t.Log(path.Split("myfile.css"))
	t.Log(path.Split(""))
}
