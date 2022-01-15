'''
使用最小二乘法拟合直线
'''
import numpy as np
from scipy.optimize import leastsq
import matplotlib.pyplot as plt

#训练数据
Xi = np.array([8.19,2.72,6.39,8.71,4.7,2.66,3.78])
Yi = np.array([7.01,2.78,6.47,6.71,4.1,4.23,4.05])

#定义拟合函数形式
def func(p,x):
    k,b = p
    return k*x+b

#定义误差函数
def error(p,x,y,s):
    print(s)
    return func(p,x)-y

#随机给出参数的初始值
p = [10,2]

#使用leastsq()函数进行参数估计
s = '参数估计次数'
Para = leastsq(error,p,args=(Xi,Yi,s))
k,b = Para[0]
print('k=',k,'\n','b=',b)

#图形可视化
plt.figure(figsize = (8,6))
#绘制训练数据的散点图
plt.scatter(Xi,Yi,color='r',label='Sample Point',linewidths = 3)
plt.xlabel('x')
plt.ylabel('y')
x = np.linspace(0,10,1000)
y = k*x+b
plt.plot(x,y,color= 'orange',label = 'Fitting Line',linewidth = 2)
plt.legend()
plt.show()
