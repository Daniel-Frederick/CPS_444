#include <stdio.h>

int perfect(int number) {
    int sum = 0;

    for (int i = 1; i <= number / 2; i++) {
        if (number % i == 0) {
            sum += i;
        }
    }

    return sum == number;
}

int main() {
    printf("Perfect numbers between 1 and 1000:\n");

    for (int num = 1; num <= 1000; num++) {
        if (perfect(num)) {
            printf("%d = ", num);

            int first = 1;
            for (int i = 1; i <= num / 2; i++) {
                if (num % i == 0) {
                    if (!first)
                        printf(" + ");
                    printf("%d", i);
                    first = 0;
                }
            }

            printf("\n");
        }
    }

    return 0;
}

