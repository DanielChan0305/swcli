#ifndef SEGTREE_H
#define SEGTREE_H

#include<vector>
#include<iostream>

// one-base segtree
template<typename T> class segtree{
    protected:
        std::vector<T> arr;
        std::vector<T> propaArr;

    private:
        int leftchild(const int& x){
            return (2 * x + 1);
        }

        int rightchild(const int& x){
            return (2 * x + 2);
        }

        // merge arr[2*x+1] and arr[2*x + 2]
        T mergeTreeCells(const T& a, const T& b) {
            cout << "please update custom merge function" << endl;
        }

        // apply stored update from propa to 
        T applyUpdateOnCell(const T&a, const T& b){
            cout << "please upadte custom range update function" << endl;
        }

        // accumulate two updates
        T accuUpdate(const T&a, const T&b){
            cout << "please update custom merge update function" << endl;
        }

        void propa(const int&x, const int& lx, const int& rx){
            arr[x] = applyUpdateOnCell(arr[x], propaArr[x]);

            if (rx - lx != 1){
                // not root, can propa
                propaArr[leftchild(x)] = accuUpdate(propaArr[leftchild(x)], propaArr[x]);
                propaArr[rightchild(x)] = accuUpdate(propaArr[rightchild(x)],propaArr[x]);
            }

            propaArr[x] = neut;
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
            propaArr.resize(2 * sz, neutral);
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

            arr[x] = mergeTreeCells(arr[leftchild(x)], arr[rightchild(x)]);
        }

        void pointAssign(int i, T val){
            pointAssign(0, 0, sz, i, val);
        }

        void rangeUpdate(int x, int lx, int rx, int l, int r, T val){
            if (rx <= l || r <= lx)
                return;

            if (l <= lx && rx <= r){
                propa[x] = accuUpdate(propa[x], val);
                return;
            }

            int mx = (lx + rx) / 2;

            rangeUpdate(leftchild(x), lx, mx, l, r, val);
            rangeUpdate(rightchild(x), mx, rx, l, r, val);
        }

        void rangeUpdate(int l, int r, T val){
            rangeUpdate(0, 0, sz, l, r, val);
        }

        T rangeGet(int x, int lx, int rx, int l, int r){
            propa(x, lx, rx);

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
