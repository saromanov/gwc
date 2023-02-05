package core

import (
	"bufio"
	"os"
	"io"
)
// ReadFile provides reading of the file
func ReadFile(path string) (int64, error) {
    f, err := os.Open(path)
    if err != nil {
        return 0, err
    }
    defer f.Close()
    r := bufio.NewReader(f)

    nr := int64(0)
    buf := make([]byte, 0, 4*1024)
    for {
        n, err := r.Read(buf[:cap(buf)])
        buf = buf[:n]
        if n == 0 {
            if err == nil {
                continue
            }
            if err == io.EOF {
                break
            }
            return nr, err
        }
        nr += int64(len(buf))

        if err != nil && err != io.EOF {
            return nr, err
        }
    }
    return nr, nil
}
