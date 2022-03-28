package bfs

import "container/list"

// 找到字典中单词转换的最短的具体路径
// bfs+dfs 综合
func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := map[string]struct{}{}
	for _, v := range wordList {
		dict[v] = struct{}{}
	}
	if _, ok := dict[endWord]; !ok {
		return nil
	}
	// 记录转换序列
	distance := markDepth(beginWord, endWord, dict)
	res := [][]string{}
	backtrackFindLadders(distance, &res, beginWord, beginWord, endWord, dict, []string{})
	return res
}
func markDepth(beginWord string, endWord string, dict map[string]struct{}) map[string]int {
	queue := list.New()
	queue.PushBack(beginWord)
	visited := map[string]bool{beginWord: true}
	distance := map[string]int{beginWord: 1}
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(string)
		if cur == endWord {
			break
		}
		for _, nextWord := range neighborWords(dict, cur) {
			if visited[nextWord] {
				continue
			}
			visited[nextWord] = true
			distance[nextWord] = distance[cur] + 1
			queue.PushBack(nextWord)
		}
	}
	return distance

}

func backtrackFindLadders(depth map[string]int, res *[][]string,
	cur, beginWord, endWord string,
	dict map[string]struct{}, path []string) {
	path = append(path, cur)
	if cur == endWord {
		*res = append(*res, append([]string{}, path...))
	}
	for _, word := range neighborWords(dict, cur) {
		// 该邻接字符串不是由 cur 生成
		if depth[cur]+1 != depth[word] {
			continue
		}
		backtrackFindLadders(depth, res, cur, beginWord, endWord, dict, path)
		// path = path[:len(path)-1]
	}

}

// 找到存在 dict 中的 word 可变换的 string 组
func neighborWords(dict map[string]struct{}, word string) []string {
	nextWordList := []string{}
	for i := 0; i < len(word); i++ {
		for char := 'a'; char <= 'z'; char++ {
			newWord := word[:i] + string(char) + word[i+1:]
			if _, ok := dict[newWord]; ok {
				nextWordList = append(nextWordList, newWord)
			}
		}
	}
	return nextWordList
}
