#include <stdio.h>

int index_of(const char *str, const char *sub)
{
    int i = 0, j = 0;

    if (*sub == '\0')
        return 0;  // empty substring begins at index 0

    while (str[i] != '\0')
    {
        j = 0;
        while (sub[j] != '\0' && str[i + j] == sub[j])
            j++;

        if (sub[j] == '\0')   // matched full substring
            return i;

        i++;
    }
    return -1;
}

int main(void)
{
    char s1[100];
    char s2[100];

    while (1)
    {
        printf("Enter main string (q to quit): ");
        if (!fgets(s1, sizeof(s1), stdin)) break;
        if (s1[0] == 'q') break;

        printf("Enter substring: ");
        fgets(s2, sizeof(s2), stdin);

        // remove newline from s2
        for (int i = 0; s2[i]; i++)
            if (s2[i] == '\n')
                s2[i] = '\0';

        int idx = index_of(s1, s2);

        printf("index_of() = %d\n", idx);
    }

    return 0;
}
