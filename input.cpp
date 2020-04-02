#include <iostream>
using namespace std;

int main() {

    //declares num as a variable that can hold a floating point number
    float num;
    
    // Displays this text to the console
    cout << "Give me a number: ";

    // Takes the user's input and stores it in num
    cin >> num;

    // Displays to the console
    cout << "This is your number doubled: " << num*2;

    return 0;
}