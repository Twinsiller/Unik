import numpy as np
import matplotlib.pyplot as plt
from scipy.optimize import minimize_scalar


def primer1():
    # Функция для вычисления объема ведра (конуса)
    def volume(theta, R):
        r = (R * theta) / (2 * np.pi)  # Радиус основания конуса
        h = np.sqrt(R**2 - r**2)       # Высота конуса
        return (1/3) * np.pi * r**2 * h  # Объем конуса

    # Найдем угол theta, при котором объем максимален
    def find_max_volume(R):
        # Оптимизация объема по углу theta
        result = minimize_scalar(lambda theta: -volume(theta, R), bounds=(0, 2 * np.pi), method='bounded')
        return result.x, -result.fun

    # Радиус исходного круга
    R = 1  # например, 1 метр

    # Найдем оптимальный угол и максимальный объем
    optimal_theta, max_vol = find_max_volume(R)

    # Вывод результатов
    print(f"Оптимальный угол вырезания сектора: {np.degrees(optimal_theta):.2f} градусов")
    print(f"Максимальный объем ведра: {max_vol:.4f} кубических метров")

    # Построение графика объема в зависимости от угла вырезания
    thetas = np.linspace(0, 2*np.pi, 100)
    volumes = [volume(theta, R) for theta in thetas]

    plt.plot(np.degrees(thetas), volumes)
    plt.xlabel("Угол вырезания сектора (градусы)")
    plt.ylabel("Объем ведра (кубические метры)")
    plt.title("Зависимость объема ведра от угла вырезания сектора")
    plt.grid(True)
    plt.show()

def primer2():
    from scipy.optimize import linprog

    # Коэффициенты функции цели (прибыли)
    c = [-1500, -2100]  # Минимизируем отрицательную прибыль (т.к. linprog ищет минимум)

    # Коэффициенты ограничений
    A = [
        [50, 40],  # Ограничение по золоту
        [30, 50]  # Ограничение по платине
    ]

    # Пределы по ресурсам
    b = [1500, 1800]

    # Границы переменных (x1 >= 0 и x2 >= 0)
    x_bounds = (0, None)

    # Решение задачи линейного программирования
    res = linprog(c, A_ub=A, b_ub=b, bounds=[x_bounds, x_bounds], method='highs')

    # Результаты
    if res.success:
        print(f"Оптимальное количество часов 'Банкир': {res.x[0]:.2f}")
        print(f"Оптимальное количество часов 'Президент': {res.x[1]:.2f}")
        print(f"Максимальная прибыль: {-res.fun:.2f} долларов")
    else:
        print("Решение не найдено")


primer1()
#primer2()
