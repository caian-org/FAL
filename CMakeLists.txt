cmake_minimum_required(VERSION 3.0)
project(FAL VERSION 0.1.0 DESCRIPTION "Ford Assembly Line")

# -std=c++11
set(CMAKE_CXX_STANDARD 11)

# -fvisibility=hidden
set(CMAKE_CXX_VISIBILITY_PRESET hidden)

# -fPIC
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)
conan_basic_setup()

add_library(FAL SHARED
    FAL/FAL.cpp
)

target_link_libraries(FAL ${CONAN_LIBS})
