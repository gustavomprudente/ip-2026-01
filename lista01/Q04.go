#include <stdio.h>

int main () {
    double kw, s, vk, vc, nv;
    scanf("%lf %lf", &s, &kw);
    vk = (0.7 * s) / 100;
    vc = vk * kw;
    nv = 0.9 * vc;
    printf("Custo por kW: R$ %.2lf\nCusto do consumo: R$ %.2lf\nCusto com desconto: R$ %.2lf\n", vk, vc, nv);
}
