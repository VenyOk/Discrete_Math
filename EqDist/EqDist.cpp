#include <iostream>
#include <vector>
#include <queue>
#include <climits>
#include <algorithm>

using namespace std;

int main() {
    int N, M, K;
    cin >> N >> M;
    vector<vector<int>> adj(N);
    for(int i = 0; i < M; i++) {
        int u, v;
        cin >> u >> v;
        adj[u].push_back(v);
        adj[v].push_back(u);
    }

    cin >> K;
    vector<int> refs(K), dist(N, INT_MAX);
    vector<vector<int>> dist_ref(N, vector<int>(K, INT_MAX));

    for(int i = 0; i < K; i++) {
        cin >> refs[i];
        queue<int> q;
        q.push(refs[i]);
        dist_ref[refs[i]][i] = 0;
        while(!q.empty()) {
            int u = q.front(); q.pop();
            for(int v : adj[u]) {
                if(dist_ref[v][i] == INT_MAX) {
                    dist_ref[v][i] = dist_ref[u][i] + 1;
                    q.push(v);
                }
            }
        }
    }

    vector<int> res;
    for(int u = 0; u < N; u++) {
        if((count(dist_ref[u].begin(), dist_ref[u].end(), dist_ref[u][0]) == K) && (dist_ref[u][0]!= INT_MAX)){
            res.push_back(u);
        }
    }

    if(res.empty()) {
        cout << "-" << endl;
    } else {
        sort(res.begin(), res.end());
        for(int u : res) {
            cout << u << " ";
        }
        cout << endl;
    }

    return 0;
}
