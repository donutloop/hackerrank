#include <bits/stdc++.h>

using namespace std;



int main()
{
    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    int k;
    k = 9;
    
    if (n > k) {
        cout << "Greater than 9";
    } else if (n == k--) {
        cout << "nine \n";
    } else if (n == k--) {
        cout << "eight \n";
    } else if (n == k--) {
        cout << "seven \n";
    }  else if (n == k--) {
        cout << "six \n";
    }  else if (n == k--) {
        cout << "five \n";
    }  else if (n == k--) {
        cout << "four \n";
    }  else if (n == k--) {
        cout << "three \n";
    }  else if (n == k--) {
        cout << "two \n";
    } else if (n == k--) {
        cout << "one \n";
    }    

    return 0;
}
