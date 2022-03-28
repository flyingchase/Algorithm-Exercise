package bfs

// 单词接龙
/*
将 dict 中每个单词视为结点，逐个入队
将逐个队头中字符串中逐个字符 26 变化，存在 dict 则入队
level 为层数即为变化次数
*/
func LadderLength(beginWord string, endWord string, wordList []string) int {
	dict := map[string]struct{}{}
	for _, v := range wordList {
		dict[v] = struct{}{}
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	queue := []string{beginWord}
	visited := map[string]bool{}
	visited[beginWord] = true
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node == endWord {
				return level
			}
			for index := 0; index < len(node); index++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					// 改变 index 位置的字符
					newWord := node[:index] + string(ch) + node[index+1:]
					if _, ok := dict[newWord]; ok {
						if !visited[newWord] {
							queue = append(queue, newWord)
							if newWord == endWord {
								return level + 1
							}
							// 访问过则标记
							visited[newWord] = true
						}
					}
				}
			}
		}
		level++
	}
	return 0
}

// 起点终点均确定，使用双向 bfs
func ladderLengthTwoBFS(beginWord string, endWord string, wordList []string) int {
	dict := map[string]struct{}{}
	for _, num := range wordList {
		dict[num] = struct{}{}
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	leftQueue, rightQueue := []string{beginWord}, []string{endWord}
	leftVisited, rightVisited := map[string]struct{}{}, map[string]struct{}{}
	leftVisited[beginWord], rightVisited[endWord] = struct{}{}, struct{}{}
	level := 1
	for len(leftQueue) != 0 && len(rightQueue) != 0 {
		// 每次走短的一边,默认左侧
		if len(leftQueue) > len(rightQueue) {
			leftQueue, rightQueue = rightQueue, leftQueue
			leftVisited, rightVisited = rightVisited, leftVisited
		}
		size := len(leftQueue)
		for size > 0 {
			size--
			node := leftQueue[0]
			leftQueue = leftQueue[1:]
			// 左侧遍历在右侧 visited 出现则找到
			if _, ok := rightVisited[node]; ok {
				return level
			}
			for index := 0; index < len(node); index++ {
				for char := 'a'; char <= 'z'; char++ {
					newWord := node[:index] + string(char) + node[index+1:]
					if _, ok := dict[newWord]; ok {
						if _, ok := leftVisited[newWord]; !ok {
							leftQueue = append(leftQueue, newWord)
							leftVisited[newWord] = struct{}{}
						}
					}
				}
			}
		}
		level++
	}
	return 0
}
