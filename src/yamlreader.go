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

func parseEntry(lineno int, line string) (entryLine string) {
    return line + "\n"
}

func parseDate(lineno int, line string) (date string) {
    var words = strings.Fields(line)
    fmt.Println(lineno, " ", words, len(words))

    if words[0] == "#" {
        date = strings.Join(words[1:], " ")
    }

    return date
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

    return key, value
}

func parseFile(path string) (header map[string]string, body map[string]string) {
    lines, err := scanLines(path)
    if err != nil {
        log.Fatalf("scanLines: %s", err)
    }

    mapHdr := make(map[string]string)
    mapBody := make(map[string]string)

    var parsingYaml bool
    var parsingEntry bool
    var k string
    var v string
    var date string

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
            mapHdr[k] = v
        }

        if len(line) > 0 {
            if line[0] == '#' {
                date = parseDate(i, line)
                parsingEntry = true
                continue
            }
        }
        if parsingEntry == true {
            mapBody[date] += parseEntry(i, line)
        }
    }

    return mapHdr, mapBody
}

func summarize(from string, to string) (summary string) {

    var s string

    return s
}

func main() {
    filename := `../test/test1.md`

    var header map[string]string
    var body map[string]string
    header, body = parseFile(filename)

    /*
    fmt.Println(header, len(header))
    fmt.Println(body, len(body))
    */

    var fromDate string
    var toDate string
    fromDate = "2014-01-01"
    toDate = "2014-03-03"


    /* do something with the header/body now... */
    var summary string
    summary = summarize(fromDate, toDate)

    fmt.Println(summary)
}
