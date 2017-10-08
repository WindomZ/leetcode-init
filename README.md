# leetcode-init

> A simple cli to creates leetcode code template.

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
