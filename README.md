# go-json-config
Primitive JSON config file lib for Golang

## Install

`go get github.com/nikserg/go-json-config`

## Usage

1. Create config file in program dir, for example `config.json` with program config values. HINT! Use only first-letter-uppercased names for config field. It will save you typing later.
2. Create `struct` implementing config structure. For example, our `config.json` is:
```
{
    "Token": "abc",
    "Id": 456
}
```

that means our structure must be:

```
type Config struct {
	Token string
	Id int
}
```

if we want to use lowercased field names, like:
```
{
    "token": "abc",
    "id": 456
}
```

struct must be:
```
type Config struct {
	Token string `json:"token"`
	Id int `json:"id"`
}
```
3. Load your file:
```
var config Config
err := ReadConfig("config.json", &config)
if err != nil {
    panic(err)
}
```

4. Awesome! Config values are in `config` struct.

```
fmt.println(config.Token)
```