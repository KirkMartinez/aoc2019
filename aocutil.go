package main

import (
  "bufio"
  "log"
  "os"
  "strconv"
)

func getData(day int) []string {
  var result []string;

  file, err := os.Open("data/day"+strconv.Itoa(day))
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    result = append(result, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return result
}
