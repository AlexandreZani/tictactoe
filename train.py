# TensorFlow and tf.keras
import tensorflow as tf
from tensorflow import keras
import numpy as np
from tensorflow.keras import optimizers


def NewModel():
  return keras.Sequential([
    keras.layers.Dense(20, input_shape=(18,), activation=tf.nn.relu, name="inputLayer"),
    keras.layers.Dense(1, activation=tf.nn.sigmoid, name="outputLayer"),
    ])


def SaveModel(model, path):
  builder = tf.saved_model.builder.SavedModelBuilder(path)
  builder.add_meta_graph_and_variables(keras.backend.get_session(),["tictactoe"])
  builder.save()

def LoadData(path):
  return numpy.loadtxt(open(path, "rb"), delimiter=",")

model = NewModel()
print(model.predict(np.array([[0]*18])))

#SaveModel(model, "go/src/tictactoe/testdata/test_model")
