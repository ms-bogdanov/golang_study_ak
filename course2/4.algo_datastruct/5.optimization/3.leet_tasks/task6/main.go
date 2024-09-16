package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max, idx := maxVal(nums)

	root := &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}

	return root
}

func maxVal(nums []int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}

	max, idx := nums[0], 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}

	return max, idx
}
