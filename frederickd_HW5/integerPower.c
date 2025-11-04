#include <stdio.h>

int integerPower(int base, int exponent) {
    int result = 1;
    
    for (int i = 0; i < exponent; i++) {
        result *= base;
    }
    
    return result;
}

int main() {
    printf("integerPower(3, 4) = %d\n", integerPower(3, 4));
    printf("integerPower(2, 5) = %d\n", integerPower(2, 5));
    printf("integerPower(5, 3) = %d\n", integerPower(5, 3));
    printf("integerPower(10, 2) = %d\n", integerPower(10, 2));
    
    return 0;
}

