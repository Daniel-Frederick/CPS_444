#include <stdio.h>

void fetch(char *dest, int n)
{
    for (int i = 0; i < n; i++)
    {
        int ch = getchar();
        if (ch == EOF)
        {
            dest[i] = '\0';
            return;
        }
        dest[i] = ch;
    }
    dest[n] = '\0';
}

int main(void)
{
    char buffer[100];
    printf("Enter text: ");
    fetch(buffer, 10);
    printf("Fetched: [%s]\n", buffer);
    return 0;
}
