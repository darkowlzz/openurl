package main

import (
  "log"

  "../../openurl"
)

func main() {
  if err := openurl.Open("http://example.com"); err != nil {
    log.Fatal(err)
  }
}
