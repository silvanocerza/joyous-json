# Joyous JSON

Joyous JSON, aka `jj`, is a tool to process Json in different ways.

# Table of Contents

- [Documentation](#documentation)
  - [Command line tool](#command-line-tool)
    - [`add`](#add)
    - [`prefix`](#prefix)
    - [`filter`](#filter)
    - [Piping](#piping)
  - [Library](#library)
- [Development](#development)

# Documentation

## Command line tool

Joyous JSON can be easily used as a command line tool process your JSON objects. It offers some commands to do that:

### `add`

Adds a new key to all processed JSON objects, if the key already exists its value will be overwritten.

```
$ echo '{"team": "team-c", "id": 1}' | jj add team team-c
{"id":1,"team":"team-c"}

$ echo '{"team": "team-x", "id": 2}' | jj add team team-c
{"id":2,"team":"team-c"}

$ echo '{"id": 3}' | jj add team team-c
{"id":3,"team":"team-c"}
```

### `prefix`

Prefix an existing key with a custom prefix, if the key doesn't exist in a JSON object nothing happens.

If the resulting key after the prefix is applied already exist it will be ovewritten.

```
$ echo '{"team": "team-c"}' | jj prefix team old_
{"old_team":"team-c"}

$ echo '{"id": 4}' | jj prefix team old_
{"id":4}

$ echo '{"old_team": "team-c", "team": "team-x"}' | jj prefix team old_
{"old_team":"team-x"}
```

### `filter`

Filter is used to include or exclude JSON objects based on the value of their keys and wether some keys are set or not. It also supports values comparison.

```
$ echo '{"team": "team-c"}' | jj filter in team
{"team":"team-c"}

$ echo '{"team": "team-c"}' | jj filter out team

$ echo '{"team": "team-c"}' | jj filter in severity

$ echo '{"team": "team-c"}' | jj filter out severity
{"team":"team-c"}

$ echo '{"severity": 2}' | jj filter in severity \>2

$ echo '{"severity": 2}' | jj filter in severity \>=2
{"severity":2}

$ echo '{"severity": 2}' | jj filter out severity \>=2

$ echo '{"severity": 2}' | jj filter out severity 2

$ echo '{"severity": 2}' | jj filter in severity 2
{"severity":2}

$ echo '{"severity": 2}' | jj filter in severity \!2

$ echo '{"severity": 2}' | jj filter in severity \!3
{"severity":2}
```

Remember to skip compators or the command might not run correctly since `bash` parses those symbols for piping and other reasons.

### Piping

Each command prints its output to stdout so it's easy to concatenate multiple operations for more advanced scenario.

The example below will get all the json objects from `log.json`, exclude all objects that don't have `team` set to `team-c`, get all the objects with a `timestamp` value older than `1642328940`, prefix the existing `team` key with `old-` and in the add back the `team` key with `team-q` value.

```
cat log.json \
    | jj filter in team team-c \
    | jj filter out timestamp \<1642328940
    | jj prefix team old- \
    | jj add team team-q
```

`jj` can also be used interactively, if it's executed without piping any value to its stdin it will listen for input and output after the operation is executed. This is a similar behaviour to `cat`.

## Library

Joyous JSON can also be used as a library to process data from and to different streams.

The main actor is the `processor.Processor`, it can be easily created using `processor.New(r io.Reader, w io.Writer)`. In the CLI it's used like so:

```
p := processor.New(io.Stdin, io.Stdout)
```

But it can easily be used to read an HTTP request body and write to stdout:

```
resp, err := http.Get("http://example.com/")
...
p := processor.New(resp.Body, io.Stdout)
```

A `Processor` by itself doesn't do much though, to make it do something we must add steps to it.

`AddStep` expects a function `StepFunc` with this signature `func(m *map[string]interface{}) (bool, error)`. `m` is the JSON object being processed, if `StepFunc` returns `false` or a non nil`error` the JSON object will be discarded.

```
p.AddStep(func(value *map[string]interface{}) (bool, error) {
    (*value)["incident"] = 1234
    return true, nil
})
```

The `step` package offers some builder function to create the most commonly used steps.

```
p.AddStep(step.NewFilterIn("team"))
p.AddStep(step.NewPrefix("team", "old_"))
```

To read a JSON object from the input stream and apply all the added steps simply call `p.Next()`, in case of failures reading from input, writing to output or executing a `StepFunc` it will return an error. If either input or output streams are closed returns an `io.EOF`.

```
err := p.Next()
if err == io.EOF {
    // Handle closure
}
```

`Processor` also offers a `ReadAll()` that will consume the whole input stream and apply all the added steps to each JSON object before returning.

If there is any failure reading from input or writing to output it will return an error. Objects processed before failure will still be written to output stream.

```
err := p.ReadAll()
```

# Development

To build:

```
go build -o jj
```

To run tests:

```
go test ./...
```

To run an example usage scenario, first build the tool with the above command and run:

```
./scenario.sh
```
