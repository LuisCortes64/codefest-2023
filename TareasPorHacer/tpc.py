case_1_k = "1"
case_1_list = [3, -1, -2]


def tareasPorHacer(tareas: list, k: int):
    for i in range(len(tareas)):
        if tareas[i] > 0:
            tareas[i], tareas[i+k] = tareas[i] + tareas[i+k], 0
    return sum(abs(n) for n in tareas)


print(tareasPorHacer(case_1_list, case_1_k))
