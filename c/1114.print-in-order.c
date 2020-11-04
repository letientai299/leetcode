#include<pthread.h>

// this is slow due to the while loop
// faster solutions use more than 1 lock.

typedef struct {
    // User defined data may be declared here.
  int c;
  pthread_mutex_t lock;
} Foo;

Foo* fooCreate() {
    Foo* obj = (Foo*) malloc(sizeof(Foo));
    obj->c = 0;

    // Initialize user defined data here.

    return obj;
}

void first(Foo* obj) {
    pthread_mutex_lock(&obj->lock);
    // printFirst() outputs "first". Do not change or remove this line.
    printFirst();
    obj->c=2;
    pthread_mutex_unlock(&obj->lock);
}

void second(Foo* obj) {

    while(obj->c <= 1){};

    pthread_mutex_lock(&obj->lock);
    // printSecond() outputs "second". Do not change or remove this line.
    printSecond();
    obj->c=3;
    pthread_mutex_unlock(&obj->lock);
}

void third(Foo* obj) {
    while(obj->c<=2){};

    pthread_mutex_lock(&obj->lock);
    // printThird() outputs "third". Do not change or remove this line.
    printThird();
    obj->c=3;
    pthread_mutex_unlock(&obj->lock);
}

void fooFree(Foo* obj) {
    // User defined data may be cleaned up here.
    pthread_mutex_destroy(&obj->lock);
}
