# https://mathhelpplanet.com/static.php?p=uravneniya-kolmogorova
# primer 2-4
from numpy.ma.core import absolute


def primer2():
    from cvxopt.modeling import variable, op
    import time

    # start = time.time()
    p = variable(4, 'p')
    lam = [[0, 1, 2, 0], [2, 0, 0, 2], [3, 0, 0, 1], [0, 3, 2, 0]]
    z = 0
    mass1 = ((lam[0][1] + lam[0][2]) * p[0] == lam[1][0] * p[1] + lam[2][0] * p[2])
    mass2 = ((lam[1][0] + lam[1][3]) * p[1] == lam[0][1] * p[0] + lam[3][1] * p[3])
    mass3 = ((lam[2][0] + lam[2][3]) * p[2] == lam[0][2] * p[0] + lam[3][2] * p[3])
    mass4 = ((lam[3][1] + lam[3][2]) * p[3] == lam[1][3] * p[1] + lam[2][3] * p[2])
    mass5 = (p[0] + p[1] + p[2] + p[3] == 1)
    x_0_1 = (0 <= p <= 1)
    problem = op(z, [mass1, mass2, mass3, mass4, mass5, x_0_1])
    problem.solve(solver='glpk')

    # print("Результат Xopt:")
    # for i in p.value:
    #     print(i)
    # print("Стоимость доставки:", problem.objective.value()[0])
    # stop = time.time()
    # print("Время :", stop - start)
    return [round(i, 2) for i in p.value]


def primer3():
    from cvxopt.modeling import variable, op
    import time
    pEarly = primer2()
    print(pEarly)
    _D = (pEarly[0] + pEarly[2]) * 10 + (pEarly[0] + pEarly[1]) * 6 - (pEarly[1] + pEarly[3]) * 4 - (
            pEarly[2] + pEarly[3]) * 2

    print("_D = ", _D)

    p = variable(4, 'p')
    lam = [[0, 1, 2, 0], [4, 0, 0, 2], [6, 0, 0, 1], [0, 6, 4, 0]]
    z = 0
    mass1 = ((lam[0][1] + lam[0][2]) * p[0] == lam[1][0] * p[1] + lam[2][0] * p[2])
    mass2 = ((lam[1][0] + lam[1][3]) * p[1] == lam[0][1] * p[0] + lam[3][1] * p[3])
    mass3 = ((lam[2][0] + lam[2][3]) * p[2] == lam[0][2] * p[0] + lam[3][2] * p[3])
    mass4 = ((lam[3][1] + lam[3][2]) * p[3] == lam[1][3] * p[1] + lam[2][3] * p[2])
    mass5 = (p[0] + p[1] + p[2] + p[3] == 1)
    x_0_1 = (0 <= p <= 1)
    problem = op(z, [mass1, mass2, mass3, mass4, mass5, x_0_1])
    problem.solve(solver='glpk')

    p = [round(i, 2) for i in p.value]
    l = 0
    for i in p:
        print(f"new p[{l}] -", i)
        l += 1
    _D1 = (p[0] + p[2]) * 10 + (p[0] + p[1]) * 6 - (p[1] + p[3]) * 8 - (p[2] + p[3]) * 4

    print("_D1 = ", _D1)

    if _D1 > _D:
        print(
            f"экономическая целесообразность ускорения\n ремонтов узлов очевидна\n (_D1 > _D примерно на {int(((round(_D1 / _D, 2)) * 100) - 100)}%)")
    else:
        print(
            f"экономическая целесообразность ускорения\n ремонтов узлов очевидна\n (_D > _D1 примерно на {int(((round(_D / _D1, 2)) * 100) - 100)}%)")


def primer4():
    p0 = (1 + 1 / 4 + (2 / 3) * (1 / 4)) ** (-1)
    print("p0 =", p0)
    p1 = (1 / 4) * p0
    print("p1 =", p1)
    p2 = (2 / 3) * (1 / 4) * p0
    print("p2 =", p2)
    print(f"Вероятность S0 = {(p0 * 100):.2f}%")
    print(f"Вероятность S1 = {(p1 * 100):.2f}%")
    print(f"Вероятность S2 = {(p2 * 100):.2f}%")


