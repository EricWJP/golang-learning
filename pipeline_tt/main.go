package main

import (
	"fmt"
	pL "./pipeline"
	"time"
	"strconv"
)

func main() {
	var pipeline *pL.Pipeline
	pipeline = pL.NewPipeline(3, 1)
	go func() {
		select {
		case <-pipeline.ChanExit:
			err := pipeline.Wait()
			if err != nil {
				fmt.Print(err)
			}
			pipeline = pL.NewPipeline(3, 1)
		}
	}()

	cnt := 0
	count := 0
	for {
		ok := pipeline.Async(func() error {
			fmt.Println(strconv.Itoa(count), strconv.Itoa(count), strconv.Itoa(count), strconv.Itoa(count))
			return nil
		}, func() error {
			fmt.Println(strconv.Itoa(count+1), strconv.Itoa(count+1), strconv.Itoa(count+1), strconv.Itoa(count+1), strconv.Itoa(count+1))
			time.Sleep(time.Second * 3)
			if cnt >5 {
				return 
			}
			return nil
		})

		cnt++

		if !ok {
			count++
			if count > 10 {
				fmt.Println("The retry exceeded the maximum limit number: ", 10)
				break
			} else {
				pipeline = pL.NewPipeline(3, 1)
			}
		}
	}
}
