#include <stdio.h>
#include <math.h>
#include <time.h>

int isPrime_half(int n) {
    if (n <= 1)
        return 0;
    for (int i = 2; i <= n / 2; i++) {
        if (n % i == 0)
            return 0;
    }
    return 1;
}

int isPrime_sqrt(int n) {
    if (n <= 1)
        return 0;
    int limit = (int)sqrt(n);
    for (int i = 2; i <= limit; i++) {
        if (n % i == 0)
            return 0;
    }
    return 1;
}

int main() {
    int count_half = 0, count_sqrt = 0;
    clock_t start, end;
    double time_half, time_sqrt;

    printf("Prime numbers between 1 and 10,000 (using sqrt optimization):\n");

    start = clock();
    for (int i = 1; i <= 10000; i++) {
        if (isPrime_sqrt(i)) {
            count_sqrt++;
            printf("%d ", i);
        }
    }
    end = clock();
    time_sqrt = (double)(end - start) / CLOCKS_PER_SEC;

    printf("\n\nTotal primes (sqrt version): %d\n", count_sqrt);
    printf("Time using sqrt method: %.6f seconds\n", time_sqrt);

    start = clock();
    for (int i = 1; i <= 10000; i++) {
        if (isPrime_half(i)) {
            count_half++;
        }
    }
    end = clock();
    time_half = (double)(end - start) / CLOCKS_PER_SEC;

    printf("Total primes (n/2 version): %d\n", count_half);
    printf("Time using n/2 method: %.6f seconds\n", time_half);

    printf("\nPerformance improvement ≈ %.2fx faster using sqrt optimization.\n",
           time_half / time_sqrt);

    return 0;
}

