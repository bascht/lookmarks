# Lookmarks

Quick helper utility initially prototyped with
[Lucas](https://github.com/moonglum/) to generate a lookup list of commonly used
URLs from a standard Yaml file.

Think of it as a kind-of "bookmark builder" which enables you to generate bookmark list from a hierarchically defined Yaml structure:

```yaml
bookmarks:
  - url: https://example.com
    key: example
    sub:
      - url: foo
      - url: bar
      - url: baz
        key: alias
```

Will generate the following lookup table:

| key           | value                     |
|---------------|---------------------------|
| example       | https://example.com       |
| example-foo   | https://example.com/foo   |
| example-bar   | https://example.com/bar   |
| example-alias | https://example.com/alias |


