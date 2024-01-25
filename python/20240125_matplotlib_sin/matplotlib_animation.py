import numpy as np
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation

#生成子图，相当于fig = plt.figure(),ax = fig.add_subplot(),其中ax的函数参数表示把当前画布进行分割，例：fig.add_subplot(2,2,2).表示将画布分割为两行两列
fig, ax = plt.subplots()

#ax在第2个子图中绘制，其中行优先，
#初始化两个数组
xdata, ydata = [], []
#第三个参数表示画曲线的颜色和线型，具体参见：https://blog.csdn.net/tengqingyong/article/details/78829596
ln, = ax.plot([], [], 'r-', animated=False)

def init():
    #设置x轴的范围pi代表3.14...圆周率，
    ax.set_xlim(0, 2*np.pi)
    #设置y轴的范围
    ax.set_ylim(-1, 1)
    #返回曲线
    return ln,

def update(n):
    #将每次传过来的n追加到xdata中
    xdata.append(n)
    ydata.append(np.sin(n))
    #重新设置曲线的值
    ln.set_data(xdata, ydata)
    return ln,

#这里的frames在调用update函数是会将frames作为实参传递给“n”
ani = FuncAnimation(fig, update, frames=np.linspace(0, 2*np.pi, 10),
                    init_func=init, blit=True)
plt.show()