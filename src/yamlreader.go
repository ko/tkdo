package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "bufio"
    "time"
    "flag"
    "path/filepath"
)



var gDays int
var gMode string
var gDir string
var gFileList []string
var gFilename string
var gSummary bool

const (
    titleMode = "title"
    dailyMode = "daily"
    activitydaysMode = "activitydays"
)


type Task struct {
    filename string
    title string
    category string
    updates map[string]string
}

var gResultTasks []Task

func (t Task) GetTaskName() string {
    return t.title
}

func (t Task) GetCategory() string {
    return t.category
}

func (t Task) GetFilename() string {
    return t.filename
}

func (t Task) GetUpdates() map[string]string {
    return t.updates
}



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
                // TODO check if date format? 
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

func getFirstSentence(msg string) (sentence string) {
    var retstr string
    var lines []string
    var idx_ellipsis int
    var idx_dot int

    idx_dot = -1
    lines = strings.Split(msg, "\n")

    for  _, line := range lines {
        if strings.Contains(line, "...") {
            idx_ellipsis = strings.Index(line, "... ")
            if idx_ellipsis == -1 {
                idx_ellipsis = strings.Index(line, "...\n")
            }
        }
        if strings.Contains(line, ".") {
           idx_dot = strings.Index(line, ". ")
           if idx_dot == -1 {
               idx_dot = strings.Index(line, ".")
               if idx_dot != len(line) - 1 {
                  idx_dot = -1
               }
           }
        }

        if strings.Index(line, "* ") == 0 {
            retstr += "\n"
        }

        if len(line) == 0 {
            retstr += "\n"
        } else {
            retstr += " "
            if idx_dot != -1 {
                // cut at the first "." we see... and add a "."
                // at the end of this. 
                retstr += strings.Split(line,". ")[0]
                retstr += "."
                break
            } else {
                retstr += line
            }
        }
    }

    //fmt.Println("returning=", retstr)
    return retstr
}

func getUpdateSummary(body string) (summary string) {

    var paragraph []string
    var paragraphs []string
    var sentences []string
    var lines []string

    // paragraphs have empty strings/lines between them
    lines = strings.Split(body,"\n")
    for _, line := range lines {
        if len(line) > 0 {
            paragraph = append(paragraph, line)
        } else {
            // delimit with "\n" so the following "range" call will do 
            // something useful
            paragraphs = append(paragraphs, strings.Join(paragraph,"\n"))
            paragraph = nil
        }
    }

    for _, para := range paragraphs {
        // We would have empty strings show up here, otherwise. 
        if len(para) > 0 {
            sentences = append(sentences , getFirstSentence(para))
        }
    }

    return strings.Join(sentences, "\n\n")
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
    var day time.Time

    var timeFrom = dateToTime(from)
    var timeTo = dateToTime(to)

    // for every day FROM to TO...
    for day = timeFrom; day != timeTo; day = day.Add(time.Hour * 24) {
        fmt.Println(day.Format(shortForm))
        summary += body[day.Format(shortForm)]
    }


    return summary
}

func usage() {
    fmt.Fprintf(os.Stderr, "Usage: %s [days]\n", os.Args[0])
    os.Exit(1)
}


func fileToTask(filename string) (t Task) {

    const (
        title = "title"
        category = "category"
    )


    var header, body map[string]string
    header, body = parseFile(filename)

    t.title = header[title]
    t.category = header[category]
    // TODO basename?
    t.filename = filename
    t.updates = body

//    fmt.Println("taskname=", t.GetTaskName())

    return t
}


func appendIfMissing(tasks []Task, task Task) []Task {
    for _, v := range tasks {
        // TODO we can do better than this
        if v.GetTaskName() == task.GetTaskName() {
            return tasks
        }
    }
    return append(tasks, task)
}




