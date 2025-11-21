#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void)
{
    int n;
    printf("How many words do you wish to enter? ");
    scanf("%d", &n);
    while (getchar() != '\n');

    char **words = (char **)malloc(n * sizeof(char *));
    if (!words)
    {
        printf("Memory allocation failed.\n");
        return 1;
    }

    char temp[100];
    printf("Enter %d words now:\n", n);
    for (int i = 0; i < n; i++)
    {
        scanf("%99s", temp);
        int len = strlen(temp);
        words[i] = (char *)malloc((len + 1) * sizeof(char));
        if (!words[i])
        {
            printf("Memory allocation failed.\n");
            return 1;
        }
        strcpy(words[i], temp);
    }

    printf("Here are your words:\n");
    for (int i = 0; i < n; i++)
        printf("%s\n", words[i]);

    // free allocated memory
    for (int i = 0; i < n; i++)
        free(words[i]);
    free(words);

    return 0;
}
