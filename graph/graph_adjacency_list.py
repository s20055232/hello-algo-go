from pprint import pprint


class GraphAdjList:
    def __init__(self, edges: list[tuple[int, int]]) -> None:
        self.adj_list: list[int] = []
        self.edge_list: list[list[int]] = []
        for edge in edges:
            self.add_vertex(edge[0])
            self.add_vertex(edge[1])
            self.add_edge(edge)

    def add_vertex(self, vertex: int):
        if vertex not in self.adj_list:
            self.adj_list.append(vertex)
            self.edge_list.append([])

    def del_vertex(self, vertex: int):
        idx = self.adj_list.index(vertex)
        self.adj_list.remove(vertex)
        self.edge_list.pop(idx)
        for edge in self.edge_list:
            try:
                edge.remove(vertex)
            except ValueError:
                pass

    def add_edge(self, edge: tuple[int, int]):
        if edge[0] not in self.edge_list[self.adj_list.index(edge[0])]:
            self.edge_list[self.adj_list.index(edge[0])].append(edge[1])
        if edge[1] not in self.edge_list[self.adj_list.index(edge[1])]:
            self.edge_list[self.adj_list.index(edge[1])].append(edge[0])

    def del_edge(self, edge: tuple[int, int]):
        idx = self.adj_list.index(edge[0])
        self.edge_list[idx].remove(edge[1])
        idx = self.adj_list.index(edge[1])
        self.edge_list[idx].remove(edge[0])

    def print(self):
        print(self.adj_list)
        pprint(self.edge_list)


if __name__ == "__main__":
    edges = [(1, 2), (2, 3), (1, 3)]
    g = GraphAdjList(edges)
    g.print()

    g.add_vertex(4)
    g.print()

    g.del_vertex(4)
    g.print()

    g.del_vertex(2)
    g.print()

    g.add_vertex(4)
    g.add_edge((4, 1))
    g.print()

    g.del_edge((1, 4))
    g.print()
