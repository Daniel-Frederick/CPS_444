#include <stdio.h>

typedef struct lens {
    float foclen;
    float fstop;
    char brand[30];
} LENS;

int main(void)
{
    LENS lenses[10] = {
        [2] = { .foclen = 500.0f, .fstop = 2.0f, .brand = "Remarkatar" }
    };

    printf("Lens 3: %s, %.1f mm, f/%.1f\n",
           lenses[2].brand, lenses[2].foclen, lenses[2].fstop);

    return 0;
}
