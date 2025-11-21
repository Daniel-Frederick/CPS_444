#include <stdio.h>

char *mystrncpy(char *s1, const char *s2, int n)
{
    int i = 0;
    while (i < n && s2[i] != '\0')
    {
        s1[i] = s2[i];
        i++;
    }
    while (i < n)
    {
        s1[i] = '\0';
        i++;
    }
    return s1;
}

int main(void)
{
    char dest[100];
    char src[100];
    int n;

    while (1)
    {
        printf("Enter source string: ");
        if (!fgets(src, sizeof(src), stdin)) break;
        if (src[0] == 'q') break;

        printf("Enter n: ");
        scanf("%d", &n);
        while (getchar() != '\n'); // flush

        mystrncpy(dest, src, n);
        dest[n] = '\0';   // ensure printing safety

        printf("Result: [%s]\n", dest);
    }

    return 0;
}
