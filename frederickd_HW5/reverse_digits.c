#include <stdio.h>

int reverseDigits(int number) {
    int reversed = 0;

    while (number != 0) {
        int digit = number % 10;
        reversed = reversed * 10 + digit;
        number /= 10;
    }

    return reversed;
}

int main() {
    int num;

    printf("Enter an integer: ");
    scanf("%d", &num);

    printf("Reversed number: %d\n", reverseDigits(num));

    return 0;
}
