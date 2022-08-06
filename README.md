# cronlog.go

Go porting of [cronlog](https://github.com/kazuho/kaztools/blob/master/cronlog).
This is my first go project :P

## Usage

### Command

```
$ cronlog cmd args...
```

### Ignore some exit code

Put `/etc/cronlog.toml`

```toml
[[Commands]]
Path        = "/path/to/comannd"
SuccessCode = 10
```

If this configuration exists, when `/path/to/comannd args...` exits with code `10`, the process will be recognized as success.

### Slack Integration

Put `/etc/cronlog.toml`

```toml
[Slack]
Url       = "https://hooks.slack.com/services/xxxxxx/yyyyyy/zzzzzz"
#Channel  = "#yourchannel"
#Username = "cronlog"
```

## Author
* @mallowlabs

