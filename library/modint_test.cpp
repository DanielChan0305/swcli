#include "modint.h"

#include <catch2/catch_test_macros.hpp>


using mint10 = modint<int, 10>;
TEST_CASE("basic terminal input output"){
    SECTION("input mint from terminal"){
        std::stringstream ss{"9"};
        mint10 a;
        ss >> a;

        REQUIRE((a == 9));
    }

    SECTION("output mint to terminal"){
        std::stringstream ss;
        mint10 b(11);

        // lvalue, rvalue
        ss << b << ' ' << mint10(-1);

        REQUIRE(ss.str() == "1 9");
    }
}

TEST_CASE("construct and converge"){
    REQUIRE((mint10(1) == 1));
    REQUIRE((mint10(10003) == 3));

    REQUIRE((mint10(-1) == 9));
    REQUIRE((mint10(-192) == 8));

    REQUIRE(bool(mint10(2)));
    REQUIRE(bool(mint10(-1)));
    REQUIRE(!bool(mint10(0)));
}


TEST_CASE("comparators and assigners"){
    mint10 a(5), b(7);

    // mint, int
    REQUIRE((a == 5));
    REQUIRE((a == mint10(15)));

    // mint, int
    REQUIRE((a != 7));
    REQUIRE((a != b));

    a = mint10(7);
    REQUIRE((a == b));

    a += mint10(4);
    REQUIRE((a == 1));

    a -= mint10(23);
    REQUIRE((a == 8));

    a *= mint10(3);
    REQUIRE((a == 4));

    a ^= 7;
    REQUIRE((a == 4));

    mint10 c(9);
    c++;
    REQUIRE((c == 0));

    c--;
    REQUIRE((c == 9));
}

TEST_CASE("arithematics"){
    REQUIRE((mint10(5) + mint10(7) == 2));
    REQUIRE((mint10(0) - mint10(1) == 9));
    REQUIRE((mint10(5) * mint10(4) == 0));
    REQUIRE((mint10(7) ^ 3 == 3));

    // inversion using fermats little theorem -> require prime mod
    REQUIRE((modint<int, 7>(5) * modint<int, 7>(5).inv() == 1));
}