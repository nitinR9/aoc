package main

import (
	"2015/common"
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"time"
)

func task(input string, padlength int, ques int) {
	startNow := time.Now()
	var count uint32 = 1
	for count < math.MaxUint32 {
		data := fmt.Sprintf("%v%d", input, count)
		hasher := md5.New()
		_, err := io.WriteString(hasher, data)

		if err != nil {
			fmt.Println("Error occured during hashing", err)
			return
		}

		md5 := fmt.Sprintf("%x", hasher.Sum(nil))

		if md5[:padlength] == strings.Repeat("0", padlength) {
			break
		}

		count++
	}

	fmt.Printf("Part%d: %d\n", ques, count)
	fmt.Printf("This part %d took: %v\n", ques, time.Since(startNow))
}

func searchNum(input string, job uint32, start uint32, end uint32, len int, w *sync.WaitGroup, ch chan<- uint32) {
	defer w.Done()
	var min uint32 = math.MaxUint32

	count := start
	for count < end {
		data := fmt.Sprintf("%v%d", input, count)
		hasher := md5.New()
		_, err := io.WriteString(hasher, data)

		if err != nil {
			fmt.Println("Error occured during hashing", err)
			return
		}

		md5 := fmt.Sprintf("%x", hasher.Sum(nil))

		if md5[:len] == strings.Repeat("0", len) {
			break
		}

		count++
	}

	if count < min && count != end {
		ch <- count
	} else {
		ch <- 0
	}
}

func taskRoutines(input string, padlength int, ques int) {
	startNow := time.Now()
	var chunkSize uint32 = 10000000 / 8
	var start uint32
	var end uint32
	var wg sync.WaitGroup
	ch := make(chan uint32)

	for i := uint32(0); i < 8; i++ {
		start = i * chunkSize
		end = (i + 1) * chunkSize
		wg.Add(1)
		go searchNum(input, i, start, end, padlength, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var min uint32 = 0

	for val := range ch {
		if val != 0 {
			min = val
			break
		}
	}

	fmt.Printf("Part%d: %d\n", ques, min)
	fmt.Printf("The part %d took: %v\n", ques, time.Since(startNow))
}

func main() {
	input := common.GetFile("input.txt")
	task(input, 5, 1)
	task(input, 6, 2)
	taskRoutines(input, 6, 3)
}
