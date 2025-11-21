#include <stdio.h>

#define CSIZE 4
#define NUM_GRADES 3

    char first[30];
    char last[30];
};

struct student {
    struct name student_name;
    float grades[NUM_GRADES];
    float average;
};

void get_grades(struct student *s, int count) {
    for (int i = 0; i < count; i++) {
        printf("Enter %d grades for %s %s:\n",
               NUM_GRADES, s[i].student_name.first, s[i].student_name.last);
        for (int j = 0; j < NUM_GRADES; j++) {
            printf("Grade %d: ", j + 1);
            scanf("%f", &s[i].grades[j]);
        }
    }
}

void calc_averages(struct student *s, int count) {
    for (int i = 0; i < count; i++) {
        float sum = 0.0f;
        for (int j = 0; j < NUM_GRADES; j++)
            sum += s[i].grades[j];
        s[i].average = sum / NUM_GRADES;
    }
}

void print_students(const struct student *s, int count) {
    for (int i = 0; i < count; i++) {
        printf("%s %s: ", s[i].student_name.first, s[i].student_name.last);
        for (int j = 0; j < NUM_GRADES; j++)
            printf("%.1f ", s[i].grades[j]);
        printf("Average: %.2f\n", s[i].average);
    }
}

void print_class_averages(const struct student *s, int count) {
    float totals[NUM_GRADES] = {0};
    for (int j = 0; j < NUM_GRADES; j++) {
        for (int i = 0; i < count; i++)
            totals[j] += s[i].grades[j];
    }
    printf("Class averages for each grade:\n");
    for (int j = 0; j < NUM_GRADES; j++)
        printf("Grade %d: %.2f\n", j + 1, totals[j]/count);
}

int main(void) {
    struct student class[CSIZE] = {
        { {"Alice", "Smith"} },
        { {"Bob", "Jones"} },
        { {"Carol", "Taylor"} },
        { {"David", "Brown"} }
    };

    get_grades(class, CSIZE);
    calc_averages(class, CSIZE);
    printf("\nStudent information:\n");
    print_students(class, CSIZE);
    printf("\n");
    print_class_averages(class, CSIZE);

    return 0;
}

