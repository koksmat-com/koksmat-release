package cmds

import (
	"os"
	"path/filepath"
)

func GetDirectory(pathname string) string {
	dir := pathname
	if (pathname == ".") || (pathname == "") {
		cwd, _ := os.Getwd()
		dir = filepath.Base(cwd)
	}

	return dir
}
