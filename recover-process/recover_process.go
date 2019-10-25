package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

type RecoverProcessConfig struct {
	ProcessName            string
	ProcessStartCommand    string
	ProcessMonitorInterval int
}

func main() {
	processName := "demo-api"
	processStartCmd := "demo-api"
	procErrChan := make(chan string, 1)

	checkInterval := 10

	for {
		var wg sync.WaitGroup

		wg.Add(1)
		go checkLocalProcess(processName, procErrChan, &wg)
		wg.Wait()
		procErr := <-procErrChan

		if procErr != "" {
			fmt.Printf("Process %s died! Restarting\n", processName)
			// Start process if its died
			startLocalProcess(processStartCmd)
		} else {
			fmt.Printf("Process %s Running! \n", processName)
		}
		time.Sleep(checkInterval * time.Second)
	}

}

func startLocalProcess(processStartCmd string) {
	c1 := exec.Command(processStartCmd, "&")
	err := c1.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	//c1.Wait()

	var b2 bytes.Buffer
	c1.Stdout = &b2
	//Println(&b2)
	fmt.Println("Process runnin with PID ", c1.Process.Pid)
}

func checkLocalProcess(processName string, procErrChan chan string, wg *sync.WaitGroup) {
	fmt.Printf("Checking for process %s \n", processName)

	checkCmd := exec.Command("pgrep", processName)

	var b2 bytes.Buffer
	checkCmd.Stdout = &b2

	checkCmd.Start()
	checkCmd.Wait()

	//Println(&b2)
	lines, err := lineCounter(&b2)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(0)
	}

	// Mark this task as done
	// Error if done after passing data to channels
	wg.Done()

	if lines == 0 {
		fmt.Printf("Error: no process %s found running!\n", processName)
		procErrChan <- "process does not exist"
	}

	procErrChan <- ""
}

// Linecounter counts the number of lines in a given output
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		fmt.Println(string(buf))
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}