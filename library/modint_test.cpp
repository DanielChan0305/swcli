#include "modint.h"

#include <catch2/catch_test_macros.hpp>


//TEST_CASE("typical 1e9 + 7 modding arithmetic", "[modint]") {
//	using mint = modint<int, 1'000'000'007>;
    
//    SECTION("constructors and int conversions") {
//        REQUIRE(bool(mint(1'000'000'006) == 1'000'000'006));
//        REQUIRE(bool(mint(-1) == 1'000'000'006));
//    };

//    mint a(-1);
//    SECTION("input and output")
//    {
//        std::stringstream ss;

//        ss << a << ' ' << mint(1'000'000'008);

//        REQUIRE(ss.str() == "10 3");
//    };
//}

//TEST_CASE("special mod"){
//    REQUIRE(bool(mint10(11) == 1));
//}

using mint = modint<int, 1'000'000'007>;
TEST_CASE("basic input output"){

    std::stringstream ss{"10"};
    mint a;
    ss >> a;
    ss.clear();

    REQUIRE(bool(a == 10));

    mint b(11);
    ss << b << ' ' << mint(-1);
    REQUIRE(ss.str() == "11 1000000006");
}

TEST_CASE("comparators and assign"){

}

TEST_CASE("basic arithematics"){

}

