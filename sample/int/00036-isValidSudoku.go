package int

// 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

// 上图是一个部分填充的有效的数独。

// 数独部分空格内已填入了数字，空白格用 '.' 表示。

// 示例 1:

// 输入:
// [
//   ["5","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// 输出: true
// 示例 2:

// 输入:
// [
//   ["8","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// 输出: false
// 解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
//      但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。

// //isValidSudoku
// func isValidSudoku(board [][]byte) bool {

// 	// map1r := make(map[byte]bool)
// 	// map2r := make(map[byte]bool)
// 	// map3r := make(map[byte]bool)
// 	// map4r := make(map[byte]bool)
// 	// map5r := make(map[byte]bool)
// 	// map6r := make(map[byte]bool)
// 	// map7r := make(map[byte]bool)
// 	// map8r := make(map[byte]bool)
// 	// map9r := make(map[byte]bool)
// 	// map1c := make(map[byte]bool)
// 	// map2c := make(map[byte]bool)
// 	// map3c := make(map[byte]bool)
// 	// map4c := make(map[byte]bool)
// 	// map5c := make(map[byte]bool)
// 	// map6c := make(map[byte]bool)
// 	// map7c := make(map[byte]bool)
// 	// map8c := make(map[byte]bool)
// 	// map9c := make(map[byte]bool)

// 	//1-3r
// 	map1 := make(map[byte]bool)
// 	for r := 0; r < 3; r++ {
// 		for c := 0; c < 3; c++ {
// 			if _, ok := map1[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map1[board[r][c]] = true
// 		}
// 	}
// 	map2 := make(map[byte]bool)
// 	for r := 0; r < 3; r++ {
// 		for c := 3; c < 6; c++ {
// 			if _, ok := map2[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map2[board[r][c]] = true
// 		}
// 	}
// 	map3 := make(map[byte]bool)
// 	for r := 0; r < 3; r++ {
// 		for c := 6; c < 9; c++ {
// 			if _, ok := map3[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map3[board[r][c]] = true
// 		}
// 	}
// 	//4-6r
// 	map4 := make(map[byte]bool)
// 	for r := 3; r < 6; r++ {
// 		for c := 0; c < 3; c++ {
// 			if _, ok := map4[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map4[board[r][c]] = true
// 		}
// 	}
// 	map5 := make(map[byte]bool)
// 	for r := 3; r < 6; r++ {
// 		for c := 3; c < 6; c++ {
// 			if _, ok := map5[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map5[board[r][c]] = true
// 		}
// 	}
// 	map6 := make(map[byte]bool)
// 	for r := 3; r < 6; r++ {
// 		for c := 6; c < 9; c++ {
// 			if _, ok := map6[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map6[board[r][c]] = true
// 		}
// 	}
// 	//6-9r
// 	map7 := make(map[byte]bool)
// 	for r := 6; r < 9; r++ {
// 		for c := 0; c < 3; c++ {
// 			if _, ok := map7[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map7[board[r][c]] = true
// 		}
// 	}
// 	map8 := make(map[byte]bool)
// 	for r := 6; r < 9; r++ {
// 		for c := 3; c < 6; c++ {
// 			if _, ok := map8[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map8[board[r][c]] = true
// 		}
// 	}
// 	map9 := make(map[byte]bool)
// 	for r := 6; r < 9; r++ {
// 		for c := 6; c < 9; c++ {
// 			if _, ok := map9[board[r][c]]; ok && string(board[r][c]) != "." {
// 				return false
// 			}
// 			map9[board[r][c]] = true
// 		}
// 	}

// 	return true
// }

func isValidSudoku(board [][]byte) bool {
	//row
	tmpRow := map[byte]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if _, ok := tmpRow[board[i][j]]; ok && string(board[i][j]) != "." {
				return false
			}
			tmpRow[board[i][j]] = true
		}
	}
	//column
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if _, ok := tmpRow[board[j][i]]; ok && string(board[j][i]) != "." {
				return false
			}
			tmpRow[board[j][i]] = true
		}
	}
	return true
}
