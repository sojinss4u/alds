package main

import (
	"log"
	"os"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func BubbleSort(ar []int) []int {
	l := len(ar)
	for i := 0; i < l-1; i++ {
		sorted := false
		for j := 0; j < (l - 1 - i); j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				sorted = true
			}
		}
		if sorted == false {
			break
		}
	}
	return ar
}

func main() {
	ar := []int{10, 2, 5, 3, 7}
	r := BubbleSort(ar)
	infoLogger.Printf("Sorted Array: %v", r)
}

