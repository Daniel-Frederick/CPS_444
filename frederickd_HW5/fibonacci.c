#include <stdio.h>

unsigned long fibonacci(unsigned int n) {
    if (n == 0)
        return 0;
    else if (n == 1)
        return 1;

    unsigned long prev = 0, curr = 1, next;

    for (unsigned int i = 2; i <= n; i++) {
        next = prev + curr;
        prev = curr;
        curr = next;
    }

    return curr;
}

int main() {
    unsigned int n;

    printf("Enter the position of the Fibonacci sequence: ");
    scanf("%u", &n);

    printf("Fibonacci(%u) = %lu\n", n, fibonacci(n));

    printf("\nFinding the largest Fibonacci number that fits in unsigned long...\n");

    unsigned long prev = 0, curr = 1, next;
    unsigned int i = 1;

    while (1) {
        next = prev + curr;
        if (next < curr)
            break;

        prev = curr;
        curr = next;
        i++;
    }

    printf("Largest Fibonacci index that fits: %u\n", i);
    printf("Fibonacci(%u) = %lu\n", i, prev);

    return 0;
}

