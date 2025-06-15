package main

import "fmt"

type Item struct {
	Name string
	Price int
}

func sack(items []Item, budget int) ([]Item, int) {
	n := len(items)
	// dp[i][j] := i番目までの品物で、予算j以下の時の最大合計価格
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, budget+1)
	}

	// 選択したかどうかの情報を保持する
	selected := make([][]bool, n+1)
	for i := range selected {
		selected[i] = make([]bool, budget+1)
	}

	// DP計算
	for i := 1; i <= n; i++ {
		itemPrice := items[i-1].Price
		for j := 0; j <= budget; j++ {
			if itemPrice <= j {
				if dp[i-1][j-itemPrice]+itemPrice > dp[i-1][j] {
					dp[i][j] = dp[i-1][j-itemPrice] + itemPrice
					selected[i][j] = true
				} else {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	res := []Item{}
	w := budget
	for i := n; i > 0; i-- {
		if selected[i][w] {
			res = append(res, items[i-1])
			w -= items[i-1].Price
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	
	return res, dp[n][budget]
}

func main() {
	items := []Item{
		{"牛丼 並盛", 400},
        {"味噌汁", 100},
        {"たまご", 80},
        {"サラダ", 150},
        {"とん汁", 190},
        {"お新香", 130},
		{"ライト", 300},
	}

	budget := 1000

	combo, total := sack(items, budget)
	fmt.Printf("予算%dで最適な組み合わせです。合計%d円\n", budget, total)
	for _, item := range combo {
		fmt.Printf(" %s: %d円\n", item.Name, item.Price)
	}
}