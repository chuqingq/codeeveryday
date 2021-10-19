import numpy as np 
import copy
import random
from matplotlib import pyplot as plt 
 
x = np.arange(1,100) 
y = copy.deepcopy(x)
random.shuffle(y)

plt.title("Matplotlib demo") 
plt.xlabel("x axis caption") 
plt.ylabel("y axis caption") 
plt.scatter(x,y)
#plt.show()
plt.savefig('1.png')

