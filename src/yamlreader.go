package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "bufio"
    "time"
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

    const shortForm = "2006-01-02";
    var words = strings.Fields(line)
    var t time.Time

    if words[0] == "#" {
        date = strings.Join(words[1:], " ")
        t = dateToTime(date)


    }
    return t.Format(shortForm)
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

func dateToTime(date string) (parsed time.Time) {
    const shortForm = "2006-01-02"

    t, err := time.Parse(shortForm, date)
    if err != nil {
        fmt.Println(err)
        return
    }

    return t
}

func summarize(body map[string]string, from string, to string) (summary string) {

    const shortForm = "2006-01-02"
    var t time.Time

    var year int
    var month time.Month
    var day int

    t = dateToTime(from)
    year, month, day = t.Date()

    // for every day FROM to TO...
    summary += body[t.Format(shortForm)]

    fmt.Println(year,month,day)

    return summary
}

func main() {
    filename := `../test/test1.md`

    var header map[string]string
    var body map[string]string
    header, body = parseFile(filename)

    fmt.Println(header, len(header))
    fmt.Println(body, len(body))

    var fromDate string
    var toDate string
    fromDate = "2014-01-01"
    toDate = "2014-01-03"


    /* do something with the header/body now... */
    var summary string
    summary = summarize(body, fromDate, toDate)

    fmt.Println("summary...")
    fmt.Println(summary)
}