func visit(path string, f os.FileInfo, err error) error {

    var bName = []byte(f.Name())

    if f.IsDir() && bName[0] == '.' {
        if len(f.Name()) == 1 || bName[1] == '.' {
            // whatever about "." and ".."
            return nil
        }
        return filepath.SkipDir
    }

    // boo directories
    if f.IsDir() {
        goto Leave
    }

    if strings.HasSuffix(path, ".md") == true {
        gFileList = append(gFileList, path)
    }

Leave:
    return nil
}





func init() {
    const (
            defaultDays = 7
            defaultMode = titleMode
            defaultDir = "."
            defaultFilename = ""
            defaultSummary = false
    )
    flag.IntVar(&gDays, "days", defaultDays, "days to look back from now")
    flag.StringVar(&gMode, "mode", defaultMode, "mode to run in")
    flag.StringVar(&gDir, "dir", defaultDir, "directory with tasks")
    flag.StringVar(&gFilename, "file", defaultFilename, "specific file to inspect")
    flag.BoolVar(&gSummary, "summary", defaultSummary, "summary of task(s)")
}

func main() {

    flag.Parse()

    const shortForm = "2006-01-02"

    //filename := `../test/test1.md`

    // time strings
    var startTime time.Time
    var endTime time.Time
    startTime = time.Now().Add(-(time.Hour * 24 * time.Duration(gDays)))
    endTime = time.Now()


    // foreach file in gDir; do visit $file
    err := filepath.Walk(gDir, visit)
    if err != nil {
        log.Fatal(err)
    }

    var files = gFileList
    for file := range files {
        var task Task
        task = fileToTask(files[file])

        // collect relevant tasks
        for k, _ := range task.updates {
            // BUG this ain't right...
            startTime.Add(-1 * time.Hour * 24)
            if dateToTime(k).Before(endTime) && dateToTime(k).After(startTime) {
                gResultTasks = appendIfMissing(gResultTasks, task)
            }
        }
    }

    if gMode == titleMode {
        var task Task
        fmt.Println("=== titles ===")
        for t := range gResultTasks {

            if gSummary == true {
                fmt.Println("---")
            }

            fmt.Printf("%s|%s\n",
                        gResultTasks[t].GetCategory(),
                        gResultTasks[t].GetTaskName())

            // if we want a summary of this task...
            if gSummary == true {
                fmt.Println("---")

                var idx int
                var counter int
                task = gResultTasks[t]
                idx = len(task.GetUpdates()) - 1
                counter = 0
                for _, msg := range task.updates {
                    if counter == idx {
                        fmt.Println()
                        fmt.Printf("%s\n", getUpdateSummary(msg))
                        fmt.Println()
                        break
                    }
                    counter++
                }
            }

        }

    } else if gMode == dailyMode {
        fmt.Println("=== daily ===")
        // let's go day by day
        for day := startTime; day.Before(endTime); day = day.Add(time.Hour * 24) {
            // and task by task
            for t:= range gResultTasks {
                // for every day a task has been updated
                for date, _ := range gResultTasks[t].GetUpdates() {
                    dateTime := dateToTime(date)
                    // find a match
                    if day.Format(shortForm) == dateTime.Format(shortForm) {
                        fmt.Printf("%s|%s|%s\n",
                                    date,
                                    gResultTasks[t].GetCategory(),
                                    gResultTasks[t].GetTaskName());
                    }
                }
            }
        }
    } else if gMode == activitydaysMode {
        fmt.Println("=== activitydays ===")
        // for each task
        for t:= range gResultTasks {

            // for this task
            if gResultTasks[t].GetFilename() == gFilename {

                // for every day this task has been updated
                for date, _ := range gResultTasks[t].GetUpdates() {

                    // print it out
                    fmt.Printf("%s|%s|%s\n",
                                date,
                                gResultTasks[t].GetCategory(),
                                gResultTasks[t].GetTaskName())
                }
            }
        }
    }

    /*
    var fromDate string
    var toDate string
    fromDate = startTime.Format(shortForm)
    toDate = endTime.Format(shortForm)

    // do something with the header/body now... 
    var summary string
    summary = summarize(task.updates, fromDate, toDate)
    */
}
