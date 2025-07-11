#ifndef FLOW_H
#define FLOW_H

#include<vector>
#include<queue>
#include<algorithm>

struct flow{
    protected:
        // number of nodes, excluding source and sink
        int sz;
        
        /*
        s -> source [size + 1]
        t -> sink [size + 2]
        */
        int s, t;

        // capacity[u][v]: residual capacity u -> v
        std::vector<std::vector<int>> capacity;

        // adj[u] = {V} : edge from u to {V}
        std::vector<std::vector<int>> adj;

    public :
        flow(int _size, int source, int sink) : sz(_size), s(source), t(sink),
                    capacity(_size + 5, std::vector<int>(_size + 5, 0)), adj(_size + 5) {}

        flow(int _size) : flow(_size, _size + 1, _size + 2) {}

        // add edge between u, v where u, v in {V - s - t}
        void addEdge(int u, int v, int cap){
            capacity[u][v] += cap;

            adj[u].push_back(v);
            adj[v].push_back(u);
        }
        
        // add edge between s, u 
        void addEdgeFromSourceToU(int u, int capacity){
            addEdge(s, u, capacity);
        }

        // add edge between u and t
        void addEdgeFromUToSink(int u, int capacity){
            addEdge(u, t, capacity);
        }

        int edmondKarp(){
            int maxFlow = 0;
            int augment;
            std::vector<int> parent(sz + 5, -1);
            
            auto bfs = [&]()
            {
                // {u, limitingCapacity}
                std::queue<std::pair<int, int>> q;

                q.push({s, __INT32_MAX__});
                parent[s] = s;
                while (q.size())
                {
                    auto [u, limiting] = q.front();
                    q.pop();
    
                    // has a augmenting path to sink
                    if (u == t)
                    {
                        return limiting;
                    }

                    for (auto v : adj[u])
                    {
                        if (parent[v] == -1 && capacity[u][v] > 0)
                        {
                            int newLimiting = std::min(limiting, capacity[u][v]);
    
                            parent[v] = u;
                            q.push({v, newLimiting});
                        }
                    }
                }

                return 0;
            };

            while ((augment = bfs()) != 0){
                maxFlow += augment;

                int v = t;
                int u = parent[v];
                while (u != v){
                    capacity[u][v] -= augment;
                    capacity[v][u] += augment;

                    v = u;
                    u = parent[v];
                }

                fill(parent.begin(), parent.end(), -1);
            }

            return maxFlow;
        }
};

#endif