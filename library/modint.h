#ifndef MATH_HPP
#define MATH_HPP

template <typename T, T MOD> class modint{
    private:
        T x;
    public: 
        modint(T value) : x(value % MOD) {}

        friend std::ostream &operator<<(std::ostream& os, const modint& m) {
            return os << m.x;
        }

        friend std::istream &operator>>(std::istream& os, const modint& m){
            return os >> m.x;
        }
};

template <typename T> T cadd(T a, T b)
{
    return (a + b);
}

template <typename T> T cminus(T a, T b){
    return (a - b);
}

template <typename T> T cmultiply(T a, T b){
    return (a * b);
}

#endif