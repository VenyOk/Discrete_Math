#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

vector<vector<int>> adj, adj_r;
vector<bool> visited;
vector<int> order, comp, res;

void dfs1(int v) {
    visited[v] = true;
    for (int u : adj[v]) {
        if (!visited[u])
            dfs1(u);
    }
    order.push_back(v);
}

void dfs2(int v, int cl) {
    comp[v] = cl;
    for (int u : adj_r[v]) {
        if (comp[u] == -1)
            dfs2(u, cl);
    }
}

int main() {

    int N, M;
    cin >> N >> M;

    adj.resize(N);
    adj_r.resize(N);
    visited.assign(N, false);
    comp.assign(N, -1);

    for (int i = 0; i < M; i++) {
        int u, v;
        cin >> u >> v;
        adj[u].push_back(v);
        adj_r[v].push_back(u);
    }

    for (int i = 0; i < N; i++) {
        if (!visited[i])
            dfs1(i);
    }

    int cl = 0;
    vector<int> root(N, -1);
    for (int i = 0; i < N; i++) {
        int v = order[N - 1 - i];
        if (comp[v] == -1) {
            dfs2(v, cl);
            root[cl] = v;
            cl++;
        }
    }

    vector<bool> inBase(cl, true);
    for (int i = 0; i < N; i++) {
        for (int v : adj[i]) {
            if (comp[i] != comp[v]) {
                inBase[comp[v]] = false;
            }
        }
    }

    for (int i = 0; i < cl; i++) {
        if (inBase[i]) {
            res.push_back(root[i]);
        }
    }

    sort(res.begin(), res.end());

    for (int v : res) {
        cout << v << " ";
    }
    cout << endl;

    return 0;
}
