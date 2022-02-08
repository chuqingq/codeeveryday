from scipy.optimize import fsolve
def func(x):
    x0,x1,x2 = x.tolist()
    return [5*x1-25,5*x0*x0-x1*x2,x2*x0-27]

initial_x = [1,1,1]
result = fsolve(func, initial_x)
print(result)
