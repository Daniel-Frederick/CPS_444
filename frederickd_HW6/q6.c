#include <stdio.h>

void reverse(double arr[], int n);

int main(void)
{
    double nums[5] = {1.1, 2.2, 3.3, 4.4, 5.5};

    reverse(nums, 5);

    printf("Reversed array: ");
    for (int i = 0; i < 5; i++)
        printf("%.1f ", nums[i]);
    printf("\n");

    return 0;
}

void reverse(double arr[], int n)
{
    int left = 0;
    int right = n - 1;

    while (left < right)
    {
        double temp = arr[left];
        arr[left] = arr[right];
        arr[right] = temp;

        left++;
        right--;
    }
}
