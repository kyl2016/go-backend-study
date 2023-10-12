

### create Reader with []byte

```
var buffer []byte
reader := bytes.NewReader(buffer)
```

image decode

```
import (
    "image"
    "image/jpeg"
    "image/png"
)
image.Decode(reader)
```

### write []byte to file

```
err = ioutil.WriteFile("temp_60.jpg", []byte, os.ModePerm)
```

### buf shouldn't be use 

**The new Buffer takes ownership of buf, and the caller should not use buf after this call.**

```
    var buf []byte
	buffer := bytes.NewBuffer(buf)
	err = jpeg.Encode(buffer, newImage, nil)

	return buffer.Bytes(), nil
```

buf 分配空间也不行

### read image file to image.Image

```
src, err := os.Open("test/testdata/flowers_160.jpg")
	if err != nil {
		panic(err)
	}

	srcImg, _, err := image.Decode(src)
	if err != nil {
		panic(err)
	}
```
