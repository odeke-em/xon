package pkger

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var prettyGitHashGetArgs = []string{"log", "-n", "1", "--pretty=format:'%H'"}

type PkgInfo struct {
	CommitHash string
	GoVersion  string
	OsVersion  string
}

func (p *PkgInfo) String() string {
	return fmt.Sprintf("Commit Hash: %s\nOS Version: %s\nGo Version: %s",
		p.CommitHash, p.OsVersion, p.GoVersion)
}

func Recon(pkgPath string) (pkgInfo *PkgInfo, err error) {
	var gitProgPath, pkgPathAbs string

	// Firstly look up if they've got Git
	gitProgPath, err = exec.LookPath("git")
	if err != nil {
		return
	}

	if pkgPathAbs, err = filepath.Abs(pkgPath); err != nil {
		return
	}

	args := []string{gitProgPath}
	args = append(args, prettyGitHashGetArgs...)

	cmd := exec.Cmd{
		Args: args,
		Dir:  pkgPathAbs,
		Path: gitProgPath,
	}

	var output []byte
	if output, err = cmd.Output(); err != nil {
		return
	}

	pkgInfo = &PkgInfo{
		CommitHash: string(output),
		OsVersion:  strings.Join([]string{runtime.GOOS, runtime.GOARCH}, " "),
		GoVersion:  runtime.Version(),
	}

	return
}
