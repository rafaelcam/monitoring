package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"net/http"
	"time"
	"io"
	"io/ioutil"
)

const amountMonitoring = 5
const sleepMonitoring = 5

func main() {

	for {
		printIntroduction()
		printOptions()
		option := chooseOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Option Not Found.")
			os.Exit(-1)
		}

		fmt.Println("\n\n\n")
	}

}

func printIntroduction() {
	name := "João Rafael"
	version := 1.1

	fmt.Println("Hello Sr.", name)
	fmt.Println("Software Version", version)
}

func printOptions() {
	fmt.Println("\n---- Options ----")
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - View Logs")
	fmt.Println("0 - Exit")
}

func chooseOption() int {
	var option int

	fmt.Print("\nChoose an option: ")
	fmt.Scan(&option)

	fmt.Println("The command chosen was", option)

	return option
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	urls := readUrlsFromFile()

	for i := 0; i < amountMonitoring; i++ {
		for _, url:= range urls {
			testUrl(url)
		}
		time.Sleep(sleepMonitoring * time.Second)
	}
}

func testUrl(url string) {
	response, err := http.Get(url)
	resolveError(err)

	if response.StatusCode == 200 {
		fmt.Println("The url", url, "is Online.")
		registerLog(url, true)
	} else {
		fmt.Println("The url", url, "is Offline. Status code is", response.StatusCode, ".")
		registerLog(url, false)
	}
}

func readUrlsFromFile() []string {
	var urls []string

	file, err := os.Open("URLs.txt")
	resolveError(err)

	reader := bufio.NewReader(file)

	for {
		url, err := reader.ReadString('\n');

		url = strings.TrimSpace(url)
		urls = append(urls, url)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return urls
}

func registerLog(url string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	resolveError(err)

	message := time.Now().Format("[02/01/2006 15:04:05]") + " The url "+ url + " is "

	if status {
		message += "Online\n"
	} else {
		message += "Offline\n"
	}

	file.WriteString(message)
	file.Close()
}

func printLogs()  {
	fmt.Println("Printing logs...")

	file, err := ioutil.ReadFile("log.txt")
	resolveError(err)

	fmt.Println(string(file))
}

func resolveError(err error)  {
	if err != nil {
		fmt.Println("Unespected Error Ocurred.", err)
		os.Exit(-1)
	}
}