# Spinner

## Install

```shell
$ go get github.com/josa42/go-spinner
```

## Examples

### Simple spinner

![](docs/screenshot-1.gif)

```go
s := spinner.New("Loading")
time.Sleep(1 * time.Second)

s.Done()
```

### Failing spinner

![](docs/screenshot-2.gif)

```go
s := spinner.New("Loading")
time.Sleep(1 * time.Second)

s.Fail()
```


### Updating message

![](docs/screenshot-3.gif)

```go
s := spinner.New("Loading...")
time.Sleep(1 * time.Second)

s.Message("Still loading...")
time.Sleep(1 * time.Second)

s.Message("Almost done")
time.Sleep(1 * time.Second)

s.Message("Done")
s.Done()

```

### Multi step

![](docs/screenshot-4.gif)

```go
s := spinner.New("One")
time.Sleep(1 * time.Second)

s.Next("Two")
time.Sleep(1 * time.Second)

s.Next("Three")
time.Sleep(1 * time.Second)

s.Done()
```

