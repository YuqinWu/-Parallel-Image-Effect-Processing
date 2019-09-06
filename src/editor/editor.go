package main

import (
	"os"
	"../pngimg"
	"sync"
	"fmt"
	"strings"
	"strconv"
	"runtime"
	"bufio"
)
var SEQUENTIAL int = 1
type goContext struct {
	mutex       *sync.Mutex
	cond        *sync.Cond
	wg   		*sync.WaitGroup
	jobs 		[][]string
	imgs 		[]*pngimg.PNGImage
	results 	[]*pngimg.PNGImage
	outer		int
	next		int
	step		int
	size 		int
	counter     int
	threadCount int
	done 		bool
}
type resultContext struct {
	mutex       *sync.Mutex
	cond        *sync.Cond
	results 	[]*pngimg.PNGImage
	next		int
	size 		int
}


func distributeWork(context *goContext, task string) string{
	// Process the string, split into each element.
	elements := strings.Split(task, ",")

	pngImg, err := pngimg.Load(elements[0])
	if err != nil {
		panic(err)
	}
	// put corresponding element into parallelism context.
	context.mutex.Lock()
	context.imgs = append(context.imgs, pngImg)
	context.jobs = append(context.jobs, elements[2:len(elements)])
	context.size++
	context.mutex.Unlock()

	return elements[1]
}
func storingImg(storeAddress []string, resultContext *resultContext, context *goContext){
	total := context.size
	for resultContext.next != total {
		resultContext.mutex.Lock()
		for resultContext.next >= resultContext.size{
			resultContext.cond.Wait()
		}
		resultContext.mutex.Unlock()

		for resultContext.next < resultContext.size {
			mycurrent := resultContext.next
			pngImg := resultContext.results[mycurrent]
			pngImg.Save(storeAddress[mycurrent])
			resultContext.next++
		}
	}

}

func workerThreads(context *goContext, resultctx *resultContext, threadID int) {
	for context.done == false || context.next < context.size{
		// Check if there is available job block. Wait if not.
		context.mutex.Lock()
		//fmt.Println("in image locked: ", context.next, context.size, threadID)
		for context.next >= context.size{
			context.cond.Wait()
		}
		context.mutex.Unlock()
		//fmt.Println("in image unlocked: ", context.next, context.size, threadID)
		// If there is at least one available job, grab it from the channel.
		for context.next < context.size {
			//fmt.Println("in for unlocked: ", context.next, context.size, threadID)
			//context.mutex.Lock()
			mycurrent := context.next
			//context.mutex.Unlock()
			pngImg := context.imgs[mycurrent]
			effects := context.jobs[mycurrent]
			//fmt.Println("mycurrent: ", context.next, context.size, threadID)
			// Start processing it.
			for _, eachTask := range(effects){
				//fmt.Println("in effects: ", eachTask, context.counter, context.step, threadID)
				//context.mutex.Lock()
				current := context.step
				//context.mutex.Unlock()
				pngImg.ProcessImg(eachTask, true, context.threadCount, threadID)

				// barrier to wait untill all threads finish their effect.
				context.mutex.Lock()
				context.counter++
				//fmt.Println("in inner barrier: ", eachTask, context.counter, context.step, threadID)
				if context.counter == context.threadCount{
					context.counter = 0
					context.step++
					//fmt.Println("in inner barrier IF: ", eachTask, context.counter, context.step, threadID)
					pngImg.ReLoad()
					if context.threadCount != 1{
						context.cond.Broadcast()
					}
					//fmt.Println("inner Broadcasted: ", eachTask, context.counter, context.step, threadID)
				} else {
					for context.counter != context.threadCount && current == context.step{
						context.cond.Wait()
					}
				}
				context.mutex.Unlock()
				//fmt.Println("out effects: ", eachTask, context.counter, context.step, threadID)
			}
			//fmt.Println("out ALL effects unlocked: ", context.counter, context.next, context.size, threadID)
			// it's time to move to next task with an outer barrier
			context.mutex.Lock()
			context.outer++
			if context.outer == context.threadCount{
				context.counter = 0
				context.step = 0
				context.outer = 0

				resultctx.mutex.Lock()
				resultctx.results = append(resultctx.results, pngImg)
				resultctx.size++
				resultctx.mutex.Unlock()
				resultctx.cond.Signal()

				context.next++
				context.cond.Broadcast()
			} else{
				for context.outer != context.threadCount && mycurrent == context.next{
					context.cond.Wait()
				}
			}
			context.mutex.Unlock()
			//fmt.Println("finished, next image: ", context.next, context.size, threadID)
		}
		//fmt.Println("finished all: ", context.next, context.size, threadID)
	}
	context.wg.Done()
}

