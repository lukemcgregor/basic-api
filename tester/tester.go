package main

//if you actually want a load tester in go, https://github.com/rakyll/hey/ seems pretty good. I wrote this to learn things not to make something cool :)
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/montanaflynn/stats"
)

type numbers struct {
	Value []int `xml:"value"`
}

type sum struct {
	Sum int64 `xml:",chardata" json:"sum"`
}

const numberOfCallsPerThread = 125000
const numberOfThreads = 4 //matches my core count

func main() {
	var chans [numberOfThreads]chan time.Duration
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = numberOfThreads

	start := time.Now()
	fmt.Println(start, "	Starting, Threads: ", numberOfThreads, " total calls: ", numberOfThreads*numberOfCallsPerThread)

	//Spawn off some threads to slam the API with maximum hulkyness
	for i := range chans {
		chans[i] = make(chan time.Duration, 5) //only let one thread get 5 calls ahead
		go callGet(numberOfCallsPerThread, chans[i])
	}

	var durations stats.Float64Data = make([]float64, numberOfCallsPerThread*numberOfThreads)

	// Wait for all our threads to complete
	printPercentage := 0
	for i := 0; i < numberOfCallsPerThread; i++ {
		if (numberOfCallsPerThread/100)*printPercentage < i {
			fmt.Println(time.Now(), "	", printPercentage, "%")
			printPercentage += 10
		}
		for j := range chans {
			duration := <-chans[j]
			ms := float64(duration) / float64(time.Millisecond) //floaty milliseconds
			durations[(i*numberOfThreads)+j] = ms
		}
	}
	elapsed := time.Since(start)

	//Stats it up
	fmt.Println("total time: ", elapsed)

	max, _ := durations.Max()
	median, _ := durations.Median()
	mean, _ := durations.Mean()
	fmt.Println("median request time:	", median, "ms")
	fmt.Println("mean request time:	", mean, "ms")
	fmt.Println("max request time:	", max, "ms")

}

func callGet(numberOfTimes int, responseTime chan time.Duration) {
	for i := 0; i < numberOfTimes; i++ {
		var result sum
		x := rand.Intn(100)
		y := rand.Intn(100)

		start := time.Now()
		response, httpError := http.Get(fmt.Sprintf("http://localhost:8080/add/%d/to/%d", x, y))
		if httpError != nil {
			close(responseTime)
			panic(httpError)
		}
		responseBody, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()

		json.Unmarshal(responseBody, &result)

		if result.Sum != int64(x+y) {
			close(responseTime)
			panic("The sum didnt match")
		}

		responseTime <- time.Since(start)
	}
	close(responseTime)
}
