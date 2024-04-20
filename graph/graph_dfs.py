from graph_adjacency_list import GraphAdjList


def dfs(graph, vertex, res, visited):
    res.append(vertex)
    visited.add(vertex)
    for edge in graph.edge_list[graph.adj_list.index(vertex)]:
        # 如果沒看過才看
        if edge not in visited:
            # 找到該節點對應的 edge list
            dfs(graph, edge, res, visited)


def graph_dfs(graph: GraphAdjList, start_vet: int) -> list[int]:
    res = []
    # 以下的流程會不斷遞迴執行，直到
    visited = set()
    # 找到對應的 edge list
    dfs(graph, start_vet, res, visited)
    return res


if __name__ == "__main__":
    edges = [
        (0, 1),
        (0, 3),
        (1, 2),
        (2, 5),
        (4, 5),
        (5, 6),
    ]
    g = GraphAdjList(edges)
    print(graph_dfs(g, 0))
    print(graph_dfs(g, 6))
