#include <stdio.h>

typedef struct{
    double* data;
    size_t size;
}myVector;

int allocateMemory(myVector* v, int size){
    v->data = (double *)calloc(size, sizeof(double));
    v->size = size;

    if (v->data == NULL){
        return -1;
    }
    return 0;
}

int append(myVector* v, double n){
    double* tmp = (double *)malloc(sizeof(double) * v->size);
    if (tmp==NULL) return -1;

    for (int i = 0; i < v->size; i++){
        tmp[i] = v->data[i];
    }

    free(v->data);

    v->size ++;

    v->data = (double *)malloc(sizeof(double) * (v->size));
    if (v-data == NULL) return -1;
    for(int i = 0;i< v->size-1, i++){
        v->data[i] = tmp[i];
    }

    v->data[v->size-1] = n;
    free(tmp);
    return 0;
}

void freeMemory(myVector* v){
    free(v->data);
    v->size = 0;
}

int main(){
    myVector a;
    allocateMemory(&a, 10);

    
}



