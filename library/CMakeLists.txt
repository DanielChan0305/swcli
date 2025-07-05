cmake_minimum_required(VERSION 3.14)

project(
    cp_templates
    DESCRIPTION "Templates for competitive programming"
    HOMEPAGE_URL "https://github.com/DanielChan0305/swcli"
)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

Include(FetchContent)

FetchContent_Declare(
  Catch2
  GIT_REPOSITORY https://github.com/catchorg/Catch2.git
  GIT_TAG        v3.8.1
)

FetchContent_MakeAvailable(Catch2)

file(GLOB_RECURSE TEST_SOURCES "*_test.cpp")
add_executable(tests ${TEST_SOURCES})
target_compile_options(tests PRIVATE -fsanitize=address PRIVATE -fsanitize=undefined)
target_link_libraries(tests PRIVATE Catch2::Catch2WithMain PRIVATE -fsanitize=address PRIVATE -fsanitize=undefined)

include(CTest)
include(Catch)
catch_discover_tests(tests)
