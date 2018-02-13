package options

import (
  "fmt"
  "os"

  "github.com/gobwas/glob"
)

type PathMatcher struct {
  Type string // 'include' or 'exclude'
  Pattern glob.Glob
}

type Options struct {
  Directories []string
  PathMatchers []PathMatcher
}

func ParseArgs(args []string) *Options {
  options := Options{
    Directories: []string{},
    PathMatchers: []PathMatcher{},
  }
  for i := 0; i < len(args); i++ {
    arg := args[i]
    if arg == "-include" || arg == "-exclude" {
      if i >= len(args) - 1 {
        panic(fmt.Sprintf("No value specified for: %s", arg))
      }
      value := args[i+1]
      var pm PathMatcher
      if arg == "-include" {
        pm = PathMatcher{ Type: "include", Pattern: glob.MustCompile(value) }
      } else {
        pm = PathMatcher{ Type: "exclude", Pattern: glob.MustCompile(value) }
      }
      options.PathMatchers = append(options.PathMatchers, pm)
      i++
    } else {
      // Directory, check if directory exist
      if _, err := os.Stat(arg); os.IsNotExist(err) {
      	panic(fmt.Sprintf("Directory '%s' does not exist", arg))
      }
      options.Directories = append(options.Directories, arg)
    }
  }
  return &options
}
