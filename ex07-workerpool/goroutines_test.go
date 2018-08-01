package goroutines

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testcases = []struct {
		MaxThreads  int
		Input       []string
		Output      []string
		WaitSeconds int
	}{
		{
			2,
			[]string{
				"0.1",
			},
			[]string{
				"worker:1 spawning",
				"worker:1 sleep:0.1",
				"worker:1 stopping",
			},
			1,
		},

		{
			2,
			[]string{
				"0.1",
				"0.2",
			},
			[]string{
				"worker:1 spawning",
				"worker:1 sleep:0.1",
				"worker:2 spawning",
				"worker:2 sleep:0.2",
				"worker:1 stopping",
				"worker:2 stopping",
			},
			1,
		},

		{
			1,
			[]string{
				"0.1",
				"0.2",
			},
			[]string{
				"worker:1 spawning",
				"worker:1 sleep:0.1",
				"worker:1 sleep:0.2",
				"worker:1 stopping",
			},
			1,
		},

		{
			10,
			[]string{
				"0.2",
				"0.3",
				"0.4",
				"0.5",
				"0.6",
				"0.7",
				"0.8",
				"0.9",
				"1.0",
				"1.1",
				"1.2",
				"1.3",
				"1.4",
			},
			[]string{
				"worker:1 spawning",
				"worker:1 sleep:0.2",
				"worker:2 spawning",
				"worker:2 sleep:0.3",
				"worker:3 spawning",
				"worker:3 sleep:0.4",
				"worker:4 spawning",
				"worker:4 sleep:0.5",
				"worker:5 spawning",
				"worker:5 sleep:0.6",
				"worker:6 spawning",
				"worker:6 sleep:0.7",
				"worker:7 spawning",
				"worker:7 sleep:0.8",
				"worker:8 spawning",
				"worker:8 sleep:0.9",
				"worker:9 spawning",
				"worker:9 sleep:1.0",
				"worker:10 spawning",
				"worker:10 sleep:1.1",
				"worker:1 sleep:1.2",
				"worker:2 sleep:1.3",
				"worker:3 sleep:1.4",
				"worker:4 stopping",
				"worker:5 stopping",
				"worker:6 stopping",
				"worker:7 stopping",
				"worker:8 stopping",
				"worker:9 stopping",
				"worker:10 stopping",
				"worker:1 stopping",
				"worker:2 stopping",
				"worker:3 stopping",
			},
			3,
		},
	}
)

func TestRun(t *testing.T) {
	test := assert.New(t)

testcases:
	for testID, testcase := range testcases {
		log.Printf("start testcase: #%d", testID)

		stdinReader, stdinWriter, err := os.Pipe()
		if err != nil {
			panic(err)
		}

		stdoutReader, stdoutWriter, err := os.Pipe()
		if err != nil {
			panic(err)
		}

		os.Stdin = stdinReader
		os.Stdout = stdoutWriter

		runner := testRun(testcase.MaxThreads)

		for _, line := range testcase.Input {
			log.Printf("    > %s", line)

			_, err := stdinWriter.WriteString(line + "\n")
			if err != nil {
				test.NoError(
					err,
					fmt.Sprintf("testcase: %d, can't write to stdin", testID),
				)
				break testcases
			}

			// In real world user will take some time to think about next
			// number, it also helps to skip concurrency hell
			time.Sleep(time.Millisecond * 10)
		}

		err = stdinWriter.Close()
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(stdoutReader)
		for number, expected := range testcase.Output {
			test.True(
				scanner.Scan(),
				fmt.Sprintf("testcase: %d, can't read stdout", testID),
			)

			actual := scanner.Text()

			log.Printf("    < %s", actual)

			test.Equal(
				expected,
				actual,
				fmt.Sprintf(
					"testcase: %d, output (line: %d) validation",
					testID,
					number+1,
				),
			)

		}

		wait(test, testID, runner, time.Duration(testcase.WaitSeconds)*time.Second)
	}
}

func testRun(poolSize int) *sync.WaitGroup {
	runner := &sync.WaitGroup{}
	runner.Add(1)

	go func() {
		defer runner.Done()
		Run(poolSize)
	}()

	return runner
}

func wait(
	test *assert.Assertions,
	testID int,
	runner *sync.WaitGroup,
	maxTime time.Duration,
) {
	done := make(chan struct{})
	go func() {
		runner.Wait()
		done <- struct{}{}
	}()

	after := time.After(maxTime)
	select {
	case <-after:
		test.FailNowf(
			"runner was not stopped",
			"testcase: %d operation timed out after %s",
			testID,
			maxTime,
		)
	case <-done:
		return
	}
}
