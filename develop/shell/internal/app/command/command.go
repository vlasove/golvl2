package command

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/vlasove/golvl2/shell/internal/app/shell"
)

var (
	// ShellTerminalPrefix ...
	ShellTerminalPrefix = `vlasove@shell:~`
	// ShellExitCommand ...
	ShellExitCommand   = `\exit`
	errConsoleInput    = errors.New("shell: can not read data")
	errPrefixBuildFail = errors.New("shell: can not check current directory place")
	errBadUser         = errors.New("shell: can not get current user info")
	successExitMessage = "shell: success exit. Bye."
)

// ShellTerminal ...
type ShellTerminal struct {
	shell  *shell.Shell
	reader io.Reader
	writer io.Writer
}

// New ...
func New(shell *shell.Shell, reader io.Reader, writer io.Writer) *ShellTerminal {
	return &ShellTerminal{
		shell:  shell,
		reader: reader,
		writer: writer,
	}
}

// Start ...
func (s *ShellTerminal) Start() error {
	fmt.Fprintln(s.writer, `ПРОСТЕЙШИЙ SHELL. ДЛЯ ВЫХОДА ИСПОЛЬЗУЙТЕ \exit`)
	scanner := bufio.NewScanner(s.reader)
	for {
		prefix, err := s.buildPrefix()
		if err != nil {
			return errPrefixBuildFail
		}
		fmt.Fprint(s.writer, prefix)
		scanner.Scan()
		text := scanner.Text()
		if text == ShellExitCommand {
			break
		}
		args := strings.Fields(text)
		s.shell.SetArgs(args)
		switch args[0] {
		case "cd":
			cdExecutor := &shell.CDExecutor{}
			s.shell.SetExecutor(cdExecutor)
		case "echo":
			echoExecutor := &shell.EchoExecutor{}
			s.shell.SetExecutor(echoExecutor)
		case "ps":
			psExecutor := &shell.PSExecutor{}
			s.shell.SetExecutor(psExecutor)
		case "pwd":
			pwdExecutor := &shell.PWDExecutor{}
			s.shell.SetExecutor(pwdExecutor)
		case "kill":
			killExecutor := &shell.KillProcessExecutor{}
			s.shell.SetExecutor(killExecutor)
		case "fork":
			forkExecutor := &shell.ForkExecutor{}
			s.shell.SetExecutor(forkExecutor)
		default:
			fmt.Fprintln(s.writer, "shell: unkown command")
		}
		res, err := s.shell.Start()
		if err != nil {
			fmt.Fprintln(s.writer, err.Error())
			continue
		}
		fmt.Fprintln(s.writer, res)
	}
	if scanner.Err() != nil {
		return errConsoleInput
	}
	if _, err := fmt.Fprintln(s.writer, successExitMessage); err != nil {
		return err
	}
	return nil
}

func (s *ShellTerminal) buildPrefix() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	var postfix string
	userName, err := user.Current()
	if err != nil {
		return "", errBadUser
	}
	if path == "/home/"+userName.Name {
		postfix = "$ "
	} else {
		postfix = " " + path + " "
	}
	return ShellTerminalPrefix + postfix, nil
}
