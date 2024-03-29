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
	// Q: Given an array which is formed by rotating a Distinct[All elements are unique in array] sorted array by 'k' times, Search for a given element in rotated array?
	//Note: You are given with values of 'k' & target.  
	
	// I/P: {4, 5, 6, 1, 2, 3}
	// The logic here is that if an array is rotated 'k' times, we will be able to find two sorted arrays ie ar1 = ar[0:(k-1)] & ar2=ar[k:n-1] in the resulting array.
	// Now we can apply BS on these two sorted arrays & will be able to find the given element in O(log(n)) time complexity, which is better than linerar search which is O(n)

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

func FindElementAfterKRotationWithoutKGiven(ar []int, e int) bool {
	// Since 'k' is not given, we need to first findout 'k'. This we can do using BS using a propery of ar given below.
	// Q: Given an array which is formed by rotating a Distinct[All elements are unique in array] sorted array by 'k' times, Search for a given element in rotated array?
        // Note: You are given with values of only target & 'k' is not given

	// Eg: [4,5,1,2,3], e = 1
	// ar1 = [4,5] ar2 = [1,2,3]
	// In the above example, k=2. We can also see that 'k' is same as the index of the first element in second array. So if we find ar2[0], we will get 'k'.
	// Also we can see that ar1[0] is always greater than any element in ar2.
	// So we will apply BS & calculate mid. Once we get mid, we will check if ar[mid] < ar[0]. If yes, we are in the right array ie ar2. However we are not sure if ar[mid] is
	// the first element of ar2. So we will also check if ar[mid] < ar[mid-1]. If this condition is true, we can conclude that this element is the starting index of ar2 == k.
	// If current mid is not the starting element of ar2, we will need to go left as we will be able to find a better answer in the left side.  So we will update high=mid-1.
	// Similarly if ar[mid] < ar[0]. Then we can confirm that we are in the left array ar1. So we need to go right to find k. We do this by setting low=mid+1.
	n := len(ar)
	low := 0
	high := n - 1
	k := -1
	for low <= high {
		mid := (low + high) / 2
		// Condition where we are in the right array
		if ar[mid] < ar[0] {
			if ar[mid] < ar[mid-1] {
				k = mid
				break
			} else {
				high = mid - 1
			}
		} else if ar[mid] > ar[0] {
			// Condition where wr are in the left array
			low = mid + 1
		}
	}
	// Now we can apply BS with k & e
	r := FindElementAfterKRotation(ar, k, e)
	return r
}

func FindSqrt(n int) int {
	// Q: Given 'n', find sqrt(n)? Also if the sqrt(n) is a flotat, we can round it off to the floor integer value.
	// Note: n is a positive integer [1 - intMax]
	// Idea: We can run a loop from i = 1 to n, multiplying i * i & checking if it is equal to or less than n.
	// For instance let's find the sqrt of 37. Let's start from 1 & calculate 1*1 = 1. Now 1<37, so 1 can be an answer. Now we will increment 'i' & calculate 2 * 2 = 4. Now 4 < 37, so it can be an answer. So we will update ans with 4. We will keep doing this till we reach 7, where 7 * 7 = 49, which is greater than 37. Now we will just return the previous answer which we stored ie 6 * 6 = 36 & this will be the final answer. So we can say that to find the sqrt(n), we need to do sqrt(n) iterations & hence the time complexity is equal to O(sqrt(n)). Space complexity will be O(1).
	// Now assume that we want to further optimize time complexity for this solution. Let's see if we can apply Binary Search to solve this problem which is having less time complexity ie O(log(n)). In order to identify if this is a searching problem, we need to first see if we have a search space. In this case the search space is from 1 - n, so this is a searching problem. Now let's see if we can apply Binary Search. In order to apply Binary Search, we need to have a target which is sqrt(n) in this case & we need to have a condition using which we can discard left or right part once we reach the middle element. In this case once we reach the middle element, we will calculate (mid*mid) & see if it is > n. If it is > n, then we will simply discard the entire right part of the search space & go to left side. ie low=mid-1. Now if mid*mid  < n, then it can be an answer, so we will store mid in answer variable & go to right as we will probably able to find a  better answer in the right side. Now we will also check if mid * mid = n, if it is, then we will simply return mid or we will wait for the BS loop to complete & will return the value of ans.
	// In this case one thing to note is that we are not given any i/p array. Here the array is defined by us just by setting low & high values ie low=1 & high = n. So we don't actually need an i/p array to be given for applying the BS. If we can define a search space ourself, for those problems as well we can apply BS.
	low := 1
	high := n
	var ans int
	for low <= high {
		mid := (low + high) / 2
		// Case 1: if square(mid) == n, then mid is the squae root of n
		if (mid * mid) == n {
			return mid
			// Here we can discard the entire right search space & go to left side
		} else if (mid * mid) > n {
			high = mid - 1
			// Here mid can be the sqrt(n), but we are not sure. So we update the ans with mid & search for a better answer in the right side
		} else {
			ans = mid
			low = mid + 1
		}
	}
	return ans
}

