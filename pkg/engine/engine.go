package engine

import (
	"fmt"
	"os"
	"os/exec"
)

type (
	// Runner can run commands on the host
	Runner interface {
		Run(tasks Task) error
	}

	// Options is the options that can be passed to New function
	Options struct {
	}

	// RunOptions is the options to run a ii
	RunOptions struct {
		Path                 string
		Arguments            []string
		EnvironmentVariables map[string]string
	}

	engine struct{}
)

// New - creates new engine
func New(opt *Options) Runner {
	return &engine{}
}

func (e *engine) Run(task Task) error {
	cmd := exec.Command(task.Path, task.Args...)

	envs := []string{}
	for k, v := range task.Envs {
		envs = append(envs, fmt.Sprintf("%s=%s", k, v))
	}
	envs = append(envs, os.Environ()...)
	cmd.Env = envs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to run task %s\n", task.Name)
		return err
	}
	return nil
}
