#ifndef MATH_HPP
#define MATH_HPP

template <typename T, T MOD> class modint{
    private:
        T x;
        T mod (T value){
            
            return (value % MOD + MOD) % MOD;
        }

    public: 

        modint(T value) : x(mod(value)) {}

        friend std::ostream &operator<<(std::ostream& os, const modint& m) {
            return os << m.x;
        }

        friend std::istream &operator>>(std::istream& is, const modint& m){
            return is >> m.x;
        }

        // + 
        modint operator+(const modint& m){
            return modint(this->x + m.x);
        }

        modint &operator+=(const modint& m){
            *this = *this + m;
            return *this;
        }

        modint &operator++(){
            *this = modint(this->x + 1);
            return *this;
        }

        modint operator++(int){
            modint cpy(*this);
            ++*this;
            return cpy;
        }

        // -
        modint operator-(const modint& m){
            return modint(this->x - m.x);
        }

        modint &operator-=(const modint& m){
            *this = *this - m;
            return *this;
        }

        modint &operator--(){
            *this = modint(this->x - 1);
            return *this;
        }

        modint operator--(int){
            modint cpy(*this);
            --*this;
            return cpy;
        }

        // *
        modint operator*(const modint& m){
            return modint(this->x * m.x);
        }

        modint &operator*=(const modint& m){
            *this = *this * m;
            return *this;
        }
};



#endif