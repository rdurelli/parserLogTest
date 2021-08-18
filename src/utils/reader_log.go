package utils

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/rdurelli/loggingParser/src/repositories"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rdurelli/loggingParser/src/lib"
	"github.com/rdurelli/loggingParser/src/parser"

	"github.com/rdurelli/loggingParser/src/models"
)

type readerLog struct {
	LogPath string
	Logs    *[]models.Log
	db      lib.Db
	logger  lib.Logger
}

func NewReaderLog(LogPath string, Logs *[]models.Log, db lib.Db, logger lib.Logger) readerLog {
	return readerLog{
		LogPath: LogPath,
		Logs:    Logs,
		db:      db,
		logger:  logger,
	}
}

func (rL readerLog) Execute() {
	rL.logger.Info("Start the method Execute")
	file, err := os.OpenFile(rL.LogPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		rL.logger.Error("Not possible to open the Log file. Reasons: ", err)
	}
	defer file.Close()

	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 500*1024)
		return lines
	}}
	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	r := bufio.NewReader(file)
	var wg sync.WaitGroup

	for {

		buf := linesPool.Get().([]byte)
		n, err := r.Read(buf)
		buf = buf[:n]
		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
			return
		}
		nextUntillNewline, err := r.ReadBytes('\n') //read entire line

		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		wg.Add(1)
		go func() {

			//process each chunk concurrently
			//start -> log start time, end -> log end time

			rL.Parse(buf, &linesPool, &stringPool)
			wg.Done()

		}()
	}
	wg.Wait()
}

func (rl readerLog) Parse(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool) {
	var wg2 sync.WaitGroup

	logs := stringPool.Get().(string)
	logs = string(chunk)

	linesPool.Put(chunk)

	logsSlice := strings.Split(logs, "\n")

	stringPool.Put(logs)

	chunkSize := 100
	n := len(logsSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}

	for i := 0; i < (noOfThread); i++ {

		wg2.Add(1)
		go func(s int, e int) {
			defer wg2.Done() //to avaoid deadlocks
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}
				is := antlr.NewInputStream(text)
				lexer := parser.NewLoggingLexer(is)

				stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

				p := parser.NewLoggingParser(stream)
				antlr.ParseTreeWalkerDefault.Walk(&LogParser{
					Logs:       rl.Logs,
					Repository: repositories.NewLogReposiry(rl.db),
				}, p.Start())

			}

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}

	wg2.Wait()
	logsSlice = nil

}
