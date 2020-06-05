package plugin

import (
	"fmt"

	"github.com/brainupdaters/drlm-common/pkg/os"

	"github.com/blang/semver"
	"github.com/goccy/go-yaml"
)

type pluginYAML struct {
	SpecVersion string                            `yaml:"specVersion,omitempty" json:"spec_version,omitempty"`
	Metadata    *metadataYAML                     `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Params      []*paramYAML                      `yaml:"params,omitempty" json:"params,omitempty"`
	Actions     map[string]interface{}            `yaml:"actions,omitempty"`
	Commands    map[string]map[string]interface{} `yaml:"commands,omitempty"`
}

type metadataYAML struct {
	Name        string   `yaml:"name,omitempty"`
	Version     string   `yaml:"version,omitempty"`
	Description string   `yaml:"description,omitempty"`
	License     string   `yaml:"license,omitempty"`
	OS          []string `yaml:"os,omitempty"`
	Arch        []string `yaml:"arch,omitempty"`
}

type paramYAML struct {
	Name     string      `yaml:"name,omitempty"`
	Required interface{} `yaml:"required,omitempty"`
}

func Unmarshal(b []byte) (*Plugin, error) {
	pY := &pluginYAML{}
	p := &Plugin{}

	if err := yaml.Unmarshal(b, pY); err != nil {
		return nil, err
	}

	if err := unmarshalSpecVersion(pY, p); err != nil {
		return nil, fmt.Errorf("unmarshal specVersion: %w", err)
	}

	if err := unmarhsalMetadata(pY, p); err != nil {
		return nil, fmt.Errorf("unmarshal metadata: %w", err)
	}

	if err := unmarshalParams(pY, p); err != nil {
		return nil, fmt.Errorf("unmarshal params: %w", err)
	}

	if err := unmarhsalActions(pY, p); err != nil {
		return nil, fmt.Errorf("unmarshal actions: %w", err)
	}

	if err := unmarshalCommands(pY, p); err != nil {
		return nil, fmt.Errorf("unmarshal commands: %w", err)
	}

	if err := validatePlugin(p, p.Actions); err != nil {
		return nil, err
	}

	return p, nil
}

func validatePlugin(p *Plugin, actions []*Action) error {
	for _, a := range actions {
		if len(a.Commands) != 0 {
			for _, c := range a.Commands {
				if _, err := p.GetCommand(c); err != nil {
					return fmt.Errorf("invalid plugin spec: action '%s': command '%s': %w", a.Name, c, err)
				}
			}
		} else {
			if err := validatePlugin(p, a.Types); err != nil {
				return err
			}
		}
	}

	return nil
}

func unmarshalSpecVersion(src *pluginYAML, dst *Plugin) error {
	v, err := semver.ParseTolerant(src.SpecVersion)
	if err != nil {
		return fmt.Errorf("parse version: %w", err)
	}
	dst.SpecVersion = v

	return nil
}

func unmarhsalMetadata(src *pluginYAML, dst *Plugin) error {
	dst.Metadata = &Metadata{
		Name:        src.Metadata.Name,
		Description: src.Metadata.Description,
		License:     src.Metadata.License,
	}

	v, err := semver.ParseTolerant(src.Metadata.Version)
	if err != nil {
		return fmt.Errorf("parse version: %w", err)
	}
	dst.Metadata.Version = v

	for _, strOS := range src.Metadata.OS {
		parsedOS := os.FromString(strOS)
		if parsedOS == os.Unknown {
			return fmt.Errorf("parse OS: unknown OS: %s", strOS)
		}

		dst.Metadata.OS = append(dst.Metadata.OS, parsedOS)
	}

	for _, strArch := range src.Metadata.Arch {
		parsedArch := os.ArchFromString(strArch)
		if parsedArch == os.ArchUnknown {
			return fmt.Errorf("parse arch: unknown arch: %s", strArch)
		}

		dst.Metadata.Arch = append(dst.Metadata.Arch, parsedArch)
	}

	return nil
}

func unmarshalParams(src *pluginYAML, dst *Plugin) error {
	dst.Params = []*Param{}

	for _, param := range src.Params {
		p := &Param{
			Name: param.Name,
		}

		if _, ok := param.Required.(bool); ok {
			p.RequiredAll = true
		}

		if v, ok := param.Required.([]interface{}); ok {
			for _, i := range v {
				if s, ok := i.(string); ok {
					p.RequiredActions = append(p.RequiredActions, s)
				}
			}
		}

		dst.Params = append(dst.Params, p)
	}

	return nil
}

func unmarhsalActions(src *pluginYAML, dst *Plugin) error {
	dst.Actions = []*Action{}

	for name, action := range src.Actions {
		dst.Actions = append(dst.Actions, unmarshalAction(name, action))
	}

	return nil
}

func unmarshalAction(name string, action interface{}) *Action {
	commands := []string{}
	types := []*Action{}

	if v, ok := action.([]interface{}); ok {
		for _, i := range v {
			if s, ok := i.(string); ok {
				commands = append(commands, s)
			}
		}

	} else {
		commands = nil

		if v, ok := action.(map[string]interface{}); ok {
			for typeName, actionType := range v {
				types = append(types, unmarshalAction(typeName, actionType))
			}

		} else {
			types = nil
		}
	}

	return &Action{
		Name:     name,
		Commands: commands,
		Types:    types,
	}
}

func unmarshalCommands(src *pluginYAML, dst *Plugin) error {
	dst.Commands = []*Command{}

	for name, command := range src.Commands {
		cmd, err := unmarshalCommand(name, command)
		if err != nil {
			return err
		}

		dst.Commands = append(dst.Commands, cmd)
	}

	return nil
}

func unmarshalCommand(name string, command map[string]interface{}) (*Command, error) {
	isCommand := false
	isGroup := false
	for k := range command {
		switch k {
		case "command", "workdir":
			isCommand = true

		default:
			isGroup = true
		}
	}

	if isGroup && isCommand {
		return nil, fmt.Errorf("command '%s': invalid syntax: has commands and is also a command group. The keys 'command' and 'workdir' are reserved", name)
	}

	if isCommand {
		cmd := []string{}
		if v, ok := command["command"].([]interface{}); ok {
			for _, i := range v {
				if s, ok := i.(string); ok {
					cmd = append(cmd, s)
				}
			}
		} else {
			return nil, fmt.Errorf("command '%s': invalid syntax: the 'command' key should be an array, but it's %T", name, command["command"])
		}

		workdir := ""
		if _, ok := command["workdir"]; ok {
			if v, ok := command["workdir"].(string); ok {
				workdir = v
			} else {
				return nil, fmt.Errorf("command '%s': invalid syntax: the 'workdir' key should be a string, but it's %T", name, command["workdir"])
			}
		}

		return &Command{
			Name:    name,
			Command: cmd,
			Workdir: workdir,
		}, nil
	}

	if isGroup {
		cmds := []*Command{}
		for k, v := range command {
			val, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("command '%s': invalid syntax: subcommand '%s' has to be a map, but it's %T", name, k, v)
			}

			cmd, err := unmarshalCommand(k, val)
			if err != nil {
				return nil, fmt.Errorf("command '%s': %v", name, err)
			}

			cmds = append(cmds, cmd)
		}

		return &Command{
			Name:        name,
			SubCommands: cmds,
		}, nil
	}

	return nil, fmt.Errorf("command '%s': invalid syntax: the command is neither a command or a command group", name)
}
