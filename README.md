# qrcat

Generate QR codes in UNIX way

## Installation

```
$ go get github.com/yuya-takeyama/qrcat
```

## Usage

### Generate as ASCII

```
$ echo http://example.com | qrcat
```

or

```
$ qrcat url.txt
```

### Generate as PNG

```
$ echo http://example.com | qrcat -format png
```

or

```
$ qrcat -format png url.txt
```

## License

The MIT License

## Author

Yuya Takeyama
