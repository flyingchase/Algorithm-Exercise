class Solution:
    # @return a undirected graph node
    def cloneGraph(self, node):
        if not node:
            return None
        head = UndirectedGraphNode(node.label)
        q, visited, dic = [], set([node]), {}
        dic[node] = head
        q.append(node)
        while q:
            cur = q.pop()
            for i in cur.neighbors:
                if i not in dic:
                    dic[i] = UndirectedGraphNode(i.label)
                dic[cur].neighbors.append(dic[i])
                if i not in visited:
                    q.append(i)
                    visited.add(i)
        return head


class UndirectedGraphNode:
    def __init__(self, x):
        self.label = x
        self.neighbors = []


def printGraph(node):
    q, visited = [], set([node])
    q.append(node)
    while q:
        cur = q.pop(0)
        print(" ", cur.label)
        for i in cur.neighbors:
            print(i.label)
            if i not in visited:
                q.append(i)
                visited.add(i)


s = Solution()
one = UndirectedGraphNode(1)
two = UndirectedGraphNode(2)
three = UndirectedGraphNode(3)
four = UndirectedGraphNode(4)

one.neighbors.append(one)
one.neighbors.append(two)
one.neighbors.append(three)

printGraph(s.cloneGraph(one))
