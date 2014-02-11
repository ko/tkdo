/*
func ReadString(filename string) {
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    r := bufio.NewReader(f)
    line, err := r.ReadString('\n')
    for err == nil {
        fmt.Println(line)
        line, err = r.ReadString('\n')
    }

    if err != io.EOF {
        fmt.Println(err)
        return
    }
}

func scanGoTokens(s scanner.Scanner, i int, line string) {
    s.Init(strings.NewReader(line))
    tok := s.Scan()
    for tok != scanner.EOF {
        fmt.Println(i, " scantok=", s.TokenText())
        tok = s.Scan()
    }
}
*/


