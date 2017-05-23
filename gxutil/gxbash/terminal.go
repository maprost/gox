package gxbash

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/maprost/gox/gxutil/gxlog"
	"os/exec"
)

func Command(logLevel gxlog.Level, cmdName string, cmdArgs ...string) (string, error) {
	gxlog.Print(logLevel, append([]string{cmdName}, cmdArgs...))
	cmd := exec.Command(cmdName, cmdArgs...)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	return evalOutput(out.String(), stderr.String(), err)
}

func Stream(logLevel gxlog.Level, cmdName string, cmdArgs ...string) (string, error) {
	gxlog.Print(logLevel, append([]string{cmdName}, cmdArgs...))
	cmd := exec.Command(cmdName, cmdArgs...)

	var out string
	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stdoutScanner := bufio.NewScanner(stdoutReader)
	go func() {
		for stdoutScanner.Scan() {
			txt := stdoutScanner.Text()
			out += txt
			gxlog.Info(txt)
		}
	}()

	var stderr string
	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	stderrScanner := bufio.NewScanner(stderrReader)
	go func() {
		for stderrScanner.Scan() {
			txt := stderrScanner.Text()
			stderr += txt
			gxlog.Warn(txt)
		}
	}()

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	return evalOutput(out, stderr, err)
}

func evalOutput(out string, stderr string, err error) (string, error) {
	// there is an error
	if err != nil {
		// maybe only the return is empty -> no error
		if len(stderr) == 0 && len(out) == 0 {
			return "", nil
		}

		// some more details in std error?
		if len(stderr) > 0 {
			return "", errors.New(err.Error() + ":" + stderr)
		}

		// something else -> error
		return "", err
	}

	// return output
	return out, nil
}
