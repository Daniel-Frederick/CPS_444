#include <stdio.h>
#include <math.h>

int main() {
    int salaryRanges[9] = {0};
    float sales;
    int salary;

    printf("Enter salesperson's gross sales (-1 to end): ");
    scanf("%f", &sales);

    while (sales != -1) {
        salary = (int)(200 + 0.09 * sales);

        if (salary >= 1000)
            salaryRanges[8]++;
        else
            salaryRanges[(salary - 200) / 100]++;

        printf("Enter salesperson's gross sales (-1 to end): ");
        scanf("%f", &sales);
    }

    printf("\nNumber of salespeople in each salary range:\n");
    printf("$200–299: %d\n", salaryRanges[0]);
    printf("$300–399: %d\n", salaryRanges[1]);
    printf("$400–499: %d\n", salaryRanges[2]);
    printf("$500–599: %d\n", salaryRanges[3]);
    printf("$600–699: %d\n", salaryRanges[4]);
    printf("$700–799: %d\n", salaryRanges[5]);
    printf("$800–899: %d\n", salaryRanges[6]);
    printf("$900–999: %d\n", salaryRanges[7]);
    printf("$1000 and over: %d\n", salaryRanges[8]);

    return 0;
}
