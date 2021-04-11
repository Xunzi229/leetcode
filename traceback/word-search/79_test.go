package word_search

import (
	"fmt"
	"testing"
)

//[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
//"ABCB"
func PackageExist() bool {
	//by := [][]byte{
	//	{'A', 'B', 'C', 'E'},
	//	{'S', 'F', 'C', 'S'},
	//	{'A', 'D', 'E', 'E'},
	//}
	//by := [][]byte{
	//	{'a', 'a'},
	//}

	by := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	return exist(by, "ABCCED")
}

func TestPackageExist(t *testing.T) {
	fmt.Println(PackageExist())
}
