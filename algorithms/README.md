

## 01
给定一个 N*N 的矩阵，例如：
1  2  3  4  5
11 12 13 14 15
21 22 23 24 25
31 32 33 34 34
66 63 65 67 62
给一个长度为 M 的数组，M > N
2 12 13 23 33 65,67
将数组映射到矩阵上，判断数组是否是相邻数组？(可以一笔画)
https://gitee.com/zmzhou-star/learnotes/raw/master/src/main/java/com/github/zmzhoustar/AdjacentArray.java

## 02
请根据输入两个字符串，求它们的最长公共字串，输出最长公共字串的长度。
找不到字串输出：0
例如：
输入：ab1abc2ddd ab1abc 最长公共字串为ab1abc，输出结果6
输入：abcdef abc1def 最长公共字串为abc，输出结果3
https://gitee.com/zmzhou-star/learnotes/raw/master/src/main/java/com/github/zmzhoustar/LongestCommonStr.java

## 03
求最大面积
假如给航天器装太阳能电池板（长方形），先在航天器一侧装上长度不一的柱子，每个柱子之间的距离为 1，
电池板必须装在相邻的柱子上
例如：一组柱子长度为 {10,9,8,7,6,5,4,3,2,1}
求得最大面积：25

https://gitee.com/zmzhou-star/learnotes/raw/master/src/main/java/com/github/zmzhoustar/MaxArea.java

## 04
Leetcode算法题：695. 岛屿的最大面积
给定一个包含了一些 0 和 1 的非空二维数组 grid 。
一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
算法：
BFS(广度优先遍历) / DFS(深度优先遍历)

https://gitee.com/zmzhou-star/learnotes/raw/master/src/main/java/com/github/zmzhoustar/MaxAreaOfIsland.java