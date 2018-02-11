package main

import (
  "fmt"
  "image/png"
  "image/jpeg"
  "os"
  "path/filepath"
  "regexp"
  "strings"

  pngquant "github.com/yusukebe/go-pngquant"
  "github.com/nickalie/go-mozjpegbin"
)

func optimizePNG(path string) bool {
  finput, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  img, err := png.Decode(finput)
  finput.Close()
  if err != nil {
    return false
  }
  cimg, err := pngquant.Compress(img, "1")
  if err != nil {
    panic(err)
  }
  f, err := os.Create(path)
  defer f.Close()
  err = png.Encode(f, cimg)
  if err != nil {
    panic(err)
  }
  return true
}

func optimizeJPEG(path string) bool {
  finput, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  img, err := jpeg.Decode(finput)
  finput.Close()
  if err != nil {
    return false
  }
  f, err := os.Create(path)
  defer f.Close()
  err = mozjpegbin.Encode(f, img, nil)
  if err != nil {
    panic(err)
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
      fmt.Println(fmt.Sprintf("Skipping %s", path))
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

      if success {
        fmt.Println(fmt.Sprintf("Processing %s", path))
      } else {
        fmt.Println(fmt.Sprintf("Skipping %s: Unknown file format", path))
      }
    }
    return nil
  })
}
