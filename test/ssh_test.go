package test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime/debug"
	"sync"
	"testing"
)

func readLog(wg *sync.WaitGroup, out chan string, reader io.ReadCloser) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r, string(debug.Stack()))
		}
	}()
	defer wg.Done()
	r := bufio.NewReader(reader)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF || err != nil {
			return
		}
		out <- string(line)
	}
}

// RunCommand run shell
func RunCommand(out chan string, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(2)
	go readLog(&wg, out, stdout)
	go readLog(&wg, out, stderr)
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func TestPING(t *testing.T) {
	out := make(chan string)
	defer close(out)
	go func() {
		for {
			str, ok := <-out
			if !ok {
				break
			}
			fmt.Println(str)
		}
	}()
	args := []string{"-c", "ping www.5bug.wang"}
	if err := RunCommand(out, "bash", args...); err != nil {
		return
	}

}
