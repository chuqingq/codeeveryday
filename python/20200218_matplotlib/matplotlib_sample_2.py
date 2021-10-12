import numpy as np 
from matplotlib import pyplot as plt 
 
x = np.linspace(0, 10) 
y = np.sin(x)
plt.title("Matplotlib demo") 
plt.xlabel("x axis caption") 
plt.ylabel("y axis caption") 
plt.plot(x,y)
plt.show()
plt.savefig('1.png')

