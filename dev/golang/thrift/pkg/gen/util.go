package gen

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"text/template"
)

func execModTidy(dirPath string) error {
	return execCommand(dirPath, "go", nil, "mod", "tidy")
}

func GenerateMockAPI(mockServiceName, fileName, thriftIDL string) error {
	sp, err := GenerateMockService(mockServiceName, thriftIDL)
	if err != nil {
		return err
	}

	tpl, err := template.New("tplServiceMock").Parse(tplServiceMock)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, sp)
	if err != nil {
		return err
	}
	os.WriteFile(fileName, b.Bytes(), 0644)

	return nil
}

func CompileThriftToGo(stagingDir, thriftIdl string, env []string) error {
	return execCommand(stagingDir, "thrift", env, "-r", "-gen", "go", "-out", stagingDir, thriftIdl)
}

func CmdExec(dir, name string, envs []string, args ...string) error {
	cmd := exec.Command(name, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	if len(envs) != 0 {
		cmd.Env = envs
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case <-c:
				cmd.Process.Release()
				cmd.Process.Kill()
				break LOOP
			}
		}
	}(ctx)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func execCommand(dir, cmd string, envs []string, args ...string) error {
	c := exec.Command(cmd, args...)

	if dir != "" {
		c.Dir = dir
	}

	if len(envs) != 0 {
		c.Env = envs
	}
	stdout, err := c.Output()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			return err
		}
		return fmt.Errorf("%w: %s", err, string(exitErr.Stderr))
	}
	if len(stdout) != 0 {
		fmt.Println("Exec output: ", stdout)
	}
	return nil
}