func FindQubeRoot(n int) int {
	// We will appy the same logic as sqaure root
	low := 1
	high := n
	var ans int
	for low <= high {
		mid := (low + high) / 2
		if mid*mid*mid == n {
			return mid
		} else if mid*mid*mid > n {
			high = mid - 1
		} else {
			ans = mid
			low = mid + 1
		}
	}
	return ans
}

func FindMaxSumOfSubArray(k int, ar []int) int {
	// Given an array of positive integers ar & k, return the max sum of sub array of lenght k
	// A sub array is formed by consecutive elements of a given array
	// Eg: ar = [1,2,3,4,5] & k=3
	// Possible Sub Arrays are [1,2,3], [2,3,4], [3,4,5]
	// Assume that sum of first sub array is sum. Then the sum of remaining sub arrays can be found from this initial sum using following formula
	// Sum of next array = (current sum) + (last element in next sub array) - (first element of current subarray)
	// Ref: https://www.geeksforgeeks.org/find-maximum-minimum-sum-subarray-size-k/

	n := len(ar)

	// Find the sum of the first sub array by iterating through each element in it
	var sum, maxSum int
	for i := 0; i < k; i++ {
		sum = sum + ar[i]
		maxSum = sum
	}

	// Run a loop from 'k' to 'n' & find the sum of reamining sub arrays by adding the last element of current sub array & removing first element of last sub array

	for i := k; i < n; i++ {
		sum = sum + ar[i] - ar[i-k] // Remove first element of previous array & add last eleement of current array
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func FindMaxKOfMaxSubArray(b int, ar []int) int {
	//  We are given an array of length 7 as shown in the above diagram. Now let's take k=1, initially. Now each individual element in the array forms a sub array of length 1. So maximum sum of these individual arrays = 7.
	// Also 7 <= 20. This means that k=1, can be an answer till we find a better solution. So we set ans=1. Now let's take k=2. Now maximum sum = 10, which is <= 20.
	// So k=2, can be an answer & we update ans=2. Now let's make k=3. Now max sum = 16 which is <20. This means that k=3 can be an answer & we update the answer with k=3. Now let's take k=4, now max sum = 20. Also 20 <= 20, so k=4 can be an answer. So we update 'k' with ans=4. Now let's make k=5, now max sum becomes 25, which is > 20 & violate the condition. So we will return k=4, as the answer for the above question. Now we can say that k>4 cannot be our answer as when k=5, max(sum) > b. This essentially means that if you add anything to max(sum), when k=5, it will exceed 20. When k=6, we add one more element to array for addition. So when k>=5, max(sum) will be definitely > 20. So we don't need to check any k after 5.
	// Time Complexity [O(log(n))]
	// In this example target is 'k', which can go from 1 - n. So the search space is between 1 to n

	low := 1
	high := len(ar)
	ans := 0
	for low <= high {
		mid := (low + high) / 2
		maxSum := FindMaxSumOfSubArray(mid, ar)
		if maxSum <= b {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

// Check function to check if we can complete the given list of tasks in t time

func Check(k, t int, ar []int) bool {
	// ar = [1, 2, 3, 4, 5]
	// k = 2
	// t = 5
	// To check if the all the tasks given in ar can be completed in 't', we need to keep assigning the tasks to each worker starting from index 0, holding the condition total time of all the tasks assigned to a given worker shouldn't exeed time
	// To do this we run a loop, adding the time of each tasks to sum variable & if sum exceeds t. If sum exceeds 't', we increment w by 1, where w indicate workers. Now within the for loop at any point if we find w > k, ie total workers needed to complete
	// the tasks exceeds the available workers, we exit the loop and return false. Otherwise we rerturn true
	// Time Complexity: O(n)
	// Space Complexity: O(1)
	w := 1
	sum := 0
	n := len(ar)
	for i := 0; i < n; i++ {
		sum = ar[i] + sum
		if sum > t {
			w++
			sum = ar[i]
		}
		if w > k {
			return false
		}
	}
	return true
}

func FindMinTimeForCompletingTasks(k int, ar []int) int {
	// Q: Given 'n' tasks, 'k' workers & the time taken for each task, find the minimum time in which we can complete all the tasks?.
	// Notes:
	// A single worker can only do continuous set of tasks from the task array.
	// All workers start their assigned tasks at the same time.
	// Each task can be assigned only to a single worker.

	// For this to become a searching problem, we need to have a search space where we can definitely find the answer
	// In this case the search space is [max(ar), sum(ar)], ie between max time in the given array to sum of all elements in the given array of tasks
	// This is because, we need to assign each task to atleast one worker & hence the maximum time in the given list is the lowest time in which we can complete all the tasks
	// Eg: ar = [1,1,100], k=2. Now let's assume we assign first two tasks to w1 & last task to w2. Even now w2 will take atleast 100 min to complete the task & hence time taken
	// to complete all the tasks will be max(ar).
	// Similarly consider the case where ar = [1,1,100] & k=1. Now the worker 1 will need to do all tasks alone & he will take sum of time taken by all tasks to complete all tasks

	// So we can conclude that this is a searching problem. Now to see if we can apply Binary Search, we need to see if we can discard left or right side of the array once we reach the middle
	// In this case we make use of following propery of the time to discard left or right part of the array
	// Propery: If we can complete a task in 'n' minutes, we can definitely complete it in n+1, n+2 , n+5 etc times
	// Eg: If we can read a book in 10 minutes, we can definitely read it in say 12, 15, 20 minutes. If we get some more time, probably we will take some breaks & complete reading it.
	// Similary if we cannot complete  a task in 'n' minute, we will definitely not able to complete it in < n minutes
	// Now using this propery we can discard left / right side of the array & apply binary search.

	// We will calculate the mid element in the search space & call check() function to see if we can complete this task in mid time
	// If true, we will disacrd the entire right part & go to left side for finding a better time. Mid can be an answer in this case & hence we update the ans variable with the mid & set high=mid-1
	// Now if we get false from check() function, we won't be able to complete all the tasks in 'mid' time & hence we can discard the entire left part & go to right side for finding a better time
	
	// Time Complexity = Time Complexity Of Check Function * Time Complexity Of Binary Search
	// Time Complexity Of Check Function = O(n), Space = O(1)
	// Time Complexity Of BS = log(x), where x = sum(ar) - max(ar)
	// Space Complexity Of BS = O(1)
	// ie Total Time Complexity = O(n) * O(log(x)) = O(nlog(x)), where x = sum(ar) - max(ar)
	// Total Space Complexity = O(1)

	// Anonimous function to calculate MaxValue in the slice
	maxValue := func(ar []int) int {
		max := 0
		for i := 0; i < len(ar); i++ {
			if ar[i] > max {
				max = ar[i]
			}
		}
		return max
	}

	// Anonimous function to return Sum of all elements in the given slice
	sumValue := func(ar []int) int {
		sum := 0
		for i := 0; i < len(ar); i++ {
			sum = sum + ar[i]
		}
		return sum
	}

	low := maxValue(ar)
	high := sumValue(ar)
	ans := math.MinInt

	// Binary Search to find minimum time
	for low <= high {
		mid := (low + high) / 2
		if Check(k, mid, ar) {
			high = mid - 1
			ans = mid
		} else {
			low = mid + 1
		}
	}

	return ans

}

func Check1(n, d int, ar []int) bool {
	// We need to check if it is possible to allocare n, cows in such a way that we will be able to keep atleast a distance 'd' between each cow
	// d is the distance which we need to check if possible to allocate
	// n, number of cows
	// ar = [1,2,3,4,5]
	// d = 3
	l := len(ar)
	// c, represent number of cows allocated & keep increasing this value every time we allocate a new cow
	c := 1
	lastCowPosition := ar[0]
	for i := 1; i < l; i++ {
		if ar[i]-lastCowPosition >= d {
			c++
			lastCowPosition = ar[i]
		}
		if c == n {
			return true
		}
	}
	return false
}

func FindMaxDistanceBetweenCows(m, n int, ar []int) int {
        // Q: You are given 'n' cows & 'm' stalls. All 'm' stalls are on 'x-axis' at different locations. Place all 'n' cows in these stalls in such a way that, minimum distance between any two cows in maximised?

	//Notes: 
	//1. You can place only one cow in a given stall.
	//2. All cows needs to placed in one of the stalls.
	// ar = [1,2,3,4,5]
	// n = 2
	// Here we have a search space betwen [min(space b/w adjacent stalls), (ar[n]-ar[0])]& target max distance we can achieve between cows, so this is a searching problem
	// Now to apply BST, we should be able to discrd either left or right side after calculating the mid
	// Here we can say that we cannot allocate all cows with a distance 'd', then we can say that any distance > d is also not possible & hence we can discard entire right side of d
	// Similarly if it is possible to allocate cows with a distance 'd', we can definitely say that it is 100% possible to have a distance < d. So we will be able to find a better answer possibly in the right side as we want to maximize the distances. So weill update the ans = mid & search in the right side

	n1 := len(ar) // 5
	minDistance := func(ar []int) int {
		// Minimum distance between stalls
		ans := math.MaxInt
		for i := 0; i < n1-1; i++ {
			d := ar[i+1] - ar[i]
			if d < ans {
				ans = d
			}
		}
		return ans
	}
	low := minDistance(ar)
	high := ar[n1-1] - ar[0]

	ans := math.MinInt

	for low <= high {
		mid := (high + low) / 2
		if Check1(n, mid, ar) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
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
        r2 := FindElementAfterKRotationWithoutKGiven(ar1, 6)
	infoLogger.Print(r2)
	r3 := FindSqrt(81)
	infoLogger.Print(r3)
        r4 := FindQubeRoot(27)
	infoLogger.Print(r4)
        ar1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r5 := FindMaxKOfMaxSubArray(19, ar1)
	infoLogger.Print(r5)
	ar2 := []int{3, 5, 1, 7, 8, 2, 5, 3, 10, 1, 4, 7, 5, 4, 6}
	minTime := FindMinTimeForCompletingTasks(4, ar2)
	infoLogger.Print(minTime)
	ar3 := []int{3, 8, 12, 18, 25, 30, 35, 41, 49}
	maxDistance := FindMaxDistanceBetweenCows(9, 4, ar3)
	infoLogger.Print(maxDistance)
}
