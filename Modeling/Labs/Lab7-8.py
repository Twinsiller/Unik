import numpy as np
import random as rm

'''states = ["Sleep", "Run", "Icecream"]
transitionName = [["SS", "SR", "SI"], ["RS", "RR", "RI"], ["IS", "IR", "II"]]
transitionMatrix = [[0.2, 0.6, 0.2], [0.1, 0.6, 0.3], [0.2, 0.7, 0.1]]
if sum(transitionMatrix[0]) + sum(transitionMatrix[1]) + sum(transitionMatrix[2]) != 3:
    print("Somewhere, something went wrong. Transition matrix, perhaps?")
else:
    print("All is gonna be okay, you should move on!! ;)")


# 实现了可以预测状态的马尔可夫模型的函数。
def activity_forecast(days):
    activityToday = "Sleep"
    print("Start state: " + activityToday)
    # 应该记录选择的状态序列。这里现在只有初始状态。
    activityList = [activityToday]
    i = 0
    # 计算 activityList 的概率
    prob = 1
    while i != days:
        if activityToday == "Sleep":
            change = np.random.choice(transitionName[0],replace=True,p=transitionMatrix[0])
            if change == "SS":
                prob = prob * 0.2
                activityList.append("Sleep")
                pass
            elif change == "SR":
                prob = prob * 0.6
                activityToday = "Run"
                activityList.append("Run")
            else:
                prob = prob * 0.2
                activityToday = "Icecream"
                activityList.append("Icecream")
        elif activityToday == "Run":
            change = np.random.choice(transitionName[1],replace=True,p=transitionMatrix[1])
            if change == "RR":
                prob = prob * 0.5
                activityList.append("Run")
                pass
            elif change == "RS":
                prob = prob * 0.2
                activityToday = "Sleep"
                activityList.append("Sleep")
            else:
                prob = prob * 0.3
                activityToday = "Icecream"
                activityList.append("Icecream")
        elif activityToday == "Icecream":
            change = np.random.choice(transitionName[2],replace=True,p=transitionMatrix[2])
            if change == "II":
                prob = prob * 0.1
                activityList.append("Icecream")
                pass
            elif change == "IS":
                prob = prob * 0.2
                activityToday = "Sleep"
                activityList.append("Sleep")
            else:
                prob = prob * 0.7
                activityToday = "Run"
                activityList.append("Run")
        i += 1
    print("Possible states: " + str(activityList))
    print("End state after "+ str(days) + " days: " + activityToday)
    print("Probability of the possible sequence of states: " + str(prob))

## 预测 2 天后的可能状态
#activity_forecast(2)'''


def zadanie5_1():
    times = 2
    start = np.array([0.1, 0.0, 0.3, 0.15, 0.25, 0.2])
    matricaPerhoda = np.array([
        [0.0, 0.0, 0.2, 0.1, 0.0, 0.7],
        [0.1, 0.0, 0.0, 0.0, 0.0, 0.9],
        [0.1, 0.0, 0.0, 0.0, 0.2, 0.7],
        [0.2, 0.0, 0.2, 0.1, 0.0, 0.5],
        [0.1, 0.2, 0.1, 0.1, 0.0, 0.5],
        [0.1, 0.0, 0.2, 0.1, 0.2, 0.4]
    ])
    # check = [[0.4,0.6],[0.3,0.7]] #Возведение в степень работет и без numpy для массива из списка
    matricaPerhodaForTimes = np.linalg.matrix_power(matricaPerhoda, times)
    #print(matricaPerhodaForTimes)
    end = start.dot(matricaPerhodaForTimes)
    print(end)
    print(sum(end))


zadanie5_1()
