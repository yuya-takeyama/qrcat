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

### Generate as Data Scheme URI

```
$ echo http://example.com | qrcat -format uri
data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAASgAAAEoAQAAAADDfFG0AAAAJHRFWHRTb2Z0d2FyZQBRUi1QTkcgaHR0cDovL3FyLnN3dGNoLmNvbS9nj329AAAFPElEQVR4AWP4TwwYVTWqalTVqKrBp4rh/////xkYGBgYGBj+MzAw/P//n+E/A8P//wz/GRgYGBgYGP7///9/VBVMFcP/////M/z/////f4b//xn+M/z/z/CfgYGBgeE/w///////Z/j/////UVUwVQz/////z/CfgYHhP8P//wz/Gf4z/P///////wz/Gf4zMDD8Z/j/////UVUwVQz/////z/CfgYHhP8N/BgaG/wwMDP8Z/jMw/P/P8J+BgeE/w/////+PqoKpYvj///9/hv8MDAz/Gf7/Z/j/n4HhPwMDw3+G//8Z/jMwMPxn+P////9RVTBVDP/////P8P//////Gf7/////P8N/BgaG/wwMDP8Z/v////8/w/////+PqoKpYvj///9/BgYGBgYGhv8M/xn+M/xn+M/wn+E/w38GBgYGBgaG/////x9VBVPF8P////////////////////9n+P+f4f////8Z/v//////////////////qCqYKob///////+f4T8DA8N/hv8MDP8Z/jMw/P//n4Hh////DP//M/z/////qCqYKob//////////3+G////M/xnYGD4z8DA8P8/AwPDf4b/////Z/j/////UVUwVQz/////z/Cf4T/DfwYGhv8MDP///2f4//8/AwPDfwYGBgYGhv////8fVQVTxfD//////xkY/v9n+P///38Ghv8MDP//MzD8Z/jP8J+B4T/D//////8fVQVTxfD//////xkYGBgYGBj+/2dg+P+fgYHh/38Ghv8M//////////////+PqoKpYvj/////////MzAw/P/PwMDA8J+BgYGBgeE/AwMDw////xkY/v///39UFUwVw/////8zMPxnYPjPwPCf4T/Df4b/DP/////P8P//////GRgY/v///39UFUwVw/////8zMDAw/Gf4z/D/P8P///8Z/jMwMDD8Z2BgYPj/n+H//////4+qgqli+P///3+G/////2dg+P+fgYGB4T/DfwYGBob//xn+/2f4z/D//////0dVwVQx/P//////////M/z/////f4b//xkYGBj+MzD8Z/j///9/BgaG/////x9VBVPF8P////8M////Z2Bg+P//PwMDA8P/////M/z/z/Cf4f///wwM/////z+qCqaK4f//////MzD8Z/j/n4HhPwMDA8N/hv8M/xn+/2f4z/D//3+G/////x9VBVPF8P////8M/xn+/2dgYGBgYPjPwPD///////8zMDAwMPz//5/h/////0dVwVQx/P///////////////z8DAwPD//////9n+P+f4f///wz//zMw/P////+oKpgqhv////9nYGBgYGBg+P+fgYGB4f9/hv8M/xkY/jP8Z/jPwMDw/////6OqYKoY/v///5/h//////8z/GdgYGBgYPjPwMDA8J/h////DP//M/z//////1FVMFUM/////8/wn4GB4T/DfwaG//8Z/v//z8DA8J+BgYGBgeH/f4b/////H1UFU8Xw/////wz/GRgY/jP8//+fgYHhP8P//wwM//////+fgeH/////////P6oKporh/////xn+MzAw/Gf4z8DAwPCf4T8Dw38GBob///8zMPz/z/D/////o6pgqhj+////n+H//////zP8////PwMDw3+G////M/z/z8Dw//9/hv//////P6oKporh/////xkYGBgYGBj+/2f4////f4b///8zMPz///8/w///DAz/////P6oKporhPzFgVNWoqlFVo6oGnyoAe0A8XC4P1FkAAAAASUVORK5CYII=
```

## License

The MIT License

## Author

Yuya Takeyama
