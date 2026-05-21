func maximumTotalDamage(spells []int) int64 {
    sort.Ints(spells)
    prevVal := [3]int{math.MinInt32, math.MinInt32, math.MinInt32}
    prevDp := [3]int64{0, 0, 0}
    var maxDamage int64 = 0
    i := 0
    for i < len(spells) {
        j := i+1
        var currAcc int64 = int64(spells[i])
        for j < len(spells) && spells[j] == spells[i] {
            currAcc += int64(spells[j])
            j++
        }
        for k := 2; k >= 0; k-- {
            if spells[i]-2 > prevVal[k] {
                maxDamage = max(maxDamage, currAcc+prevDp[k]);
                break
            }
        }
        prevVal[0], prevVal[1], prevVal[2] = prevVal[1], prevVal[2], spells[i]
        prevDp[0], prevDp[1], prevDp[2] = prevDp[1], prevDp[2], maxDamage
        i = j;
    }
    return maxDamage
}