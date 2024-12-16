def zadanie_8():
    # Данные задачи
    total_lamps = 1000  # Общее количество ламп
    failed_3000_4000 = 50  # Число отказов за интервал 3000-4000 часов
    time_interval = 1000  # Интервал времени (4000 - 3000)

    # Расчёт частоты отказов
    failure_frequency = failed_3000_4000 / (total_lamps * time_interval)

    # Расчёт среднего числа работающих ламп на интервале
    average_working_lamps = total_lamps - failed_3000_4000 / 2

    # Расчёт интенсивности отказов
    failure_intensity = failed_3000_4000 / (time_interval * average_working_lamps)

    # Вывод результатов
    print(f"Частота отказов (q): {failure_frequency:.6f} отказов/час")
    print(f"Интенсивность отказов (λ): {failure_intensity:.6f} отказов/час")


def zadanie_36():
    # Данные задачи
    P_s = 0.999  # Вероятность безотказной работы системы
    n = 5  # Количество подсистем

    # Вычисление вероятности безотказной работы одной подсистемы
    P_e = P_s ** (1 / n)

    # Вывод результата
    print(f"Минимально допустимая вероятность безотказной работы элемента: {P_e:.6f}")


def zadanie_62():
    import numpy as np
    from scipy.stats import t, chi2

    # Данные выборки
    data = [125, 85, 80, 250, 85, 85, 160, 80]
    n = len(data)
    alpha = 0.95  # Уровень доверия

    # Вычисляем среднее T и стандартное отклонение σ
    T = np.mean(data)
    sigma = np.std(data, ddof=1)

    # Доверительный интервал для T
    t_critical = t.ppf(1 - (1 - alpha) / 2, df=n - 1)
    T_low = T - t_critical * sigma / np.sqrt(n)
    T_high = T + t_critical * sigma / np.sqrt(n)

    # Доверительный интервал для σ
    chi2_low = chi2.ppf((1 - alpha) / 2, df=n - 1)
    chi2_high = chi2.ppf(1 - (1 - alpha) / 2, df=n - 1)
    sigma_low = np.sqrt((n - 1) * sigma ** 2 / chi2_high)
    sigma_high = np.sqrt((n - 1) * sigma ** 2 / chi2_low)

    # Результаты
    print(f"Среднее значение T: {T:.2f}")
    print(f"Стандартное отклонение σ: {sigma:.2f}")
    print(f"Доверительный интервал для T: ({T_low:.2f}, {T_high:.2f})")
    print(f"Доверительный интервал для σ: ({sigma_low:.2f}, {sigma_high:.2f})")


print("Задание 8:")
zadanie_8()
print("\n\n\nЗадание 36:")
zadanie_36()
print("\n\n\nЗадание 62:")
zadanie_62()