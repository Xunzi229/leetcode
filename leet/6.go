//https://leetcode-cn.com/problems/zigzag-conversion/solution/z-zi-bian-huan-you-hua-by-yujwhe-ta-de-ba-kuai-fu-/
package main

import "fmt"

func main (){
    str := "heltfchqssrwqgwanggkjlsownsdpoowubszfzratjwlpuldarnmehcbvuemiulcxdedcxfygbjyyxbyqqmvxoyukchszuxwxdbbagzjklhiikiyavvzltwwyfqxzpvwszxvfzerknbuxkszhoaujwqhbjecycyrbyoizucjhddgpxfynftxelehulktnkkqkaajucsdgxjvvoukvphzamjvxtomfacqaezwhuzntkkqagbvxkxywgtvbjjijnylsajzwioruaiujlrgvoguwzrzkbivogggiphgzvytygnhtfnovwkuvctidbdrkkaubhbddzwbhmkatzqqvbktdgbgjezvqzqshtxmutpbhzdcyvvwwhpbnqjxujunkmhtfehzzwchxhlydiubqjddbmcxxzkilrdrvlsvjvehcrfhabjqkmvnaykyxviimnbkyufirlpvcwdcxmsjaowaogandkxsybcwvjgouxjytobscvdclbfzkfonqmfqpjmksvaoslnoaqgelmhxnmyxtnllbsbqcocwjendparrsywdkfazrbxmoiyrczjgplfypseguvymvuphzshsteejoccsclzrwesnyytsttgppvwqpfikjpvztxsxirrgxlvvjpnckttaqqqivbshsogllylwrccopylypaabvwbomuwjxqspezcszpqtrsjgsvgjxhltdohrifchvvyawbuxqkskecszzzkyixrnmagwfiebfcdbfxbyjtipxcoybzxjyowkrcjwnpxstawbzxzisjysloqnpnyoevavzjrmarhutdvtcwdwfdoqsffhuexazyvajpnkiugbzdwdzazedowxvchrgeshephogwaosiqtlmwmowssmopjswayduhhkrxqnzhijxbulyiawauirjtjitk"
    fmt.Println(len(str))
    fmt.Println(convert(str, 1))

}


/*
当进行Z 型移动时 ， 每次恢复到下一个 列开始的时候
1行 需要 1 列 1 个元素恢复到下一个 列开始的时候
2行 需要 1 列 2 个元素恢复到下一列开始的时候
3行 需要 2 列 4 个元素恢复到下一列开始的时候
4行 需要 3 列 6 个元素恢复到下一列开始的时候
5行 需要 4 列 8 个元素恢复到下一列开始的时候

提交后发现这不是一个好的方法， 如何优化呢
*/

func convert(s string, numRows int) string {
  strLen := len(s)
  if numRows <= 1 || strLen <= 1 {return s}

  var numColumns = 0
  numColumns = strLen / 2 + 1
  var chess = make([][]string, numRows)

  for  i := 0; i < numRows; i++ {
      for j := 0; j < numColumns; j++ {
          chess[i] = append(chess[i], "0")
      }
  }
  down(0, s, 0, 0, numRows, numColumns, chess)

  var newStr string
  for  i := 0; i < numRows; i++ {
      for j := 0; j < numColumns; j++ {
          if chess[i][j] != "0" {
              newStr += chess[i][j]
          }
      }
   }
   return  newStr
}

/*
site  是str的位置
str   string
x y   当前的棋盘坐标
numRows 当前棋盘的 行数
numColumns 当前棋盘的 列长度
chess 棋盘
*/
/*
向下的时候纵坐标不变
*/

func down(site int, str string, x, y, numRows, numColumns int, chess [][]string){
    for {
        // 如果遍历的字符串到最后， 就return
        if site + 1 > len(str){ return }
        // 当前的 位置改变
        chess[x][y] = string(str[site])

        if x + 1 == numRows {
            up(site + 1, str, x - 1, y + 1, numRows, numColumns, chess)
            return
        }
        site++; x++
    }
}

/*
向上的时候横纵坐标同时加1
*/
func up(site int, str string, x, y, numRows, numColumns int, chess [][]string){
    for {
        if x == 0 {
            down(site, str, x, y, numRows, numColumns, chess)
            return
        }
        if site + 1 > len(str) { return }
        chess[x][y] = string(str[site])
        x--; y++; site++
    }
}