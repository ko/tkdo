package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "bufio"
)

func scanLines(path string) ([]string, error) {

    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}


func errorExit(i int) {
    fmt.Println("bad format @ line ", i)
    os.Exit(1)
}

func parseLine(lineno int, line string, words []string) {
    words = strings.Fields(line)
    fmt.Println(lineno, " ", words, len(words))

    if lineno == 0 {
        if words[0] != "---" {
            errorExit(lineno)
        }
    }
}

func parseYaml(lineno int, line string) (key string, value string) {
    var words = strings.Fields(line)
    if len(words) < 2 {
        errorExit(lineno)
    }
    if last := len(words[0]) - 1; last >= 0 && words[0][last] == ':' {
        words[0] = words[0][:last]
    }

    key = words[0]
    value = strings.Join(words[1:], " ")

    fmt.Println("key=",key)
    fmt.Println("val=",value)

    return key, value
}

func main() {
    filename := `../test/test1.md`
    //    ReadString(filename)
    lines, err := scanLines(filename)
    if err != nil {
        log.Fatalf("scanLines: %s", err)
    }

    //var s scanner.Scanner
    var words []string
    var parsingYaml bool
    kv := make(map[string]string)
    var k string
    var v string

    for i, line := range lines {
        if i == 0 {
            if line == "---" {
                parsingYaml = true
                continue
            }
        }
        if parsingYaml == true {
            if line == "---" {
                parsingYaml = false
                continue
            }
            k,v = parseYaml(i, line)
            kv[k] = v
        }
        parseLine(i, line, words)
    }
    fmt.Println(kv)
}
