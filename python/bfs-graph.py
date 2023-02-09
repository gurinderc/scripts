from collections import deque

def BFS(graph, start_vertex):
    visited = set()
    queue = deque([start_vertex])
    visited.add(start_vertex)
    
    while queue:
        current_vertex = queue.popleft()
        print(current_vertex)
        
        for neighbor in graph[current_vertex]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append(neighbor)

# Example usage
if __name__ == "__main__":
    graph = {
        0: [1, 2],
        1: [2],
        2: [3],
        3: [1, 2]
    }
    BFS(graph, 0)
