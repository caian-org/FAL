int _addAndMultiples(int num) {
    return (num + 1) * num;
}

extern "C" {
    int
    __attribute__((visibility("default")))
    addAndMultiplies(int a) {
        return _addAndMultiples(a);
    }
}
