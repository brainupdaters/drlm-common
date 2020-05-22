# DRLM Plugin Spec

The DRLM plugins are written in YAML (and by extension, JSON is also valid). They have the following sections:

## Spec Version

The spec version sets which version of the spec is used. This specification is for the version `v1`

```yaml
specVersion: v1
```

## Metadata

The metadata section is used to get information of the plugin

```yaml
metadata:
  name: <plugin-name> (mysql)
  version: <plugin-version> (v1.0.3) # Uses semver
  description: <plugin-description> (This plugin adds support for MySQL to DRLM)
  license: <license> (AGPL-3.0-only)
  os:
    - <os>
    - <os>
  arch:
    - <arch>
    - <arch>
```

- `name`: the name of the plugin
- `version`: the version of the plugin. It uses SemVer
- `description`: the description of the plugin
- `license`: the license of the plugin
- `os`: list of supported OSes. If this parameter isn't provided, it's assumed that all OSes are supported. Availiable OSes:
  + `linux`
  + `windows`
  + `darwin` (macOS)
  + `aix`
  + `dragonfly`
  + `freebsd`
  + `netbsd`
  + `openbsd`
  + `plan9`
  + `solaris`
- `arch`: list of supported CPU architectures. If this parameter isn't provided, it's assumed that all CPU architectures are supported. Availiable CPU architectures:
  + `amd64`

## Params

The params section are all the parameters that are sent through `drlmctl` when scheduling a task. They can be used inside the commands as variables. They will be prefixed with `DRLM_PLUGIN`, so the parameter `user`, would be accessed as `$DRLM_PLUGIN_USER`. The names will be uppercased automatically

```yaml
params:
  - name: <param-name> (user) # $DRLM_PLUGIN_USER
    required:
      - <action-name> (backup)
      - <action-name> (restore)
  - name: <param-name> (password) # $DRLM_PLUGIN_PASSWORD
```

- `name`: the name of the parameter
- `required`: the actions where the parameter is required. If it's required and it's not provided when scheduling a job, the scheduling will fail

## Actions

The actions section is the section where you define the functionality of the plugin itself and what it's going to be able to do.

```yaml
actions:
  <action-name>: (backup)
    <action-type>: (full)
      - <command-name> (mount)
      - <command-name> (backup-full)
      - <command-name> (umount)

    <action-type>: (partial)
      - <command-name> (mount)
      - <command-name> (backup-partial)
      - <command-name> (umount)

  <action-name>: (restore)
    - <command-name> (mount.check)
    - <command-name> (restore)
    - <command-name> (umount)

```

- `action-name`: is the name of the action
- `action-type`: some actions have different "targets"; you could want to make a full backup, or only the changes since yesterday
- `command-name`: is the name of a command defined in the commands section. In case you want to access to a command inside a group, you can do so by using a dot notation: `command-group.subcommand`

## Commands

The commands section is where all the diferent commands of the plugin are specified. They can be combined in groups. All the command names and command group names have to be unique.

```yaml
commands:
  <command-name>: (mount)
    workdir: <workdir> (/home/drlm)
    command:
      - <command> (mount)
      - <parameter> (/dev/sda1)
      - <parameter> (/mnt)

  <command-group>:
    <command-group>:
      <command-name>:
        command:
          - <command>
          - <parameter>
          - <parameter>

      <command-name>:
        command:
          - <command>
          - <parameter>
          - <parameter>

    <command-name>:
      workdir: <workdir>
      command:
        - <command>
        - <parameter>
        - <parameter>
```

- `command-name`: the command name is the name that is going to be used in the actions section
- `workdir`: the directory from where the command is going to be run. This is optional. If no workdir is provided, the home directory of the user that is running the DRLM Agent is going to be used
- `command`: the binary that is going to be used
- `parameter` the parameters passed to the command
- `command-group`: the commands can be grouped in command groups. When using a command group in an action, all the methods inside are going to be executed sequentially one after the other. Also, command groups can have subgroups inside

## Tips and tricks

### Importing commands

Let's say you have commands (for example, mount and umount) that get executed with each command. You can use the default YAML syntax to do so:

```yaml
commands:
  mount: &mount
    workdir: /home/drlm
    command:
      - mount
      - /dev/sda1
      - ./mount

  umount: &umount
    workdir: /home/drlm
    command:
      - umount
      - ./mount

  backup:
    mount: *mount
    backup:
      workdir: /home/drlm
      command:
        - ./bak.sh
        - /var/lib/program
        - ./mount
    umount: *umount

  restore:
    mount: *mount
    restore:
      workdir: /home/drlm
      command:
        - ./restore.sh
        - ./mount
        - /var/lib/program
    umount: *umount
```

The key here is to set a reference to a command with the `&refName` and then use it with `*refName`. If the reference is within a command group, the recommendation is using lowerCamelCase: `&commandgroupCommandsubgroupCommandname`

## Full example

```yaml
specVersion: v1
metadata:
  name: tar
  version: 1.0.3
  description: This plugin adds support for the tar command to DRLM
  license: AGPL-3.0-only
  os: [ "linux", "darwin" ]

params:
  - name: source
    required:
      - backup
  - name: destination
    required:
      - restore
  - name: prefix

actions:
  backup:
    compressed:
      - backup-compressed
      # Example of how you can access a command within a command group
      - group-example.command-example
    uncompressed:
      - backup-uncompressed
  restore:
    compressed:
      - restore-compressed
    uncompressed:
      - restore-uncompressed

commands:
  backup-compressed:
    command:
      - tar
      - czvf
      - "$DRLM_STORAGE/drlm.tar.gz"
      - $DRLM_PLUGIN_SOURCE

  backup-uncompressed:
    command:
      - tar
      - cvf
      - "$DRLM_STORAGE/drlm.tar"
      - $DRLM_PLUGIN_SOURCE

  restore-compressed:
    restore:
      command:
        - tar
        - xzvf
        - "$DRLM_STORAGE/drlm.tar.gz"
        - -C
        - $DRLM_PLUGIN_DESTINATION
    check: &restore-compressedCheck
      workdir: /
      command:
        - test
        - -f
        - $DRLM_PLUGIN_DESTINATION

  restore-uncompressed:
    restore:
      command:
        - tar
        - xf
        - "$DRLM_STORAGE/drlm.tar"
        - -C
        - $DRLM_PLUGIN_DESTINATION
    check: *restoreCompressedCheck

  group-example:
    command-example:
      command:
        - echo
        - \"It works! :D\"
```