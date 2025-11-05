#include <stdio.h>

int multiple(int num1, int num2) {
    if (num1 == 0) {
        return 0;
    }
    return (num2 % num1 == 0) ? 1 : 0;
}

int main() {
    int num1, num2;

    printf("Enter two integers (num1 num2): ");
    scanf("%d %d", &num1, &num2);
    
    if (multiple(num1, num2)) {
      printf("%d is a multiple of %d\n", num2, num1);
      return 1;
    } else {
      printf("%d is NOT a multiple of %d\n", num2, num1);
      return 0;
    }
    
    return 0;
}
