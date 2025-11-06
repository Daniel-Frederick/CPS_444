#include <stdio.h>

int main() {
    int counts[10];
    for (int i = 0; i < 10; i++) {
        counts[i] = 0;
    }

    int bonus[15];
    for (int i = 0; i < 15; i++) {
        bonus[i] = 0;
        bonus[i] += 1;
    }

    float monthlyTemperatures[12];
    printf("Enter 12 monthly temperatures:\n");
    for (int i = 0; i < 12; i++) {
        printf("Month %d: ", i + 1);
        scanf("%f", &monthlyTemperatures[i]);
    }

    int bestScores[5] = {95, 87, 78, 92, 88};
    printf("\nBest Scores:\n");
    for (int i = 0; i < 5; i++) {
        printf("%d\n", bestScores[i]);
    }

    return 0;
}
