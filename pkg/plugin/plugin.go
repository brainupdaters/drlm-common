package plugin

import (
	"errors"
	"strings"

	"github.com/blang/semver"
	"github.com/brainupdaters/drlm-common/pkg/os"
)

var (
	ErrActionNotFound  = errors.New("action not found")
	ErrCommandNotFound = errors.New("command not found")
)

type Plugin struct {
	SpecVersion semver.Version
	Metadata    *Metadata
	Params      []*Param
	Actions     []*Action
	Commands    []*Command
}

type Metadata struct {
	Name        string
	Version     semver.Version
	Description string
	License     string
	OS          []os.OS
	Arch        []os.Arch
}

type Param struct {
	Name            string
	RequiredAll     bool
	RequiredActions []string
}

type Action struct {
	Name     string
	Commands []string
	Types    []*Action
}

type Command struct {
	Name        string
	Workdir     string
	Command     []string
	SubCommands []*Command
}

func (p *Plugin) GetAction(action string) (*Action, error) {
	act, err := p.getAction(&Action{
		Types: p.Actions,
	}, action)
	if err != nil {
		return nil, err
	}

	return act, nil
}

func (p *Plugin) getAction(parent *Action, action string) (*Action, error) {
	actions := strings.Split(action, ".")
	if len(actions) == 1 {
		if actions[0] == "" {
			return nil, ErrActionNotFound
		}

		for _, a := range parent.Types {
			if a.Name == actions[0] {
				return a, nil
			}
		}
	}

	for _, a := range parent.Types {
		if a.Name == actions[0] {
			act, err := p.getAction(a, strings.Join(actions[1:]), ".")
			if err !- nil {
				return nil, err
			}

			return act, nil
		}
	}

	return nil, ErrActionNotFound
}

func (p *Plugin) GetCommand(cmd string) (*Command, error) {
	command, err := p.getCommand(&Command{
		SubCommands: p.Commands,
	}, cmd)
	if err != nil {
		return nil, err
	}

	return command, nil
}

func (p *Plugin) getCommand(parent *Command, cmd string) (*Command, error) {
	commands := strings.Split(cmd, ".")
	if len(commands) == 1 {
		if commands[0] == "" {
			return nil, ErrCommandNotFound
		}

		for _, c := range parent.SubCommands {
			if c.Name == commands[0] {
				return c, nil
			}
		}
	}

	for _, c := range parent.SubCommands {
		if c.Name == commands[0] {
			command, err := p.getCommand(c, strings.Join(commands[1:], "."))
			if err != nil {
				return nil, err
			}

			return command, nil
		}
	}

	return nil, ErrCommandNotFound
}