func main() {

	// this is sequential case.
	if len(os.Args) == 2 {
		// Open file for scanning.
		f, err := os.Open(os.Args[1])
		if err != nil {
		    fmt.Println("error opening file =", err)
		    os.Exit(1)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		//Start scanning
		for scanner.Scan() {
			// Start processing it.
			elements := strings.Split(scanner.Text(), ",")
			// Loads the png image and returns the image or an error
			pngImg, err := pngimg.Load(elements[0])
			// Process effects
			for _, eachTask := range(elements[2:len(elements)]){
				pngImg.ProcessImg(eachTask, false, SEQUENTIAL, SEQUENTIAL)
				// assign out back to in for next effect
				pngImg.ReLoad()
			}
			// Then save it.
			err = pngImg.Save(elements[1])
			//Checks to see if there were any errors when saving.
			if err != nil {
				panic(err)
			}
		}
		fmt.Printf("Done.\n")
	  // for parallel case
	} else if len(os.Args) == 3 {
	
		numOfThreads := runtime.NumCPU() // the default number of threads.
		flag := strings.Split(os.Args[1], "=")

		//check if it is -p flag
		if flag[0] == "-p" {
			// check if number is provided.
			if len(flag) == 2 {
				numOfThreads, _ = strconv.Atoi(flag[1])
			}
		} else {
			fmt.Println("Usage: editor [-p=[num of threads]]  <csv file>")
			return
		}
		// Prepare for the lock.
		var wg sync.WaitGroup
		var mux sync.Mutex
		condVar := sync.NewCond(&mux)
		context := goContext{&mux, condVar, &wg, make([][]string, 0), make([]*pngimg.PNGImage, 0), make([]*pngimg.PNGImage, 0), 0, 0, 0, 0, 0, numOfThreads, false}
		var resultmux sync.Mutex
		resultcond := sync.NewCond(&resultmux)
		resultcontext := resultContext{&resultmux, resultcond, make([]*pngimg.PNGImage, 0), 0, 0}

		// Initialize the number of threads.
		for i := 0; i < numOfThreads; i++ {
			wg.Add(1)
			go workerThreads(&context, &resultcontext, i)
		}

		// Perpare for reading tasks.
		f, err := os.Open(os.Args[2])
		if err != nil {
		    fmt.Println("error opening file= ", err)
		    os.Exit(1)
		}
		defer f.Close()

		// Declare an array to store the address of saving images.
		// This is unnecessary to share among threads so make it local.
		var storeAddress []string
		
		//Start scanning
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			store := distributeWork(&context, scanner.Text())
			// put the store address into the local array.
			storeAddress = append(storeAddress, store)
			condVar.Broadcast()
		}
		// Jobs are done, notify all thread to quit
		context.done = true
		condVar.Broadcast()

		//Start storing imgs.
		storingImg(storeAddress, &resultcontext, &context)
		wg.Wait()
		//Wait until all finished to exit.
		fmt.Printf("Done.\n")
		
	} else {
		fmt.Println("Usage: editor [-p=[num of threads]]  <csv file>")
		return
	}
}


