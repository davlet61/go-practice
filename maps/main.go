package main

import "fmt"

func main() {
	userRoles := make(map[string]string)

	userRoles["Mike"] = "admin"
	userRoles["Sheela"] = "user"
	userRoles["Maleek"] = "editor"
	userRoles["George"] = "guest"

	fmt.Println(userRoles)

	delete(userRoles, "George")
	fmt.Println("after deletion", userRoles)

	emails := map[string]string{"Mike": "mike@jordan.com", "Sheela": "sheela30@me.com"}
	fmt.Println(emails)

	// range
	nums := []int{54, 23, 12, 45, 67, 89, 90, 12, 34, 56, 78, 90}

	for i, num := range nums {
		fmt.Printf("%d -> ID: %d\n", i, num)
	}

	// with no index
	for _, num := range nums {
		fmt.Printf("id => %d\n", num)
	}

	newNums := []int{1, 2, 3, 4, 5}

	// multiply all the numbers in the slice
	result := 1
	for _, num := range newNums {
		result *= num
	}
	fmt.Printf("Multiplied result => %d\n", result)

	// multiply each number by 2
	for i, num := range newNums {
		newNums[i] = num * 2
	}
	fmt.Println("Multiplied by 2 =>", newNums)

	// Range with map
	for key, value := range userRoles {
		fmt.Printf("%s ---> %s\n", key, value)
	}
}
