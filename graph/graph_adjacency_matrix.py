from collections import namedtuple


Edge = namedtuple("Edge", ["a", "b"])


class GraphAdjMat:
    def __init__(self, vertices: list[int], edges: list[Edge]) -> None:
        # wrong example: reference first rows
        # self.matrix: list[list[int]] = [[0] * len(vertices)] * len(vertices)
        self.matrix: list[list[int]] = []
        self.vertices: list[int] = []
        for vertex in vertices:
            self.add_vertex(vertex)

        for edge in edges:
            self.add_edge(edge)
            self.matrix[self.vertices.index(edge.a)][self.vertices.index(edge.b)] = 1
            self.matrix[self.vertices.index(edge.b)][self.vertices.index(edge.a)] = 1

    def size(self):
        return len(self.vertices)

    def add_vertex(self, vertex: int):
        self.vertices.append(vertex)
        n = self.size()
        new_rows = [0] * n
        self.matrix.append(new_rows)
        for n in range(len(self.matrix) - 1):
            self.matrix[n].append(0)

    def remove_vertex(self, vertex: int):
        idx = self.vertices.index(vertex)
        self.vertices.pop(idx)
        for rows in self.matrix:
            rows.pop(idx)
        self.matrix.pop(idx)

    def _find_and_change(self, edge: Edge, link=True):
        if edge.a == edge.b:
            raise ValueError()
        index_a = self.vertices.index(edge.a)
        index_b = self.vertices.index(edge.b)
        self.matrix[index_a][index_b] = 1 if link is True else 0
        self.matrix[index_b][index_a] = 1 if link is True else 0

    def add_edge(self, edge: Edge):
        self._find_and_change(edge, True)

    def del_edge(self, edge: Edge):
        self._find_and_change(edge, False)


if __name__ == "__main__":
    from pprint import pprint

    vertices = [5, 4, 3, 2, 1]
    edges = [Edge(5, 3), Edge(4, 3), Edge(3, 5), Edge(2, 1), Edge(4, 1)]
    g = GraphAdjMat(vertices, edges)
    print(g.vertices)
    pprint(g.matrix)
    for i in range(g.size()):
        assert g.matrix[i][i] == 0, "WRONG!!"

    g.add_edge(Edge(5, 4))
    print(vertices)
    pprint(g.matrix)

    g.del_edge(Edge(5, 4))
    print(g.vertices)
    pprint(g.matrix)

    try:
        g.add_edge(Edge(6, 4))
    except ValueError:
        print("Expected Wrong!")

    try:
        g.add_edge(Edge(1, 1))
    except ValueError:
        print("Expected Wrong!")

    g.remove_vertex(3)
    print(g.vertices)
    pprint(g.matrix)
