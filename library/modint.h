#ifndef MODINT_H 
#define MODINT_H

#include<iostream>

template <typename T, T MOD> class modint{
    private:
        T x;
        T mod (T value){
            
            return (value % MOD + MOD) % MOD;
        }

    public: 

        modint(T value) : x(mod(value)) {}
        modint(): x(0) {}

        template<class T2, T2 MOD2>
        friend std::ostream &operator<<(std::ostream &os, const modint<T2, MOD2> &m);
        template<class T2, T2 MOD2>
        friend std::istream &operator>>(std::istream &is, modint<T2, MOD2> &m);

        bool operator==(const modint& m){
            return (this->x == m.x);
        }

        modint &operator=(const modint& m){
            this->x = m.x;
            return *this;
        }

        template<typename T2> modint &operator=(const T2 m){
            this->x = m;
            return *this;
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

        // ^
        template<typename T2> modint operator^(T2 power){
            modint ans = modint(1);
            modint base = modint(this->x);

            while (power){
                if (power % 2){
                    ans *= base;
                }

                base *= base;
                power /= 2;
            }

            return ans;
        }

        template<typename T2> modint &operator^=(T2 power){
            *this = *this ^ power;
            return *this;
        }

        modint inv(){
            return (*this ^ (MOD - 2));
        }
};


template<class T2, T2 MOD2>
std::ostream &operator<<(std::ostream& os, const modint<T2, MOD2>& m) {
    return os << m.x;
}

template<class T2, T2 MOD2>
std::istream &operator>>(std::istream& is, modint<T2, MOD2>& m){
    return is >> m.x;
}

#endif