def n1():
    from cvxopt.modeling import variable, op
    import time

    start = time.time()
    x = variable(12, 'x')
    c = [100, 90, 110, 75.5, 80, 95, 80, 85, 100, 70, 95, 105]
    z = (c[0] * x[0] + c[1] * x[1] + c[2] * x[2] + c[3] * x[3] + c[4] * x[4] + c[5] * x[5] + c[6] * x[6] + c[7] * x[7] +
         c[
             8] * x[8] + c[9] * x[9] + c[10] * x[10] + c[11] * x[11])
    mass1 = (x[0] + x[1] + x[2] == 155.5)
    mass2 = (x[3] + x[4] + x[5] == 250)
    mass3 = (x[6] + x[7] + x[8] == 170)
    mass4 = (x[9] + x[10] + x[11] == 134.5)
    mass5 = (x[0] + x[3] + x[6] + x[9] <= 350)
    mass6 = (x[1] + x[4] + x[7] + x[10] <= 250)
    mass7 = (x[2] + x[5] + x[8] + x[11] <= 400)
    x_non_negative = (x >= 0)
    problem = op(z, [mass1, mass2, mass3, mass4, mass5, mass6, mass7, x_non_negative])
    problem.solve(solver='glpk')
    print("Результат Xopt:")

    for i in x.value:
        print(i)
    print("Стоимость доставки:")
    print(problem.objective.value()[0])
    stop = time.time()
    print("Время :")
    print(stop - start)


def n2():
    import numpy as np
    from python_tsp.exact import solve_tsp_dynamic_programming

    distance_matrix = np.array([
        [0, 15, 7, 10, 9, 21, 5, 11],
        [17, 0, 10, 15, 7, 12, 6, 9],
        [11, 9, 0, 13, 25, 14, 8, 10],
        [12, 7, 13, 0, 21, 24, 10, 17],
        [23, 8, 9, 13, 0, 15, 21, 16],
        [17, 21, 8, 11, 13, 0, 10, 14],
        [9, 11, 20, 15, 10, 17, 0, 8],
        [7, 12, 17, 10, 9, 11, 22, 0]
    ])
    permutation, distance = solve_tsp_dynamic_programming(distance_matrix)
    print(permutation, distance)
