import matplotlib.pyplot as plt
import math


def pr1(R, E, C, a=0, b=1000):
    i = []

    for t in range(a, b):
        expu = math.exp(-(1 / (R * C)) * t)
        i.append((E / R) * expu)
        '''i.append(E/R)'''

    plt.plot(range(a, b), i)
    plt.show()


def pr2(R, E, L, a=0, b=1000):
    u = []

    for t in range(a, b):
        expu = math.exp(-1 * (R / L) * t)
        u.append(E * expu)
        '''i.append(E/R)'''

    plt.plot(range(a, b), u)
    plt.show()


def pr3(R1, R2, E, C, L, a=0, b=1000):
    h = 1 - (4 * C * L * R2 * (R1 + R2))
    o = ((E * R2) / (C * L * (R1 + R2))) * (1 + ((L + R1 * R2 * C) / (2 * C * L * R2)) * (
            1 + (h / (L + R1 * R2 * C))))
    l = 2 * h
    a2 = o / l
    a1 = -a2 - (E * R) / (C * L * (R1 + R2))

    u = []

    for t in range(a, b):
        z = a1 * math.exp(-((L + R1 * R2 * C) / (2 * C * L * R2)) * (1 + h) * t)
        x = a2 * math.exp(-((L + R1 * R2 * C) / (2 * C * L * R2)) * (1 - h) * t)
        c = (E * R2) / (C * L * (R1 + R2))
        u.append(z + x + c)
        '''i.append(E/R)'''

    plt.plot(range(a, b), u)
    plt.grid()
    plt.show()


# Запуск 1
R = 2000  # Сопротивление (На 1 резисторе)
E = 12  # Напряжение
C = 0.15  # Ёмкость
a = 0  # Начало времени
b = 1000  # Конец времени
pr1(R, E, C, a, b)

# Запуск 2
# R = 100  # Сопротивление (На 1 резисторе)
# E = 12.0  # Напряжение
# L = 0.01  # Индуктивность
# a = 0  # Начало времени
# b = 3  # Конец времени
# pr2(R, E, L, a, b)

# Запуск 3
# R = 100  # Сопротивление (На 1 резисторе)
# R2 = 300  # Сопротивление (На 2 резисторе)
# E = 12.0  # Напряжение
# C = 10  # Ёмкость
# L = 0.01  # Индуктивность
# a = 0  # Начало времени
# b = 3  # Конец времени
# pr3(R, R2, E, C, L, a, b)
