#include <stdio.h>

void squareOfAsterisks(int side) {
    for (int i = 0; i < side; i++) {
        for (int j = 0; j < side; j++) {
            printf("*");
        }
        printf("\n");
    }
}

int main() {
    printf("Square with side = 4:\n");
    squareOfAsterisks(4);
    return 0;
}

