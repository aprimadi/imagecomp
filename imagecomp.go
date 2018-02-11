package main

import (
  "bytes"
  "fmt"
  "image/png"
  "image/jpeg"
  "io"
  "io/ioutil"
  "os"
  "path/filepath"
  "regexp"
  "strings"

  pngquant "github.com/yusukebe/go-pngquant"
  "github.com/nickalie/go-mozjpegbin"
)

func optimizePNG(path string) bool {
  // Read image
  finput, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  input, err := ioutil.ReadAll(finput)
  if err != nil {
    panic(err)
  }
  in := bytes.NewReader(input)
  img, err := png.Decode(in)
  finput.Close()
  if err != nil {
    return false
  }

  // Encode image
  out := new(bytes.Buffer)
  cimg, err := pngquant.Compress(img, "1")
  if err != nil {
    panic(err)
  }
  err = png.Encode(out, cimg)
  if err != nil {
    panic(err)
  }

  outlen := int64(out.Len())
  if outlen < in.Size() {
    // Write to file
    f, err := os.Create(path)
    if err != nil {
      panic(err)
    }
    io.Copy(f, out)
    f.Close()

    saved := (in.Size() - outlen) * 100 / in.Size()
    fmt.Println(fmt.Sprintf("%02d%% %s", saved, path))
  } else {
    fmt.Println(fmt.Sprintf("--- %s", path))
  }
  return true
}

func optimizeJPEG(path string) bool {
  // Read image
  finput, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  input, err := ioutil.ReadAll(finput)
  if err != nil {
    panic(err)
  }
  in := bytes.NewReader(input)
  img, err := jpeg.Decode(in)
  finput.Close()
  if err != nil {
    return false
  }

  // Encode image
  out := new(bytes.Buffer)
  err = mozjpegbin.Encode(out, img, &mozjpegbin.Options{
    Quality: 80,
    Optimize: true,
  })
  if err != nil {
    panic(err)
  }

  outlen := int64(out.Len())
  if outlen < in.Size() {
    // Write to file
    f, err := os.Create(path)
    if err != nil {
      panic(err)
    }
    io.Copy(f, out)
    f.Close()

    saved := (in.Size() - outlen) * 100 / in.Size()
    fmt.Println(fmt.Sprintf("%02d%% %s", saved, path))
  } else {
    fmt.Println(fmt.Sprintf("--- %s", path))
  }

  return true
}

func main() {
  re, err := regexp.Compile("\\/original\\/")
  if err != nil {
    panic(err)
  }

  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil {
      panic(err)
    }
    if re.Find([]byte(path)) != nil {
      // fmt.Println(fmt.Sprintf("Skipping %s", path))
      return nil
    }

    if info.IsDir() {
      return nil
    }
    ext := strings.ToLower(filepath.Ext(info.Name()))
    if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
      success := optimizePNG(path)
      if !success {
        success = optimizeJPEG(path)
      }

      if !success {
        fmt.Println(fmt.Sprintf("Skipping %s: Unknown file format", path))
      }
    }
    return nil
  })
}
