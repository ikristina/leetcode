package problems

func waysToMakeFair(nums []int) int {

    var totalEven, totalOdd int
    for i, num := range nums {
        if i % 2 == 0 {
            totalEven += num
        } else {
            totalOdd += num
        }
    }

    var leftEven, leftOdd int
    var count int

    for i, num := range nums {
        if i % 2 == 0 { // removing the number from the array
            totalEven -= num
        } else {
            totalOdd -= num
        }

        // after the removal left (visited so far) part stays whatever it was and the rest of number flip odd - even
        if leftEven + totalOdd == leftOdd + totalEven {
            count++
        }

        if i % 2 == 0 { // adding the number to the "visited" leftEven or leftOdd
            leftEven += num
        } else {
            leftOdd += num
        }
    }
    return count
}
