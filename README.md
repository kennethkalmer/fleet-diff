# fleet-diff

Helps you diff systemd unit files as serialized by [fleet](https://github.com/coreos/fleet).

## Background

Lets assume you've got a directory full of unit files, and you're uncertain
which files in your remote fleet differ from the ones in your local repository.

You could try and diff the output of `fleetctl cat`, but if your unit file has
any newlines in you'll get a false negative.

For example, lets say our unit has something like this:

```
ExecStart=/bin/bash -c ' \
  echo $(etcdctl get /hello)
'
```

This would look more or less like the following when using `fleetctl cat`

```
$ fleetctl cat hello.service
ExecStart=/bin/bash -c '  echo $(etcdctl get /hello) '
```

Doing a simple diff no longer works.

## Usage

`fleet-diff` takes the output from `fleetctl cat` and the path to a local unit
and then parses both using the [unit package](https://github.com/coreos/fleet/blob/master/unit/unit.go) and then compares the units.

```
$ fleetctl cat hello.service | fleet-diff hello.service -
Everything looks fine.
$ echo $?
0
```

`fleet-diff` accepts two arguments, paths to unit files. If one of the paths is a `-` character, `fleet-diff` will read STDIN for the content.

`fleet-diff` will have an exit code of `0` if the units match, or `1` if they don't, making it easier to use in scripting.

### Diffing two files

```
$ fleetctl cat hello.service > hello-submitted.service
$ fleet-diff hello.service hello-submitted.service
Units are different!
$ echo $?
1
```

## Installation

You need [go](https://golang.org/) installed to build the binary, then simply run:

```
$ go install https://bitbucket.org/mylightstone/fleet-diff
```

Go will place the final binary in your `$GOPATH/bin` directory, so make sure it is available in your `$PATH`.

## License


