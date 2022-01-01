#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <set>
#include <map>
#include <algorithm>
#include <iterator>
using namespace std;


int main() {

    map<string,int>m;
    
    int size; 
    cin >> size;
        
    for (int i = 0; i < size; i++) {
        int type;
        string name; 
        cin >> type;
        cin >> name;
        
        if (type == 1) {
            int value;
            cin >> value;
            std::map<string, int>::iterator it = m.find(name.c_str()); 
            if (it != m.end()) {
                it->second = it->second + value;
            } else {
                m.insert(make_pair(name.c_str(), value));
            }    
       }
        
        if (type == 2) {
            m.erase(name.c_str());
        }
        
        
        if (type == 3) {
            printf("%d\n", m[name.c_str()]);
        }
    }
    
        
    return 0;
}


