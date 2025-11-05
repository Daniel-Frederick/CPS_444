#include <stdio.h>

int timeInSeconds(int hours, int minutes, int seconds) {
    return hours * 3600 + minutes * 60 + seconds;
}

int timeDifference(int hours1, int minutes1, int seconds1, int hours2, int minutes2, int seconds2) {
    int time1 = timeInSeconds(hours1, minutes1, seconds1);
    int time2 = timeInSeconds(hours2, minutes2, seconds2);
    return (time2 > time1) ? (time2 - time1) : (time1 - time2);
}

int main() {
    int h1, m1, s1, h2, m2, s2;
    printf("Enter first time (hours minutes seconds): ");

    scanf("%d %d %d", &h1, &m1, &s1);
    printf("Enter second time (hours minutes seconds): ");

    scanf("%d %d %d", &h2, &m2, &s2);
    int seconds1 = timeInSeconds(h1, m1, s1);
    int seconds2 = timeInSeconds(h2, m2, s2);
    int difference = timeDifference(h1, m1, s1, h2, m2, s2);

    printf("\nFirst time: %d:%02d:%02d = %d seconds since 12:00\n", h1, m1, s1, seconds1);
    printf("Second time: %d:%02d:%02d = %d seconds since 12:00\n", h2, m2, s2, seconds2);
    printf("Time difference: %d seconds\n", difference);
    return 0;
}
