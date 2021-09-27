#include <stdio.h>
#include <stdlib.h>

#define SIZE 256

int main(){
  FILE* file;

  char line[SIZE];
  line[0]='\0';
  file = fopen("hello.txt", "r");
  if (file == NULL){
    printf("Fail\n");
    return 1;
  }
  while (fgets(line, SIZE, file) != NULL){
    printf("%s", line);
  }
  fclose(file);
  return 0;
}
