#include <iostream>
#include <fcntl.h>
#include <unistd.h>
#include <filesystem>

#include "index.h"
using namespace std;
int main(int argc, char **argv)
{
    int fd = open("input.in", O_RDWR);
    string s = "";
    while (true)
    {
        char c;
        int a = read(fd, &c, 1);
        if (a == 0)
        {
            break;
        }
        s.push_back(c);
    }
    close(fd);

    s = handle(s);
    fd = open("output.out", O_RDWR);
    write(fd, s.c_str(), s.size());
    close(fd);
}
