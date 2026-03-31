#include<stdio.h>

int main() {
    int n1, n2, n3, conc, q;
    scanf("%d %d %d", &n1, &n2, &n3);
    if (n1 > 9 || n1 < 0 || n2 > 9 || n2 < 0 || n3 > 9 || n3 < 0) {
        printf("DIGITO INVALIDO\n");
        return 0;
    } else {
        conc = (n1 * 100) + (n2 * 10) + n3;
        q = conc * conc;
        printf("%d, %d\n", conc, q);
        return 0;
    }
    return 0;
}
