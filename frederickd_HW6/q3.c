#include <stdio.h>

int max_in_array(const int arr[], int n);

int main(void)
{
    int nums[5] = {12, 48, 7, 99, 23};

    int largest = max_in_array(nums, 5);

    printf("The largest value is: %d\n", largest);

    return 0;
}

int max_in_array(const int arr[], int n)
{
    int max = arr[0];

    for (int i = 1; i < n; i++)
    {
        if (arr[i] > max)
            max = arr[i];
    }

    return max;
}

