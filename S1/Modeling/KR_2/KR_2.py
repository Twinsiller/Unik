# Вариант 2 - Болдинов Алексей ЭФМО-02-24 (6 по списку)
def zadanie1():
    from cvxopt.modeling import variable, op
    import time

    # start = time.time()
    x = variable(9, 'x')
    c = [6, 4, 2, 3, 5, 4, 3, 6, 3]
    z = (c[0] * x[0] + c[1] * x[1] + c[2] * x[2] + c[3] * x[3] + c[4] * x[4] + c[5] * x[5] + c[6] * x[6] + c[7] * x[7] +
         c[8] * x[8])
    mass1 = (x[0] + x[1] + x[2] == 24)
    mass2 = (x[3] + x[4] + x[5] == 28)
    mass3 = (x[6] + x[7] + x[8] == 23)
    mass4 = (x[0] + x[3] + x[6] == 20)
    mass5 = (x[1] + x[4] + x[7] == 25)
    mass6 = (x[2] + x[5] + x[8] == 30)
    x_non_negative = (x >= 0)
    problem = op(z, [mass1, mass2, mass3, mass4, mass5, mass6, x_non_negative])
    problem.solve(solver='glpk')
    print("Результат Xopt:")

    for i in x.value:
        print(i)
    print("Стоимость доставки:", problem.objective.value()[0])
    # stop = time.time()
    # print("Время :", stop - start)


def zadanie2():
    import numpy as np
    from python_tsp.exact import solve_tsp_dynamic_programming

    distance_matrix = np.array([
        [0, 24, 14, 20, 6],
        [17, 0, 8, 18, 11],
        [10, 30, 0, 26, 3],
        [2, 19, 22, 0, 1],
        [7, 12, 5, 15, 0]
    ])
    permutation, distance = solve_tsp_dynamic_programming(distance_matrix)
    print(permutation, distance)


zadanie2()
