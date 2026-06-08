import heapq

class Solution:
    def assignTasks(self, servers: list[int], tasks: list[int]) -> list[int]:
        server_heap = [(server, idx) for idx, server in enumerate(servers)]
        heapq.heapify(server_heap)
        queue_heap = []
        results = []
        t = 0
        for i, task in enumerate(tasks):
            if not server_heap:
                t = queue_heap[0][0]
            while queue_heap and queue_heap[0][0] == t:
                _, server, idx = heapq.heappop(queue_heap)
                heapq.heappush(server_heap, (server, idx))
            server, idx = heapq.heappop(server_heap)
            heapq.heappush(queue_heap, (task+t, server, idx))
            results.append(idx)
            t = max(t, i+1)

        return results
