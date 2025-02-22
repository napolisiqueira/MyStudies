

N = int(input())

for p in range(0, N):
    x, y = map(float, input().split(" "))

    if y == 0:
        print("divisao impossivel")
    else:
        print(x/y)
        

    