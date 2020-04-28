# coding:utf-8

import tensorflow as tf
from tensorflow import keras
from tensorflow.keras import layers

import numpy as np

model = tf.keras.models.load_model('my_model.h5')

# predict 预计结果是 0.5*2.5+0.8=2.05
data = [[0.5]]
res = model.predict(data)
print(res)
print('expected: ', np.sin(data)*2.5+0.8)
