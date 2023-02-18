package main

import (
"bytes"

"github.com/skip2/go-qrcode"

)

// Create buffer to store image in png format
var buf bytes.Buffer

func main() { // Create qr code
qrCode, err := qrcode.New("http://example.com", qrcode.Medium)

if err != nil {
panic(err)
}

// Write image data to buffer // Encode as png
png.Encode(&buf, qrCode.Image(250))

// Save image data from buffer to file on disk if err = ioutil.WriteFile("qrcode.png", buf.Bytes(), 0644); err != nil { panic(err) } 
}