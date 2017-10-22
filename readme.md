# Simplifying processing barcodes

[![Build Status](https://travis-ci.org/bykovme/gobarcode.svg?branch=master)](https://travis-ci.org/bykovme/gobarcode)  [![codecov](https://codecov.io/gh/bykovme/gobarcode/branch/master/graph/badge.svg)](https://codecov.io/gh/bykovme/gobarcode)

## Included functions

``` go
func CheckEAN13(barcode string) error // Checks if EAN13 code is correct (length, content, checksum)
```