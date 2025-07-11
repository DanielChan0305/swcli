#ifndef SEGTREE_H
#define SEGTREE_H

#include<vector>
#include<iostream>

// one-base segtree
template<typename T> class segtree{
    protected:
        std::vector<T> arr;

    private:
        int leftchild(const int& x){
            return (2 * x + 1);
        }

        int rightchild(const int& x){
            return (2 * x + 2);
        }

    public:
        int sz;
        T neut;

        segtree(int size, T neutral): sz(1), neut(neutral){
            while (sz < (size + 5))
            {
                sz *= 2;
            }

            arr.resize(2 * sz, neutral);
        }

        // custom merge function
        T merge(const T& a, const T& b) {
            cout << "please update custom merge function" << endl;
        }

        void pointAssign(int x, int lx, int rx, int i, T val){
            if (rx - lx == 1){
                arr[x] = val;
                return;
            }

            int mx = (lx + rx) / 2;

            if (i < mx){
                pointAssign(leftchild(x), lx, mx, i, val);
            }else{
                pointAssign(rightchild(x), mx, rx, i, val);
            }

            arr[x] = merge(arr[leftchild(x)], arr[rightchild(x)]);
        }

        void pointAssign(int i, T val){
            pointAssign(0, 0, sz, i, val);
        }

        T rangeGet(int x, int lx, int rx, int l, int r){
            if (rx <= l || r <= lx){
                return neut;
            }

            if (l <= lx && rx <= r){
                return arr[x];
            }

            int mx = (lx + rx) / 2;

            T leftRes = rangeGet(leftchild(x), lx, mx, l, r);
            T rightRes = rangeGet(rightchild(x), mx, rx, l, r);

            return merge(leftRes, rightRes);
        }

        // return range (l, r]
        T rangeGet(int l, int r){
            return rangeGet(0, 0, sz, l, r);
        }
};

#endif
