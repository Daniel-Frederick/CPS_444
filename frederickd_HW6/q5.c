#include <stdio.h>

double diff_largest_smallest(const double arr[], int n);

int main(void)
{
    double nums[5] = {3.5, 10.1, 2.0, 7.7, -1.2};

    double diff = diff_largest_smallest(nums, 5);

    printf("Difference between largest and smallest: %.2f\n", diff);

    return 0;
}

double diff_largest_smallest(const double arr[], int n)
{
    double min = arr[0];
    double max = arr[0];

    for (int i = 1; i < n; i++)
    {
        if (arr[i] < min)
            min = arr[i];
        if (arr[i] > max)
            max = arr[i];
    }

    return max - min;
}