# https://mathhelpplanet.com/static.php?p=smo-s-otkazami
# primer 5-7
def primer5():
    lam = 90  # Заявок в час
    t_ob = 2  # Минуты 1 обращение
    intensive_potoka = (1 / t_ob) * 60  # Интенсивность обращения 30 за час
    print(intensive_potoka)
    otn_pr_spo_SMO = intensive_potoka / (lam + intensive_potoka)  # Относительная пропускная способность СМО
    print(otn_pr_spo_SMO)
    p_otkaza = 1 - otn_pr_spo_SMO  # Вероятность отказа в обслуживании составит
    absolute = lam * otn_pr_spo_SMO  # Абсолютная пропускная способность СМО,
    # т.е. в среднем в час будут обслужены 22,5 заявки на переговоры
    print(f"В среднем в час будут обслужены "
          f"{absolute} заявки на переговоры. Очевидно, что при наличии только одного телефонного номера СМО будет плохо справляться с потоком заявок.")


def primer6():
    import math

    def calculate_smo(arrival_rate, service_rate, target_Q):
        n = 1
        results = []
        while True:
            rho = arrival_rate / service_rate
            P0 = 1 / sum(rho ** k / math.factorial(k) for k in range(n + 1))
            Q = 1 - (rho ** n / math.factorial(n)) * P0
            A = arrival_rate * Q
            avg_busy = A / service_rate
            results.append({"n": n, "Q": Q, "A": A, "avg_busy": avg_busy})
            if Q >= target_Q:
                break
            n += 1
        return n, results

    arrival_rate = 90
    service_rate = 30
    target_Q = 0.9

    optimal_n, results = calculate_smo(arrival_rate, service_rate, target_Q)

    print(f"Оптимальное число каналов: {optimal_n}")
    for row in results:
        print(f"n={row['n']}, Q={row['Q']:.2f}, A={row['A']:.1f}, avg_busy={row['avg_busy']:.2f}")

def primer7():
    import math

    # Исходные данные
    n = 3  # число каналов (ЭВМ)
    arrival_rate = 0.25  # интенсивность поступления заявок (лямбда)
    service_rate = 1 / 3  # интенсивность обслуживания заявок (1/t_об)

    # Расчёт интенсивности нагрузки
    rho = arrival_rate / service_rate  # интенсивность нагрузки

    # Предельные вероятности состояний
    p0 = 1 / sum(rho ** k / math.factorial(k) for k in range(n + 1))
    p1 = rho * p0
    p2 = (rho ** 2 / math.factorial(2)) * p0
    p3 = (rho ** 3 / math.factorial(3)) * p0

    # Вероятность отказа
    P_otk = p3

    # Относительная пропускная способность
    Q = 1 - P_otk

    # Абсолютная пропускная способность
    A = arrival_rate * Q

    # Среднее число занятых каналов
    avg_busy = A / service_rate

    # Средняя занятость одной ЭВМ
    avg_utilization = avg_busy / n * 100

    # Вывод результатов
    print(f"Интенсивность нагрузки: ρ = {rho:.2f}")
    print(f"Предельные вероятности состояний: p0 = {p0:.3f}, p1 = {p1:.3f}, p2 = {p2:.3f}, p3 = {p3:.3f}")
    print(f"Вероятность отказа: P_otk = {P_otk:.3f}")
    print(f"Относительная пропускная способность: Q = {Q:.3f}")
    print(f"Абсолютная пропускная способность: A = {A:.3f}")
    print(f"Среднее число занятых ЭВМ: avg_busy = {avg_busy:.3f}")
    print(f"Средняя занятость одной ЭВМ: avg_utilization = {avg_utilization:.1f}%")


primer7()
