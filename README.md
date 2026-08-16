# url-query

Parse or encode application/x-www-form-urlencoded query strings (CLI).

## Usage

```
url-query -mode parse -in query.txt
url-query -mode encode -in query.txt -out -
```

- `-mode`  `parse` (default) or `encode`
- `-in`    input path, `-` means stdin
- `-out`   output path, `-` means stdout

`+` is a space. Percent escapes must be two hex digits.
A leading `?` is stripped.
