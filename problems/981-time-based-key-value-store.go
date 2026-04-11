package problems

import "sort"

type TimeValue struct {
    Val string
    Timestamp int
}

type TimeMap struct {
    Data map[string][]TimeValue
}

func ConstructTimeMap() TimeMap {
    return TimeMap{Data: make(map[string][]TimeValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    tsv := TimeValue{Val: value, Timestamp: timestamp}
    this.Data[key] = append(this.Data[key], tsv)
}

func (this *TimeMap) Get(key string, timestamp int) string {
    values, ok := this.Data[key]
    if !ok || len(values) == 0 {
        return ""
    }
    // Binary search to find the first timestamp GREATER than the target

    // Option 1: beats 80.4% - but locally performed better - see the test file with benchmark comment
    i := sort.Search(len(values), func(i int) bool {
        return values[i].Timestamp > timestamp
    })

    // If i is 0, it means even the first timestamp is greater than target
    if i == 0 {
        return ""
    }
    // Otherwise, the "largest timestamp_prev" is at index i-1
    return values[i-1].Val

    // Option 2: beats 87.54%
    // low := 0
    // high := len(values) - 1
    // ans := -1

    // for low <= high {
    //     mid := low + (high-low)/2

    //     if values[mid].Timestamp <= timestamp {
    //         // This is a potential candidate, but there might be a
    //         // larger timestamp later in the slice that is still <= target.
    //         ans = mid
    //         low = mid + 1
    //     } else {
    //         // This timestamp is too large, look in the left half.
    //         high = mid - 1
    //     }
    // }

    // if ans == -1 {
    //     return ""
    // }
    // return values[ans].Val

}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
