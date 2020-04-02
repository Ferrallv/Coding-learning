#include <iostream>
using namespace std;
int main()
{
// declare array A of length 2 and assign its elements
// to have values 2 and 3
int A[2] = {2,3};
// declare array B of length 2, but don't assign its
// elements (yet)
int B[2];
// loop through both arrays A and B assigning to each
// element of B the corresponding value of A
for (int i = 0; i < 2; i++)
{
B[i] = A[i];
}
// change the first element of B 
B[0]= 100;
// loop through A and B printing out each element
for (int i = 0; i < 2; i++)
{
cout << "A,B: " << A[i] << ", " << B[i] << endl;
}
return 0;
} 