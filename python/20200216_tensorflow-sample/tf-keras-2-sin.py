# coding:utf-8

import tensorflow as tf
from tensorflow import keras
from tensorflow.keras import layers

import numpy as np

# data 自己写个简单的线性
x_data = np.random.random((6000,1))
y_data = np.sin(x_data) * 2.5 + 0.8

# model
model = tf.keras.Sequential([
    # Adds a densely-connected layer with 64 units to the model:
    layers.Dense(64, activation='relu', input_shape=(1,)),
    # Add another:
    layers.Dense(64),
    layers.Dense(64),
    layers.Dense(64, activation='relu'),
    # Add an output layer with 10 output units:
    layers.Dense(1)
])

# compile
model.compile(optimizer=tf.keras.optimizers.Adam(0.01),
              #loss=tf.keras.losses.CategoricalCrossentropy(from_logits=True),
              loss='mse',
              metrics=['accuracy'])

# fit
model.fit(x_data, y_data, epochs=16, batch_size=32)

model.save('my_model.h5')

# predict 预计结果是 0.5*2.5+0.8=2.05
data = [[0.5]]
res = model.predict(data)
print(res)
print('expected: ', np.sin(data)*2.5+0.8)
