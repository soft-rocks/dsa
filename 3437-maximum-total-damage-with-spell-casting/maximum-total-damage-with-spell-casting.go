type Spell struct {
    totalDamage int64
    power       int
}

func maximumTotalDamage(p []int) int64 {
    sort.Ints(p)
    dp := make([]Spell, 5)
    cnt := 0
    
    for i := 0; i < len(p); i++ {
        cnt++
        // Merge spells with same power
        if i != len(p)-1 && p[i] == p[i+1] {
            continue
        }
        
        // Calculate damage if we select current spell
        curr := int64(p[i]) * int64(cnt)
        cnt = 0
        
        // Find optimal position in window
        candidate := 5
        for j := 4; j >= 0; j-- {
            if p[i]-dp[j].power <= 2 {
                continue
            }
            if j == 4 || candidate == 5 {
                candidate = j
            } else if dp[j+1].power-dp[j].power > 2 && p[i]-dp[j+1].power > 2 {
                break
            } else {
                if dp[j].totalDamage > dp[candidate].totalDamage {
                    candidate = j
                }
            }
        }
        
        var damage int64
        if candidate != 5 {
            damage = dp[candidate].totalDamage
        }
        
        // Update sliding window
        dp = append(dp[1:], Spell{
            totalDamage: damage + curr,
            power:       p[i],
        })
    }
    
    return max(dp[2].totalDamage, dp[3].totalDamage, dp[4].totalDamage)
}