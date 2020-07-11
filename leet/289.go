//https://leetcode-cn.com/problems/game-of-life/
package main

/*
定义 0 为 死亡
定义 1 为 活着
定义 3 为 待活(0 -> 3)
定义 4 为 待死(1 -> 4)

x 表示横坐标
y 表示纵坐标

(x+1, y)   向右
(x, y + 1) 向上
(x - 1, y) 向左
(x, y - 1) 向下
(x + 1, y + 1) 右上
(x + 1, y - 1) 右下
(x - 1, y + 1) 左上
(x - 1, y - 1) 左下
*/
func main (){
    var a = [][]int{{1, 1}, {0, 0}}
    gameOfLife(a)
}

func gameOfLife(board [][]int)  {
    totalX := len(board) - 1
    totalY := len(board[0]) - 1

    for i := 0; i <= totalX; i++ {
        for j := 0; j <= totalY; j++ {
            changeState(i, j, totalX, totalY, board)
        }
    }

    for i := 0; i <= totalX; i++ {
        for j := 0; j <= totalY; j++ {
            if board[i][j] == 3{
                board[i][j] = 1
            }
            if board[i][j] == 4{
                board[i][j] = 0
            }
        }
    }
}


func changeState(x, y, totalX, totalY int, chess [][]int){
    live, died := 0, 0
    if x + 1 <= totalX && y <= totalY {
        if chess[x+1][y] == 1 || chess[x+1][y] == 4{
            live++
        } else {
            died++
        }
    }
    if x <= totalX && y + 1 <= totalY {
        if chess[x][y + 1] == 1 || chess[x][y + 1] == 4{
            live++
        } else {
            died++
        }
    }

    if x - 1 >= 0 && y <= totalY {
        if chess[x - 1][y] == 1 || chess[x - 1][y] == 4{
            live++
        } else {
            died++
        }
    }

    if x <= totalX && y - 1 >= 0 {
        if chess[x][y - 1] == 1 || chess[x][y - 1] == 4{
            live++
        } else {
            died++
        }
    }

    if x + 1 >= 0 && x + 1 <= totalX && y + 1 >= 0 && y + 1 <= totalY {
        if chess[x + 1][y + 1] == 1 || chess[x + 1][y + 1] == 4 {
            live++
        } else {
            died++
        }
    }
    if x + 1 >= 0 && x + 1 <= totalX && y - 1 >= 0 && y - 1 <= totalY {
        if chess[x + 1][y - 1] == 1 || chess[x + 1][y - 1] == 4 {
            live++
        } else {
            died++
        }
    }
    if x - 1 >= 0 && x - 1 <= totalX && y + 1 >= 0 && y + 1 <= totalY {
        if chess[x - 1][y + 1] == 1 || chess[x - 1][y + 1] == 4{
            live++
        } else {
            died++
        }
    }
    if x - 1 >= 0 && x - 1 <= totalX && y - 1 >= 0 && y - 1 <= totalY {
        if chess[x - 1][y - 1] == 1 || chess[x - 1][y - 1] == 4{
            live++
        } else {
            died++
        }
    }

    if chess[x][y] == 1 {
        if live < 2 {
            chess[x][y] = 4
        }
        if live > 3 {
            chess[x][y] = 4
        }
    } else {
        if live == 3 {
            chess[x][y] = 3
        }
    }
}

