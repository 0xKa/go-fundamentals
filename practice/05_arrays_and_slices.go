package practice

import "fmt"

func ArraysAndSlicesEx5() {
	array := [...]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("Array: %v (type %T)\n", array, array)
	fmt.Printf("Slice: %v (type %T)\n", slice, slice)

	nums := []int{10, 20, 30, 40, 50}
	part := nums[1:3]

	fmt.Println("\nLength and capacity:")
	fmt.Printf("nums: %v, len=%d, cap=%d\n", nums, len(nums), cap(nums))
	fmt.Printf("part: %v, len=%d, cap=%d\n", part, len(part), cap(part))

	part[0] = 99

	fmt.Println("\nSlices share their backing array:")
	fmt.Println("nums:", nums)
	fmt.Println("part:", part)

	nums = []int{10, 20, 30, 40, 50}
	part = nums[1:3]
	part = append(part, 99)

	fmt.Println("\nAppend reuses available capacity:")
	fmt.Println("nums:", nums)
	fmt.Println("part:", part)

	nums = []int{10, 20, 30, 40, 50}
	part = nums[1:3]
	part = append(part, 60, 70, 80)

	fmt.Println("\nAppend allocates when capacity is exceeded:")
	fmt.Println("nums:", nums)
	fmt.Println("part:", part)
}
