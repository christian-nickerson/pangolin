package testing

import (
	"os"
	"path"
	"runtime"
)

// testing init function to run go tests from the
// root folder, not internal pr pkg
func init() {
	_, filename, _, _ := runtime.Caller(0)
	// move up 2 sub-dirs to get to root
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
