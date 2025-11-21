#include <stdio.h>

int contains(char ch, const char *str)
{
    while (*str)
    {
        if (*str == ch)
            return 1;
        str++;
    }
    return 0;
}

int main(void)
{
    char text[80];
    char c;

    while (1)
    {
        printf("Enter a string: ");
        if (!fgets(text, 80, stdin)) break;
        if (text[0] == 'q') break;

        printf("Enter a character: ");
        c = getchar();
        while (getchar() != '\n');

        if (contains(c, text))
            printf("'%c' IS in the string.\n", c);
        else
            printf("'%c' is NOT in the string.\n", c);
    }
}
