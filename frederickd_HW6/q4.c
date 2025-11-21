#include <stdio.h>

int index_of_largest(const double arr[], int n);

int main(void)
{
    double nums[5] = {3.3, 9.9, 1.1, 7.7, 4.4};

    int index = index_of_largest(nums, 5);

    printf("Index of largest value: %d\n", index);
    printf("Largest value: %.2f\n", nums[index]);

    return 0;
}

int index_of_largest(const double arr[], int n)
{
    int max_index = 0;

    for (int i = 1; i < n; i++)
    {
        if (arr[i] > arr[max_index])
            max_index = i;
    }

    return max_index;
}
