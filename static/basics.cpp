/* C++ Basics - This is practice program for learning C++*/

// Must include iostream to use cout
#include <iostream>
// If not, use namespace std
using namespace std;

// Every C++ program must have a main function, this is the function that the program will execute
int main() {
    std::cout << "Hello World! We are printing this statement with including iostream.";
    cout << endl << "Hellow world, we are printing this statement using namespace std! We also use 'endl' as linespace!";

    int num1 = 70;
    double num2 = 256.783;
    char ch = 'D';
    // endl manipulator is used to insert a new line
    cout << endl << num1 << endl;
    cout << num2 << endl;
    cout << "character: " << ch << endl;

    // cin takes an input
    int num;
    // Note the arrows are pointing right now
    cout << "Please specify an integer!";
    cin >> num;
    cout << "The number is: " << num << "." << endl;

    char a;
    cout << "Now, specify an integer and a character!" << endl;
    cin >> num >> a;
    cout << "Integer: " << num << endl;
    cout << "Character: " << a; 
    
    // Data conversion
    int num_int = 9;
    double num_double;

    // Implicit conversion - assigning int value to a double variable
    num_double = num_int;
    cout << "num_int = " << num_int << endl;
    cout << "num_double = " << num_double << endl;

    // Explicit converison 
    // cast notation - C-type type casting
    int num_int_one = 1;
    num_double = (double)num_int_one;
    // or
    num_double = double(num_int_one);

    // Operators
    // ++ increment and -- decrement operators, increase and decrease by 1
    // && AND, || OR and ! NOT
    return 0;
}

// Variables
// Data type then name of variable, always end with semicolon
int age = 24;

// Constant variables cannot be changed
const int LIGHT_SPEED = 299792458;

// Fundamental Data Types
int salary = 90000;

// float is 4 bytes and double is 8 bytes. double has two times the precision of float
float area = 64.74;
double volume = 134.64534;
double distance = 45E12;

// char is used for characters, enclosed in single quotes
char test = 'h';

// normal boolean
bool cond = false;

// Modifiers
long b = 4523232;
long int c = 2345342;
long double d = 233434.56343;