package main

import (
	"fmt"
	mat "learning/linearAlgebra/matrices"
)

func main() {
	old := mat.MakeMat(3, 3)
	old.RandFill()

	new := mat.MakeMat(2, 2)

	var nums []int

	for i := range old.Mat {
		nums = append(nums, i)
	}

	fmt.Println("Nums:")
	fmt.Println(nums)

	for i := range nums {
		var nums1 []int
		for j := range nums {
			if nums[j] != i {
				nums1 = append(nums1, j)
			}
		}
		for j := range nums {
			var nums2 []int
			for k := range nums {
				if nums[k] != j {
					nums2 = append(nums2, k)
				}
			}
			fmt.Println(nums1, nums2)

			for newX, oldX := range nums1 {
				for newY, oldY := range nums2 {
					new.Mat[newX][newY] = old.Mat[oldX][oldY]
				}
			}
			old.PrintMat()
			new.PrintMat()
		}

	}

}
