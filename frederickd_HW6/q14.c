#include <stdio.h>

char *find_char(const char *str, char ch)
{
    while (*str)
    {
        if (*str == ch)
            return (char *)str;
        str++;
    }
    return NULL;
}

int main(void)
{
    char word[80];
    char c;

    while (1)
    {
        printf("Enter a word: ");
        if (!fgets(word, 80, stdin)) break;
        if (word[0] == 'q') break;

        printf("Enter a character to search for: ");
        c = getchar();
        while (getchar() != '\n');

        char *p = find_char(word, c);

        if (p)
            printf("Found '%c' at position %ld\n", c, p - word);
        else
            printf("Character not found.\n");
    }
}
