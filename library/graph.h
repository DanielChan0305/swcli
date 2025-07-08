#include<vector>
#include<map>
#include<algorithm>

#define pii std::pair<int, int>

// Weighted graph
template <typename T>
class graph {
    
};


// Unweighted graph
template<>
class graph<void> : public std::vector<std::vector<int>> {
    using edge = pii;

    protected:
        int sz;
        int numEdge;
        std::vector<edge> edges;
    
    public:
        std::vector<int> indegree;

        graph(int size) : std::vector<std::vector<int>>(size + 1), sz(size), numEdge(0), indegree(size + 1){}

        void addUndirectedEdge(const int &u, const int &v){
            edges.push_back({u, v});

            numEdge++;

            (*this)[u].push_back(v);
            (*this)[v].push_back(u);

            indegree[u]++, indegree[v]++;
        }

        void addDirectedEdge(const int &u, const int& v){
            edges.push_back({u, v});

            numEdge++;

            (*this)[u].push_back(v);

            indegree[v]++;
        }

        void dfsOrder(int u, std::vector<bool>& vis, std::vector<int>& ans){
            vis[u] = 1;
            ans.push_back(u);

            for (auto &v : (*this)[u]){
                if (!vis[v]){
                    dfsOrder(v, vis, ans);
                }
            }
        }

        // visits node with small index first
        std::vector<int> dfsOrder(int u){
            std::vector<bool> vis(sz + 1);
            std::vector<int> ans;

            // sort by index
            for (int v = 1; v <= sz; v++){
                std::sort((*this)[v].begin(), (*this)[v].end());
            }

            dfsOrder(u, vis, ans);
            return ans;
        }

        void findBridges(std::map<pii, bool>& ma, std::vector<int>& low, std::vector<int>& tin, int time, int u, int p){
            low[u] = tin[u] = time++;

            for (auto v : (*this)[u]){
                // not visited
                if (v == p) {
                    continue;
                }

                if (low[v] == -1) {
                    findBridges(ma, low, tin, time, v, u);
                }

                low[u] = std::min(low[u], low[v]); 

                if (tin[u] < low[v]){
                    ma[{u, v}] = 1;
                }
            }
        }

        std::vector<int> findBridges(){
            int time = 0;
            std::vector<int> low(sz + 1, -1), tin(sz + 1, -1);
            std::map<pii, bool> ma;

            findBridges(ma, low, tin, time, 1, 1);
            std::vector<int> re;

            for (int i = 0; i < numEdge; i++)
            {
                auto& [u, v] = edges[i];
                if (ma.count({u, v}) + ma.count({v, u}) != 0)
                {
                    re.push_back(i + 1);
                }
            }

            return re;
        }

        int getIndegree(int u){
            return indegree[u];
        }

        int getOutdegree(int u){
            return (*this)[u].size();
        }


        
        // return nodes with 0 indegree
        std::vector<int> getRoots(){
            std::vector<int> re;

            for (int u = 1; u <= sz; u++){
                if (indegree[u] == 0){
                    re.push_back(u);
                }
            }

            return re;
        }

        // return nodes with 0 outdegree
        std::vector<int> getLeaves(){
            std::vector<int> re;
    
            for (int u = 1; u <= sz; u++){
                if ((*this)[u].empty()){
                    re.push_back(u);
                }
            }
    
            return re;
        }
};