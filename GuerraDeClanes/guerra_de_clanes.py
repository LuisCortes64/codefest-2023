import itertools

case_1_target = [3, 1717, 2058]
case_1_clanes = [
    [7, [455, 634], [414, 774], [371, 377], [
        436, 557], [513, 753], [803, 496], [401, 809]],
    [7, [508, 731], [854, 556], [666, 521], [274, 987],
     [910, 418], [249, 1055], [807, 1640]],
    [5, [225, 230], [501, 503], [410, 427], [381, 570], [610, 755]]]


def orderMafiosos(target: list, clanes: list[list]):
    res = []

    def getMax(clan: list):
        ans = 0
        for i in range(1, clan[0]+2):
            for j in itertools.combinations(clan[1:], i):
                t = list(zip(*j))
                if sum(t[0]) == target[1] and sum(t[1]) == target[2]:
                    ans = i
        return ans

    for i in clanes:
        res.append(getMax(i))
    return res


print(orderMafiosos(case_1_target, case_1_clanes))
