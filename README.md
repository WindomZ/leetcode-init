# leetcode-init

> A tool to creates leetcode code template via cli.

[![Build Status](https://travis-ci.org/WindomZ/leetcode-init.svg?branch=master)](https://travis-ci.org/WindomZ/leetcode-init)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/leetcode-init)](https://goreportcard.com/report/github.com/WindomZ/leetcode-init)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/WindomZ/leetcode-init/leetcode?status.svg)](https://godoc.org/github.com/WindomZ/leetcode-init/leetcode)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FWindomZ%2Fleetcode-init.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FWindomZ%2Fleetcode-init?ref=badge_shield)

## Install
```bash
go get -u github.com/WindomZ/leetcode-init/...
```

## Usage
Take the first problem([1. Two Sum](https://leetcode.com/problems/two-sum/description/)) for example: 

```bash
leetcode-init -t 'Two Sum'
# or
leetcode-init -t two-sum
# or
leetcode-init -u 'https://leetcode.com/problems/two-sum'
```

Then, a directory(name `twosum`) is generated under the current path.

```
.
└── twosum
    ├── README.md
    ├── twosum.go
    └── twosum_test.go
```

If with `-m TEMPLATE.md`:

```bash
leetcode-init -t two-sum -m TEMPLATE.md
```

loads `TEMPLATE.md` file and renders it in the directory.

```
.
└── twosum
    ├── README.md
    ├── TEMPLATE.md
    ├── twosum.go
    └── twosum_test.go
```

## Related

- [WindomZ/leetcode.go](https://github.com/WindomZ/leetcode.go) LeetCode Problem's Solutions(Golang).

## Contributing

Welcome to pull requests, report bugs, suggest ideas and discuss 
**leetcode-init** on [issues page](https://github.com/WindomZ/leetcode-init/issues).

If you like it then you can put a :star: on it.

## Roadmap

- [x] Support Golang.
- [ ] Support JavaScript.
- [ ] Support Python.

## License

[MIT](https://github.com/WindomZ/leetcode-init/blob/master/LICENSE)

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FWindomZ%2Fleetcode-init.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FWindomZ%2Fleetcode-init?ref=badge_large)
