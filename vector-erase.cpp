#include <vector>
#include <iostream>
#include <algorithm>
#include <iterator>
using namespace std;

int main() {
    vector<int> v;
    int x, a, b;  

    int size; 
    cin >> size;
        
    
    copy_n(istream_iterator<int>(cin), size, back_inserter(v));  

    cin >> x;  
    v.erase(v.begin()+(--x));

    cin >> a >> b;  
    v.erase(v.begin()+(--a), v.begin()+(--b));

    cout << v.size() << endl; 
        
    
    copy(v.begin(), v.end(), ostream_iterator<int>(cout, " ")); 
        
    return 0;
}
