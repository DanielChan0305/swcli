#ifndef BIT_H
#define BIT_H

// 0-based bit
template<typename T> class bit{
    private:
        int sz;
        std::vector<T> a;

    public:
        bit(int size): sz(size), a(sz + 5) {}
        // set default size and value 
        bit(int size, T defaultValue) : sz(size), a(sz + 5) {
            for (int i = 0; i < size; i++){
                set(i, defaultValue);
            }
        };

        // set ix position as T value
        void set(int ix, T value){
            // shift for 0 base
            ix++;
            while (ix <= sz)
            {
                a[ix] += value;
                ix += (ix) & (-ix);
            }
        }

        // get the range sum of [1 - ix]
        T getStart(int ix){
            // shift for 0 base
            ix++;
            T ans = 0;
            while (ix > 0){
                ans += a[ix];
                ix -= (ix) & (-ix);
            }

            return ans;
        }

        // get range from [l - r]
        T getRange(int l, int r){
            return getStart(r) - getStart(l - 1);
        }

        // get point value
        T getPoint(int ix){
            return getStart(ix) - getStart(ix - 1);
        }

        // getPoint[1 - size] => Output each getPoint
        void printAllPoints(){
            for (int i = 0; i < sz; i++){
                std::cout << getPoint(i) << ' ';
            }
            std::cout << std::endl;
        }
};

#endif