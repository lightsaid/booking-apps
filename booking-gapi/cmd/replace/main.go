package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
)

// 这个程序时为了解决 sqlc 生成的代码后，将 sql.NullXXX 类型转换为 *XXX

// go run ./cmd/replace_null/*.go
// goimports -w ./db/postgres
func main() {
	var dirpath string
	flag.StringVar(&dirpath, "path", "./db/postgres", "需要替换sql.NullXXX的文件目录")
	flag.Parse()
	files, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal("os.ReadDir: ", err)
	}
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file fs.DirEntry) {
			if !file.IsDir() {
				path := dirpath + "/" + file.Name()
				f, err := os.Open(path)
				if err != nil {
					log.Fatal("os.Open: ", err)
				}
				defer f.Close()

				scan := bufio.NewScanner(f)
				var lines []string
				for scan.Scan() {
					s := strings.ReplaceAll(scan.Text(), "sql.NullString", "*string")
					s = strings.ReplaceAll(s, "sql.NullBool", "*bool")
					s = strings.ReplaceAll(s, "sql.NullByte", "*byte")
					s = strings.ReplaceAll(s, "sql.NullFloat64", "*float64")
					s = strings.ReplaceAll(s, "sql.NullInt16", "*int16")
					s = strings.ReplaceAll(s, "sql.NullInt32", "*int32")
					s = strings.ReplaceAll(s, "sql.NullInt64", "*int64")
					s = strings.ReplaceAll(s, "sql.NullTime", "*time.Time")
					lines = append(lines, s)
				}
				if err = scan.Err(); err != nil {
					log.Fatal("scan.Err(): ", err)
				}
				err = os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0664)
				if err != nil {
					log.Fatal("os.WriteFile: ", err)
				}
			}
			wg.Done()
		}(file)
	}

	wg.Wait()
	fmt.Println("File replace success.")
}
