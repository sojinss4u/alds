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

func FindPeak(ar []int) int {
	// {3, 5, 4, 1}
	// Handling the corner cases at the beginning itself to avoid out of bound error for array.
	// This happens when the peak is ar[0] or ar[n-1] ie first or last element
	n := len(ar)
	if ar[0] > ar[1] {
		return ar[0]
	} else if ar[n-1] > ar[n-2] {
		return ar[n-1]
	}
	// Case 1: ar[mid-1] < ar[mid] > ar[mid+1]
	// Case 2: ar[mid-1] < ar[mid] < ar[mid+1]
	// Case 3: ar[mid-1] > ar[mid] < ar[mid+1]
	// Case 4: ar[mid-1] > ar[mid] > ar[mid+1]
	low := 0
	high := len(ar) - 1
	for low <= high {
		mid := (low + high) / 2
		// Case 1:
		if ar[mid] > ar[mid-1] && ar[mid] > ar[mid+1] {
			return ar[mid]
			// Case 2:
		} else if ar[mid+1] > ar[mid] && ar[mid-1] < ar[mid] {
			// Go to right side
			low = mid + 1
			// Case 3:
		} else if ar[mid-1] > ar[mid] && ar[mid+1] > ar[mid] {
			// Both left & right are increasing, so we can go either side. He we chose left side
			high = mid - 1
			// Case 4:
		} else if ar[mid-1] > ar[mid] && ar[mid+1] < ar[mid] {
			// Go to left side
			high = mid - 1
		}
	}
	return math.MinInt
}

func FindUnique(ar []int) int {
	// Q: In given array of elements, each element occurs twice except for one element. Find the unique element.
	// Note: Duplicate elements are adjacent to each other. 
	// 
	// Time Complexity: O(log(n))
	// Space Complexity: O(1)
	
	// EdgeCase1: Check if array has a single element. If yes, return ar[0] as answer

	n := len(ar)

	if n == 1 {
		return ar[0]
	}
	// EdgeCase2: Check if first or last element is unique element. If yes, retun them to avoid out of index errors in Binary Search

	if ar[0] != ar[1] {
		return ar[0]
	}

	if ar[n-1] != ar[n-2] {
		return ar[n-1]
	}
	// If above conditions are not satsfied we can start Binary Search
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		// Check if ar[mid] itself is the answer
		if ar[mid] != ar[mid-1] && ar[mid] != ar[mid+1] {
			return ar[mid]
		}
		// Check if we have landed on second occurrence of repeating element. If yes, change mid to mid-1
		if ar[mid] == ar[mid-1] {
			mid = mid - 1
		}
		// Now check if we are at the right side or left side by checking if mid is oven or odd
		if mid%2 == 0 {
			// We are at left side, so go to right side. Also skip second occurrence of mid by adding 2, instead of 1
			low = mid + 2
		} else {
			// We are at right sie, so go to left side. Since we at first index & we are moving left, we don't reduce 2, but reduce 1
			high = mid - 1
		}
	}
	return math.MinInt
}

// This function does normal binary search & return true|false
func BS(ar []int, e int) bool {
	n := len(ar)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if ar[mid] == e {
			return true
		} else if ar[mid] > e {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func FindElementAfterKRotation(ar []int, k, e int) bool {
	// Given an array which is formed by rotating a distinct sorted array by 'k' times, Search for a given element in the rotated array?
	// Note: You are given with values of 'k' & target.
	// I/P: {4, 5, 6, 1, 2, 3}
	n := len(ar)
	// First Binary Search
	// ar = {4, 5, 6}, e = 3
	rl := BS(ar[0:k], e)
	// Second Binary Search
	// ar = {1,2,3}, e = 3
	rr := BS(ar[k:n], e)
	if rl || rr {
		return true
	}
	return false
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
        ar1 = []int{3, 4, 1, 2}
	r = FindPeak(ar1)
	infoLogger.Print(r)
        ar1 = []int{3, 3, 1, 1, 8, 8, 10, 10, 19, 6, 6, 2, 2, 4, 4}
	r = FindUnique(ar1)
	infoLogger.Print(r)
        ar1 = []int{4, 5, 6, 1, 2, 3}
	r1 := FindElementAfterKRotation(ar1, 3, 7)
	infoLogger.Print(r1)
}
