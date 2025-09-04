#ifndef TRIE_H
#define TRIE_H

struct trie{
    private:
        int val;
        trie *left, *right;

    public:
        trie() : val(0) {}
};

#endif