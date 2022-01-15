import numpy as np
from scipy.optimize import curve_fit

def func(x, a, b):
    return a * x + b

x = np.linspace(0, 10, 100)
y = func(x, 1, 2)

yn = y + 0.9 * np.random.normal(size=len(x))

popt, pconv = curve_fit(func, x, yn)
print(popt)
