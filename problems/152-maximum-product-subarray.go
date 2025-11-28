package problems

func maxProduct(nums []int) int {
    
    maxProduct := nums[0]
    minProduct := nums[0]
    result := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            minProduct, maxProduct = maxProduct, minProduct
        }
        minProduct = min(nums[i], nums[i]*minProduct)
        maxProduct = max(nums[i], nums[i]*maxProduct)
        result = max(maxProduct, result)
    }
    return result
}
