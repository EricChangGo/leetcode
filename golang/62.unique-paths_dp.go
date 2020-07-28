/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {
    graph:= make([][]int, n)
    for i := range graph {
        graph[i] = make([]int, m)
        graph[i][0] = 1
    }
    
    for i := range graph[0] {
        graph[0][i] = 1
    }
    
    for i:=1; i < n; i++{
        for j:=1; j < m; j++ {
             graph[i][j] = graph[i-1][j] + graph[i][j-1]
        }
    }
    return graph[n-1][m-1]
}
// @lc code=end
