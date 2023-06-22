#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

const int MAXN = 1'000'000;

void dfs(int v, int c, vector<int>& used, vector<vector<int>>& g, vector<int>& vertices, vector<int>& edges) {
    used[v] = c;
    vertices[c]++;
    for (auto i : g[v]) {
        edges[c]++;
        if (used[i] == 0) {
            dfs(i, c, used, g, vertices, edges);
        }
    }
    edges[c] /= 2;
}

int main() {
    int n, m;
    cin >> n >> m;

    vector<vector<int>> g(n);
    vector<pair<int, int>> edges(m);

    for (int i = 0; i < m; i++) {
        int x, y;
        cin >> x >> y;
        edges[i] = make_pair(x, y);
        g[x].push_back(y);
        g[y].push_back(x);
    }

    vector<int> used(n, 0);
    int c = 1;
    vector<int> vertices(n+1, 0);
    vector<int> edges_count(n+1, 0);
    for (int i = 0; i < n; i++) {
        if (used[i] == 0) {
            dfs(i, c, used, g, vertices, edges_count);
            c++;
        }
    }

    int idx = max_element(vertices.begin(), vertices.end()) - vertices.begin();
    for (int i = 0; i < c; i++) {
        if (vertices[i] == vertices[idx] && edges_count[i] > edges_count[idx]) {
            idx = i;
        }
    }

    vector<vector<bool>> s(c, vector<bool>(n, false));
    for (int i = 0; i < n; i++) {
        s[used[i]][i] = true;
    }

    cout << "graph {\n";
    for (int i = 0; i < n; i++) {
        cout << "\t" << i;
        if (s[idx][i]) {
            cout << " [color = red]";
        }
        cout << "\n";
    }

    for (const auto& edge : edges) {
        cout << "\t" << edge.first << " -- " << edge.second;
        if (s[idx][edge.first] && s[idx][edge.second]) {
            cout << " [color = red]";
        }
        cout << "\n";
    }
    cout << "}\n";

    return 0;
}
