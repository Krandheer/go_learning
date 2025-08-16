package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func getLinesChannel(f io.Reader) <-chan string {
    out := make(chan string, 1)

    go func() {
        defer close(out)
        var buf bytes.Buffer
        chunk := make([]byte, 8)

        for {
            n, err := f.Read(chunk)

            if n > 0 {
                data := chunk[:n]
                for {
                    if i := bytes.IndexByte(data, '\n'); i >= 0 {
                        buf.Write(data[:i])
                        out <- buf.String() 
                        buf.Reset()

                        if i+1 < len(data) {
                            data = data[i+1:]
                            continue
                        }
                        break 
                    }
                    buf.Write(data)
                    break
                }
            }

            if err != nil {
                if errors.Is(err, io.EOF) {
                    if buf.Len() > 0 {
                        out <- buf.String()
                    }
                    return 
                }
                return
            }
        }
    }()

    return out 
}

func ReadByte() {
    f, err := os.Open("hello.txt")
    if err != nil {
        panic(err)
    }
	defer f.Close()
    lines := getLinesChannel(f)

    for line := range lines {
        fmt.Println("Line:", line)
    }
}
