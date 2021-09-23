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

## License
The MIT License

Copyright (c) 2014 mallowlabs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

