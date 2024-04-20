from collections import deque

from graph_adjacency_list import GraphAdjList


def graph_bfs(graph: GraphAdjList, start_vet: int) -> list[int]:
    res = []
    # 待查詢表
    queue = deque()
    queue.append(start_vet)
    # 已看過表
    visited: set[int] = set()
    visited.add(start_vet)
    # 只要還有需要查看的，就一直不斷查看
    while len(queue) > 0:
        # 從待查詢表中的取得要查看的元素
        vertex = queue.popleft()
        res.append(vertex)
        # 先找到頂點表的位置
        idx = graph.adj_list.index(vertex)
        # 查看該頂點的邊，將邊所連接的頂點都加入
        for edge in graph.edge_list[idx]:
            # 如果頂點看過了，就跳過，反之則加入查看表
            if edge not in visited:
                visited.add(edge)
                queue.append(edge)

    return res


if __name__ == "__main__":
    edges = [
        (0, 1),
        (0, 3),
        (1, 2),
        (1, 4),
        (2, 5),
        (3, 4),
        (3, 6),
        (4, 5),
        (4, 7),
        (5, 8),
        (6, 7),
        (7, 8),
    ]
    g = GraphAdjList(edges)
    print(graph_bfs(g, 0))
    print(graph_bfs(g, 6))
