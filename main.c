#include <stdio.h>
#include <stdbool.h>

bool stringsEqual(const char *s1, const char *s2)
{
    int s1_size = sizeof(s1) / sizeof(char);
    int s2_size = sizeof(s2) / sizeof(char);
    if (s1_size != s2_size)
    {
        return false;
    }

    for (int i = 0; i < s1_size; i++)
    {
        if (s1[i] != s2[i])
        {
            return false;
        }
    }
    return true;
}

int main()
{
    char s1[] = "hello world";
    char s2[] = "hello world";
    printf("equal?: %d\n", stringsEqual(s1, s2));
    return 0;
}