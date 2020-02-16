# coding:utf-8

# 调用tensorflow
import tensorflow as tf
import numpy as np

# 这里生成了100对数字，作为整个神经网络的input
x_data = np.random.rand(100).astype("float32")
y_data = x_data * 2.5 + 0.8 #权重2.5，偏移设置2.5

W = tf.Variable(tf.random_uniform([1], -200.0, 200.0))
b = tf.Variable(tf.zeros([1]))

y = W * x_data + b

# 最小化均方
loss = tf.reduce_mean(tf.square(y - y_data))

# 定义学习率，我们先使用0.7来看看效果
optimizer = tf.train.GradientDescentOptimizer(0.7)
train = optimizer.minimize(loss)

# 初始化TensorFlow参数
init = tf.initialize_all_variables()

# 运行数据流图
sess = tf.Session()
#合并到Summary中
#选定可视化存储目录

sess.run(init)

# 开始计算
for step in range(500):
    sess.run(train)
    if step % 5 == 0:
        print(step, "W:",sess.run(W),"b:", sess.run(b))
