#include <stdio.h>

int main() {
    int numbers[20];
    int count = 0;
    int num, i, isDuplicate;
    
    printf("Enter 20 numbers (between 10 and 100):\n");
    
    for (int input = 0; input < 20; input++) {
        printf("Number %d: ", input + 1);
        scanf("%d", &num);
        
        if (num < 10 || num > 100) {
            printf("Invalid! Number must be between 10 and 100.\n");
            input--;
            continue;
        }
        
        isDuplicate = 0;
        for (i = 0; i < count; i++) {
            if (numbers[i] == num) {
                isDuplicate = 1;
                break;
            }
        }
        
        if (!isDuplicate) {
            numbers[count] = num;
            count++;
            printf("  -> %d (unique)\n", num);
        } else {
            printf("  -> %d is a duplicate, not added.\n", num);
        }
    }
    
    printf("\nUnique numbers entered: ");
    for (i = 0; i < count; i++) {
        printf("%d ", numbers[i]);
    }
    printf("\nTotal unique numbers: %d\n", count);
    
    return 0;
}
