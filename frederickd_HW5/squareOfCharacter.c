#include <stdio.h>

void squareOfCharacters(int side, char fillCharacter) {
    for (int i = 0; i < side; i++) {
        for (int j = 0; j < side; j++) {
            printf("%c", fillCharacter);
        }
        printf("\n");
    }
}

int main() {
    printf("Square with side = 5 and fillCharacter = '#':\n");
    squareOfCharacters(5, '#');
    
    printf("\nSquare with side = 4 and fillCharacter = '*':\n");
    squareOfCharacters(4, '*');
    return 0;
}
