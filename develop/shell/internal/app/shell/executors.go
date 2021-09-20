package shell

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
)

var (
	errEmptyEcho             = errors.New("shell: echo should have some data")
	errNeedProcess           = errors.New("shell: kill command needs name of process")
	errCanNotKillProcess     = errors.New("shell: this process can not be killed")
	errForkNeedAmount        = errors.New("shell: fork command need amount of processes")
	errForkChildNotAvailable = errors.New("shell: fork no child avaliable, exit")
)

// Executor ...
type Executor interface {
	Execute(s *Shell) (string, error)
}

// CDExecutor ...
type CDExecutor struct{}

// Execute ...
func (c *CDExecutor) Execute(s *Shell) (string, error) {
	// сброс до домашнего каталога
	if len(s.Args) == 1 {
		cUser, err := user.Current()
		if err != nil {
			return "", err
		}
		if err := os.Chdir(cUser.HomeDir); err != nil {
			return "", err
		}
	} else {
		// тут как минимум есть указанная директория
		if err := os.Chdir(s.Args[1]); err != nil {
			return "", err
		}

	}

	return "Successfully changed dir", nil
}

// EchoExecutor ...
type EchoExecutor struct{}

// Execute ...
func (e *EchoExecutor) Execute(s *Shell) (string, error) {
	// если аргумент не передан
	if len(s.Args) == 1 {
		return "", errEmptyEcho
	}
	return s.Args[1], nil
}

// PSExecutor ...
type PSExecutor struct{}

// Execute ...
func (p *PSExecutor) Execute(s *Shell) (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	builder.WriteString("PID\t|\tCOMMAND\n")
	builder.WriteString("---------------\n")
	for _, proc := range processes {
		builder.WriteString(
			fmt.Sprintf("%v\t|\t%v\n", proc.Pid(), proc.Executable()),
		)
	}
	builder.WriteString("---------------\n")
	return builder.String(), nil
}

// PWDExecutor ...
type PWDExecutor struct{}

// Execute ...
func (p *PWDExecutor) Execute(s *Shell) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return "Current work dir :" + path, err
}

// KillProcessExecutor ...
type KillProcessExecutor struct{}

// Execute ...
func (k *KillProcessExecutor) Execute(s *Shell) (string, error) {
	// проверим, что есть , что убивать
	if len(s.Args) < 2 {
		return "", errNeedProcess
	}
	processes, err := process.Processes()
	if err != nil {
		return "", err
	}
	for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			return "", err
		}
		if name == s.Args[1] {
			if err := process.Kill(); err != nil {
				return "", errCanNotKillProcess
			}
		}
	}
	return fmt.Sprintf("Process %v successfully killed", s.Args[1]), nil
}

// ForkExecutor ...
type ForkExecutor struct{}

// Execute ...
// убрать чилды
func (f *ForkExecutor) Execute(s *Shell) (string, error) {
	if len(s.Args) < 2 {
		return "", errForkNeedAmount
	}
	fork, err := strconv.Atoi(s.Args[1])
	if err != nil {
		return "", err
	}
	children := []int{}
	var builder strings.Builder
	pid := os.Getpid()
	ppid := os.Getppid()
	builder.WriteString(
		fmt.Sprintf("pid: %d, ppid: %d, forks: %d\n", pid, ppid, fork),
	)
	if _, isChild := os.LookupEnv("CHILD_ID"); !isChild {
		for i := 0; i < fork; i++ {
			args := append(os.Args, fmt.Sprintf("#child_%d_of_%d", i, os.Getpid()))
			childENV := []string{
				fmt.Sprintf("CHILD_ID=%d", i),
			}
			pwd, err := os.Getwd()
			if err != nil {
				return "", err
			}
			childPID, _ := syscall.ForkExec(args[0], args, &syscall.ProcAttr{
				Dir: pwd,
				Env: append(os.Environ(), childENV...),
				Sys: &syscall.SysProcAttr{
					Setsid: true,
				},
				Files: []uintptr{0, 1, 2}, // print message to the same pty
			})
			builder.WriteString(
				fmt.Sprintf("parent %d fork %d\n", pid, childPID),
			)
			if childPID != 0 {
				children = append(children, childPID)
			}
		}
		// print children
		builder.WriteString(
			fmt.Sprintf("parent: PID=%d children=%v", pid, children),
		)
		if len(children) == 0 && fork != 0 {
			return "", errForkChildNotAvailable
		}

		// set env
		for _, childID := range children {
			if c := os.Getenv("CHILDREN"); c != "" {
				os.Setenv("CHILDREN", fmt.Sprintf("%s,%d", c, childID))
			} else {
				os.Setenv("CHILDREN", fmt.Sprintf("%d", childID))
			}
		}
	}
	return "", nil
}
