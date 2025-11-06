#include <stdio.h>

int qualityPoints(int average) {
    if (average >= 90 && average <= 100)
        return 4;
    else if (average >= 80)
        return 3;
    else if (average >= 70)
        return 2;
    else if (average >= 60)
        return 1;
    else
        return 0;
}

int main() {
    int average;

    printf("Enter the student's average (0–100): ");
    scanf("%d", &average);

    printf("Quality points: %d\n", qualityPoints(average));

    return 0;
}
