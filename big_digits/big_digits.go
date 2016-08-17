//big digits only 0 - 3
package main

import (
  "fmt"
  "os"
  "log"
  "path/filepath"
)

var bigDigits = [][]string{
  {
  "   000   ",
  " 0     0 ",
  "0       0",
  "0       0",
  "0       0",
  " 0     0 ",
  "   000   " },
  {
  "   111   ",
  "    11   ",
  "    11   ",
  "    11   ",
  "    11   ",
  "    11   ",
  " 1111111 "},
  {
  "  222222 ",
  " 2      2",
  "        2",
  "       2 ",
  "     2   ",
  "   2     ",
  " 22222222"},
  {
  "  333333 ",
  " 3      3",
  "        3",
  "  333333 ",
  "        3",
  " 3      3",
  "  333333 "},
}

func main() {
  if len(os.Args) == 1 {
    fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }

  stringOfDigits := os.Args[1]
  for row := range bigDigits[0] {
    line := ""
    for column := range stringOfDigits {
      digit := stringOfDigits[column] - '0'
      if 0 <= digit && digit <= 3 {
        line += bigDigits[digit][row] + " "
      } else {
        log.Fatal("Numbers must be between 0..3")
      }
    }
    fmt.Println(line)
  }
}
