# Formatter

It does what it says, lint and fix any file based on a configuration.

## Examples

With a simple config:

``` yaml
formatter:
  linters:
    - name: go
      command: gofmt -d "$FILENAME"
  fixers:
    - name: go
      command: gofmt -l -w "$FILENAME"
  matchers:
    - path_regex: .*\.go
      fixer_name: go
      linter_name: go
```

``` bash
FORMATTER_CONFIG=config.yaml
```

The lint command shows errors:

``` bash
formatter lint file.go
```

The fix command updates the file in place:

``` bash
formatter fix file.go
```

NOTE: lint and fix are just conventions, you can write linters and
fixers that do anything bash will let you.

## Config

``` yaml
formatters:
  linters:
    - name: linter-name
      command: echo any bash command
  fixers:
    - name: big-fixer
      command: |
        echo multi-line
        echo command
  matchers:
    - path_regex: optional regex.*fullmatch (filename|filepath)
    - shebang_regex: regex?matching the (#! .*)
      linter: linter-name
      fixer: big-fixer
```
