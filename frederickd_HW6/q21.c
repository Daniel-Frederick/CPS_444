#include <stdio.h>

struct Month {
    char name[20];
    char abbrev[4];
    int days;
    int month_no;
};

extern struct Month months[12];

int total_days_up_to(int month_no)
{
    int total = 0;
    for (int i = 0; i < 12; i++)
    {
        if (months[i].month_no <= month_no)
            total += months[i].days;
    }
    return total;
}

struct Month months[12] = {
    {"January", "Jan", 31, 1},
    {"February", "Feb", 28, 2},
    {"March", "Mar", 31, 3},
    {"April", "Apr", 30, 4},
    {"May", "May", 31, 5},
    {"June", "Jun", 30, 6},
    {"July", "Jul", 31, 7},
    {"August", "Aug", 31, 8},
    {"September", "Sep", 30, 9},
    {"October", "Oct", 31, 10},
    {"November", "Nov", 30, 11},
    {"December", "Dec", 31, 12}
};

int main(void)
{
    int month = 4;
    printf("Total days up to month %d: %d\n", month, total_days_up_to(month));
    return 0;
}
