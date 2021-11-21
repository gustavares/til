package main

import (
	"fmt"
)

/*

We have some clickstream data that we gathered on our client's website. Using cookies, we collected snippets of users' anonymized URL histories while they browsed the site. The histories are in chronological order and no URL was visited more than once per person.
Write a function that takes two users' browsing histories as input and returns the longest contiguous sequence of URLs that appears in both.
Sample input:
user0 = ["/start", "/pink", "/register", "/orange", "/red", "a"]
user1 = ["/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"]
user2 = ["a", "/one", "/two"]
user3 = ["/red", "/orange", "/yellow", "/green", "/blue", "/purple", "/white", "/amber", "/HotRodPink", "/BritishRacingGreen"]
user4 = ["/red", "/orange", "/amber", "/random", "/green", "/blue", "/purple", "/white", "/lavender", "/HotRodPink", "/BritishRacingGreen"]
user5 = ["a"]
Sample output:
findContiguousHistory(user0, user1)
   /pink
   /register
   /orange
findContiguousHistory(user1, user2)
   (empty)
findContiguousHistory(user2, user0)
   a
findContiguousHistory(user5, user2)
   a
findContiguousHistory(user3, user4)
   /green
   /blue
   /purple
   /white
findContiguousHistory(user4, user3)
   /green
   /blue
   /purple
   /white
n: length of the first user's browsing history
m: length of the second user's browsing history

*/

func main() {
	user0 := []string{"/start", "/pink", "/register", "/orange", "/red", "a"}
	user1 := []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"}
	user2 := []string{"a", "/one", "/two"}
	user3 := []string{"/red", "/orange", "/yellow", "/green", "/blue", "/purple", "/white", "/amber", "/HotRodPink", "/BritishRacingGreen"}
	user4 := []string{"/red", "/orange", "/amber", "/random", "/green", "/blue", "/purple", "/white", "/lavender", "/HotRodPink", "/BritishRacingGreen"}
	user5 := []string{"a"}

	fmt.Println(findContiguousHistory(user0, user1))
	fmt.Println(findContiguousHistory(user0, user2))
	fmt.Println(findContiguousHistory(user2, user0))
	fmt.Println(findContiguousHistory(user5, user2))
	fmt.Println(findContiguousHistory(user3, user4))
	fmt.Println(findContiguousHistory(user4, user3))
}

func findContiguousHistory(userA []string, userB []string) []string {
	start := 0
	max := 0

	for i := 0; i <= len(userA); i++ {
		for j := 0; j <= len(userB); j++ {
			for k := 0; k+i < len(userA) && k+j < len(userB); k++ {
				if userA[i+k] != userB[j+k] {
					break
				}

				if max < k+1 {
					max = k + 1
					start = i
				}
			}

		}
	}

	return userA[start : start+max]
}
