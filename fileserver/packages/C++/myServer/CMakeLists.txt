cmake_minimum_required (VERSION 3.12)

project(server)

aux_source_directory(. CURRENTDIR_SRCS)
add_executable(server ${CURRENTDIR_SRCS})
#target_link_libraries(server)
