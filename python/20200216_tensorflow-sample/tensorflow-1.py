# coding:utf-8

# 调用tensorflow
import tensorflow as tf
import numpy as np

# 这里生成了100对数字，作为整个神经网络的input
x_data = np.random.rand(100).astype("float32")

# 使用with，让我们的数据以节点的方式落在tensorflow的报告上。
with tf.name_scope('y_data'):
    y_data = x_data * 2.5 + 0.8 #权重2.5，偏移设置2.5
    tf.summary.histogram("method_demo"+"/y_data",y_data) #可视化观看变量y_data


# 指定W和b变量的取值范围，随机在[-200,200]
with tf.name_scope('W'):
    W = tf.Variable(tf.random_uniform([1], -200.0, 200.0))
    tf.summary.histogram("method_demo"+"/W",W) #可视化观看变量

# 指定偏移值b，同时shape等于1
with tf.name_scope('b'):
    b = tf.Variable(tf.zeros([1]))
    tf.summary.histogram("method_demo"+"/b",b) #可视化观看变量

with tf.name_scope('y'):
    y = W * x_data + b #sigmoid神经元
    tf.summary.histogram("method_demo"+"/y",y) #可视化观看变量

# 最小化均方
with tf.name_scope('loss'):
    loss = tf.reduce_mean(tf.square(y - y_data))
    tf.summary.histogram("method_demo"+"/loss",loss) #可视化观看变量
    tf.scalar_summary("method_demo"+'loss',loss) #可视化观看常量

# 定义学习率，我们先使用0.7来看看效果
optimizer = tf.train.GradientDescentOptimizer(0.7)
with tf.name_scope('train'):
    train = optimizer.minimize(loss)

# 初始化TensorFlow参数
init = tf.initialize_all_variables()

# 运行数据流图
sess = tf.Session()
#合并到Summary中
merged = tf.merge_all_summaries() # TODO chuqq
#选定可视化存储目录
writer = tf.train.SummaryWriter(LOG_PATH,sess.graph)

sess.run(init)

# 开始计算
for step in xrange(500):
    sess.run(train)
    if step % 5 == 0:
        print(step, "W:",sess.run(W),"b:", sess.run(b))
        result = sess.run(merged) #merged也是需要run的
        writer.add_summary(result,step) #result是summary类型的 # TODO chuqq
