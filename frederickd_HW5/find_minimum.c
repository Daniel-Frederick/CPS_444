#include <stdio.h>

float findMinimum(float a, float b, float c) {
    float min = a;
    if (b < min)
        min = b;
    if (c < min)
        min = c;
    return min;
}

int main() {
    float x, y, z;

    printf("Enter three floating-point numbers: ");
    scanf("%f %f %f", &x, &y, &z);

    printf("The smallest number is: %.2f\n", findMinimum(x, y, z));

    return 0;
}

