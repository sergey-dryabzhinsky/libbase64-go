#include <libbase64-go.h>
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("Compiled libbase64-go version: %s\n", libbase64_go_Version);
  printf("Loaded libbase64-go version: %s\n", libbase64_go__Version());
  printf("go runtime version used: %s\n", libbase64_go__Version_go());
  return 0;
}