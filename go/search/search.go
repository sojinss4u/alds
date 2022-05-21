package main

import (
	"log"
	"os"
	"math"
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

func UnorderedLinerarSearch(k interface{}, ar []interface{}) int {
	// Here the i/p array is not sorted & hence we don't know the positions of the elements in the array
	// So we will need to scan the entire array to find a given element 'k', in linear fashion
	// Function return index of element if it is found, else return -1
	for i := 0; i < len(ar); i += 1 {
		if ar[i] == k {
			return i
		}
	}
	return -1
}

func FindGreatest(k int, ar []int) int {
        // Q: Given a sorted array find the greatest element which is <= k ?
	// We will apply search here as we have a search space [i/p array] & we have a target [k].
	// We will apply Binary Search here, as the i/p array is sorted & we can discard some part of the
	// array after reaching the mid portion. ie Once we reach the mid, we will check
	// if ar[mid] == k, if yes we will just return ar[mid]
	// if ar[mid] > k, we don't need to check the right part of ar[mid] & we can completely discard right part
	// if ar[mid] < k, ar[mid] can be an answer, so we will update the answer with this value. So will be possibly
	// able to find a better answer on the right side of ar[mid], so we will searching for a new value in the right space,
	// untill we have only one element in the search space ie when low == high. We don't need to check if the newly found
	// element is > previous max element as array is sorted & if you go right part the element will be obviously greater
	low := 0
	high := len(ar) - 1
	secondHighest := math.MinInt
	for low <= high {
		mid := (high + low) / 2
		if ar[mid] == k {
			secondHighest = ar[mid]
			return secondHighest
		} else if ar[mid] > k {
			high = mid - 1
		} else {
			low = mid + 1
			secondHighest = ar[mid]
		}
	}
	return secondHighest
}

func FindFrequency(k int, ar []int) int {
	// Q: Given a sorted array of 'n' elements, find the frequency of a given element 'k'
	// Since we have a search space & target the questions is a search question
	// Assume that we want to find the element 5 in array ar = [1,2,3,4,5,5,5,7,8,9,10,12]
	// Since the array is sorted & we can discard some searh space after finding the mid element, we can use BinarySearch
	// Now we need to find the index of the element 5 using BinarySeatrch, then iterate backward & forward check if ar[i]= 5, and update frequency
	// The problem with this approach is it will still take O(n), in worst case where all the elements in the array is 'k'
	// So the optimized solution is to use two binary searches to find the index of starting k & ending k & just substract them to find the frquency
	// Frequency := Last index of k - First index of k + 1
	// In the above example, initially mid = 6. Also ar[mid] == 5. So mid can be the starting index for 5. However we can find a better answer on the left side
	// as we are looking for initial index of 5
	// Similaly initially mid = 6. Also ar[mid] == 5. So mid can be the endIndex for 5. However we can find a better answer on the right side as we are looking for last index of 5

	low := 0
	high := len(ar) - 1
	// Initialize indexes with invalid index to avoid confusion. Note: 0 is a valid index, so we can't initialize with 0
	startIndex := math.MinInt
	endIndex := math.MaxInt
	// Find Start Index here
	for low <= high {
		mid := (low + high) / 2
		if ar[mid] == k {
			startIndex = mid
			high = mid - 1
		} else if ar[mid] > k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	low = 0
	high = len(ar) - 1

	// Find endIndex here

	for low <= high {
		mid := (low + high) / 2
		if ar[mid] == k {
			endIndex = mid
			low = mid + 1
		} else if ar[mid] > k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return endIndex - startIndex + 1
}

func main() {
	ar := []interface{}{3, 1, 7, 2, 6}
	i := UnorderedLinerarSearch(6, ar)
	infoLogger.Printf("%v", i)
	ar = []interface{}{"a", "c", "b", "g", "h"}
	i = UnorderedLinerarSearch("h", ar)
	infoLogger.Printf("%v", i)
        ar1 := []int{1, 3, 7, 10}
        i = FindGreatest(11, ar1)
	infoLogger.Print(i)
	ar1 = []int{1, 3, 7, 7, 7, 10}
	r := FindFrequency(7, ar1)
	infoLogger.Print("Frequency: ", r)
}

