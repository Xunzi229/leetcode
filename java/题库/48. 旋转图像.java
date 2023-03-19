// https://leetcode.cn/problems/rotate-image/

/*
1 2 3
4 5 6
7 8 9
*/

class Solution {
    public void rotate(int[][] matrix) {
        if(matrix.length == 0 || matrix.length != matrix[0].length) {
            return;
        }

        int maxRowLen = matrix.length;
        /*
         * 先纵向对折
         * 3 2 1
         * 6 5 4
         * 9 8 7
         */
        for (int i = 0; i < matrix.length; ++i) {
            int[] row = matrix[i];
            int maxColumnLen = row.length;
            for (int y = 0; y < row.length / 2; ++y) {
                int t = row[y];
                // ==> 指向新的坐标
                matrix[i][y] = matrix[i][maxColumnLen - y - 1];
                matrix[i][maxColumnLen - y - 1] = t;
            }
        }

        /*
         * 再斜对折
         * 7 4 1
         * 8 5 2
         * 9 6 3
         *
         */
        for (int i = 0; i < matrix.length; i++) {
            if (maxRowLen - i - 1 <= 0) {
                continue;
            }

            int[] row = matrix[i];
            int maxColumnLen = row.length;

            for (int y = 0; y < row.length; y++) {
                if (maxColumnLen - i - y - 1 <= 0) {
                    continue;
                }
                int t = row[y];
                // ==> 指向新的坐标
                matrix[i][y] = matrix[maxRowLen - y - 1][maxColumnLen - i - 1];
                matrix[maxRowLen - y - 1][maxColumnLen - i - 1] = t;
            }
        }
    }
}







