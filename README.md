# Glass

Run, report back, and keep a full history of the command execution life cycle.

Easily run and manage commands with Glass, Glass can run any command, and it will write the stdout & the stderr to easily human/machine-readable/interpreted log file(unique YAML log file per command) along with everything about what happened during the running process, you can give the command an ID and follow its progress later and check to see whether it's completed or no and its exit status, log, and so on.

## Commands

```shell
Easily run and manage commands with Glass, 
        Glass can run any command, and it will write the 
        stdout & the stderr to log file along with everything, like 
        you can give the command an ID and follow its progress later 
        and check to see whether it's completed or not and its exit status 
        and so on.

Usage:
  glass [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  log         Return the command log
  purge       Delete all logs or single log
  report      Return the status of command
  run         Run command
  status      Return exit code of the command
  version     Return Glass's version

Flags:
  -h, --help   help for glass

Use "glass [command] --help" for more information about a command.
```

## Examples

### Run a command

```shell
glass run -u a1 -- ls -l
```

### Get the command log

```shell
glass log -u a1
```

### Get the exit code of the command

```shell
glass status -u a1
```

### Get the status of the command

```shell
glass report -u a1
```

## Logs Location

You can pass a writable absolute folder path using `GLASS_LOG_PATH` environment variable, otherwise, all the logs will be in a folder called `glass` beside the binary.

## Notes

- Glass ignores the hup signal to remove the interruption that may be caused by it.
