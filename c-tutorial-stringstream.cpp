#include <sstream>
#include <vector>
#include <iostream>
using namespace std;

vector<int> parseInts(string str) {
    stringstream ss(str);
    
    vector<int> elements;
    
    for ( ; ; ) {
        int a;
        auto ok = static_cast<bool>(ss >> a);  
        if (ok) {
            elements.push_back(a);
            char ch;        
            auto ok = static_cast<bool>(ss >> ch);   
            if (!ok) {
                break;           
            }                
        } else {
            break;
        }
    }
    return elements;
}

int main() {
    string str;
    cin >> str;
    auto integers = parseInts(str);
    for(int i = 0; i < integers.size(); i++) {
        cout << integers[i] << "\n";
    }
    
    return 0;
}
