# s = "abcde"
# n = len(s)

# for length in range(2, n + 1):
#     print(f"\nlength = {length}")
#     for i in range(n - length + 1):
#         j = i + length - 1
#         print(f"i={i}, j={j}, s[{i}:{j + 1}] = '{s[i:j + 1]}'")

class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        if n < 2:
            return s # 1文字か空文字はそのまま回文にする
        
        # dp[i][j] = s[i..j]が回文かどうか
        dp = [[False] * n for _ in range(n)]
        start = 0
        max_len = 1

        # 1文字はすべて回文
        for i in range(n):
            dp[i][i] = True

        # 長さ2以上の部分文字列をチェックする
        for length in range(2, n + 1):
            for i in range(n - length + 1):
                j = i + length - 1

                if s[i] == s[j]:
                    if length == 2:
                        dp[i][j] = True
                    else:
                        dp[i][j] = dp[i+1][j-1]
                    
                    # 最長なら更新する
                    if dp[i][j] and length > max_len:
                        start = i
                        max_len = length
                
                else:
                    dp[i][j] = False
        
        return s[start:start + max_len]


