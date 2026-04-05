# main.py
import logging
import os
from typing import Tuple

import numpy as np
from tensorflow import keras
from tensorflow.keras import layers
import matplotlib.pyplot as plt

from models import build_model
from utils import load_dataset

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def main() -> None:
    logger.info("Loading dataset")
    train_data, test_data = load_dataset()

    logger.info("Building model")
    model = build_model(input_shape=train_data.shape[1:])

    logger.info("Compiling model")
    model.compile(optimizer='adam',
                loss='mean_squared_error',
                metrics=['accuracy'])

    logger.info("Training model")
    history = model.fit(train_data,
                        epochs=10,
                        validation_data=test_data,
                        verbose=2)

    logger.info("Evaluating model")
    loss, accuracy = model.evaluate(test_data)
    logger.info(f"Test loss: {loss:.2f}, Test accuracy: {accuracy:.2f}")

    logger.info("Plotting results")
    plt.plot(history.history['accuracy'], label='Training Accuracy')
    plt.plot(history.history['val_accuracy'], label='Validation Accuracy')
    plt.legend()
    plt.show()

if __name__ == "__main__":
    main()