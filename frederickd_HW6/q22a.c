#include <stdio.h>
#include <string.h>

typedef struct lens {    /* lens descriptor */
    float foclen;        /* focal length, mm */
    float fstop;         /* aperture */
    char brand[30];      /* brand name */
} LENS;

int main(void)
{
    LENS lenses[10];  // 10-element array of LENS

    // Assign values to the third element (index 2)
    lenses[2].foclen = 500.0f;
    lenses[2].fstop  = 2.0f;
    strcpy(lenses[2].brand, "Remarkatar");

    // Example print to verify
    printf("Lens 3: %s, %.1f mm, f/%.1f\n",
           lenses[2].brand, lenses[2].foclen, lenses[2].fstop);

    return 0;
}
