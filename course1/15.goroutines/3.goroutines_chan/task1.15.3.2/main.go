package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://httpbin.org/get"
	requestsAtOneTimeLimit := 5
	requestCount := 50

	result := benchRequest(url, requestsAtOneTimeLimit, requestCount)

	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("Ошибка при отправке запроса: %d", statusCode))
		}
	}

	fmt.Println("Все горутины завершили работу")
}

func benchRequest(url string, parallelRequest int, requestCount int) <-chan int {
	lim := make(chan struct{}, parallelRequest)
	res := make(chan int, requestCount)

	for i := 0; i < requestCount; i++ {
		lim <- struct{}{}
		go func(url string) {
			status, err := httpRequest(url)
			if err != nil {
				fmt.Println(err)
			}
			<-lim
			res <- status
		}(url)
	}

	return res
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
