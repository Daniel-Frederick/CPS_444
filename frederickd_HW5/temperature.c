#include <stdio.h>

int celsius(int fahrenheit) {
  return (5 * (fahrenheit - 32)) / 9;
}

int fahrenheit(int celsius) {
  return (9 * celsius) / 5 + 32;
}

int main() {
  printf("Celsius to Fahrenheit Conversion Chart\n");
  printf("========================================\n");
  printf("Celsius  Fahrenheit  |  Celsius  Fahrenheit\n");
  printf("-------------------------------------------\n");

  for (int c = 0; c <= 50; c++) {
    printf("%4d     %4d       |  %4d     %4d\n", c, fahrenheit(c), c + 51, fahrenheit(c + 51));
  }
  if (100 % 2 != 0) {
    printf("%4d     %4d\n", 100, fahrenheit(100));
  }

  printf("\n\nFahrenheit to Celsius Conversion Chart\n");
  printf("========================================\n");
  printf("Fahrenheit  Celsius  |  Fahrenheit  Celsius\n");
  printf("-------------------------------------------\n");

  for (int f = 32; f <= 122; f++) {
    printf("%4d        %4d     |  %4d        %4d\n", f, celsius(f), f + 91, celsius(f + 91));
  }
  if ((212 - 32) % 2 != 0) {
    printf("%4d        %4d\n", 212, celsius(212));
  }
  return 0;
}

