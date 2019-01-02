# TensorFlow and tf.keras
import tensorflow as tf
from tensorflow import keras
import numpy as np
from tensorflow.keras import optimizers
import argparse
import sys
import os.path
import subprocess
import csv

THIS_DIR = os.path.abspath(os.path.dirname(__file__))
ARENA_BIN = os.path.join(THIS_DIR, 'go', 'src', 'main')
WORKDIR = os.path.join(THIS_DIR, 'workdir')


def NewModel():
  return keras.Sequential([
    keras.layers.Dense(64, input_shape=(27,), activation=tf.nn.relu, name="inputLayer"),
    keras.layers.Dense(64, activation=tf.nn.relu),
    keras.layers.Dense(1, activation=tf.nn.sigmoid, name="outputLayer"),
    ])

def ExportModel(model, path):
  builder = tf.saved_model.builder.SavedModelBuilder(path)
  builder.add_meta_graph_and_variables(keras.backend.get_session(),["tictactoe"])
  builder.save()

def SaveModel(model, path):
  model.save(path)

def LoadModel(path):
  return keras.models.load_model(path)

class Trainer(object):
  def __init__(self, workdir):
    self.workdir = workdir

  def IterPath(self, iter_num):
    return "{:}/{:04d}".format(self.workdir, iter_num)

  def SavedModelPath(self, iter_num):
    return "{:}/model.h5".format(self.IterPath(iter_num))

  def ExportModelPath(self, iter_num):
    return "{:}/model".format(self.IterPath(iter_num))

  def DataPath(self, iter_num):
    return "{:}/data.csv".format(self.IterPath(iter_num))

  def GenerateData(self, model_num, iter_num):
    cmd = [ARENA_BIN, '-n', '10000',
        '-player1', self.ExportModelPath(model_num),
        '-player2', self.ExportModelPath(model_num),
        '-p1_file', self.DataPath(iter_num)]
    ret = subprocess.run(cmd)
    if ret.returncode != 0:
      raise Exception("There was a problem")

  def LoadData(self, iter_num):
    data = np.genfromtxt(self.DataPath(iter_num), delimiter=',')
    x = data[:,:27]
    y = data[:,27:]
    return x, y

  def RunIter0(self, model=None):
    if not model:
      model = NewModel()
    self.CompileModel(model)
    ExportModel(model, self.ExportModelPath(0))
    SaveModel(model, self.SavedModelPath(0))

  def CompileModel(self, model):
    opt = keras.optimizers.Adam(lr=1e-1)
    model.compile(optimizer=opt, loss='binary_crossentropy')

  def TrainModel(self, model, x, y):
    self.CompileModel(model)
    model.fit(x, y, epochs=10, verbose=0)

  def RunIter(self, iter_num, model=None):
    if iter_num == 0:
      self.RunIter0()
      return

    # TODO: Figure out how to save and load models "correctly".
    keras.backend.clear_session()
    os.mkdir(self.IterPath(iter_num))
    self.GenerateData(iter_num-1, iter_num)
    model = LoadModel(self.SavedModelPath(iter_num-1))
    x, y = self.LoadData(iter_num)
    self.TrainModel(model, x, y)
    ExportModel(model, self.ExportModelPath(iter_num))
    SaveModel(model, self.SavedModelPath(iter_num))

def main():
  parser = argparse.ArgumentParser(description='Train "TicTacToe Zero".')
  parser.add_argument('-iter', type=int, default=10,
      help='How many iterations to run.')
  parser.add_argument('-workdir', type=str, default=WORKDIR,
      help='Where to store the intermediate data.')
  args = parser.parse_args()
  #model = NewModel()
  trainer = Trainer(args.workdir)
  for i in range(args.iter):
    print("Iter: ", i)
    trainer.RunIter(i)
  return 0

if __name__ == "__main__":
  sys.exit(main())
