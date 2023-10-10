import numpy as np
import sys
from sklearn.linear_model import Ridge

x_values = []
y_values = []

class Prediction:
    def __init__(self, lower_limit, upper_limit):
        self.lower_limit = lower_limit
        self.upper_limit = upper_limit

def calculate_range(x_values, y_values, number):
    x_values = np.array(x_values).reshape(-1, 1)
    y_values = np.array(y_values)

    if len(x_values) < 2:
        return

    ridge = Ridge(alpha=1.0)
    ridge.fit(x_values, y_values)

    next_x = np.array([[len(x_values) + 1]])
    prediction = ridge.predict(next_x)
    temp_low = prediction[0]
    temp_high = prediction[0]
    print(round(temp_low), round(temp_high))


for line in sys.stdin:
    try:
        number = int(line.strip())

        if number > 0:
            x_values.append(len(x_values))
            y_values.append(number)
            calculate_range(x_values, y_values, number)
    except ValueError:
        print("Invalid input. Please enter a valid number.")
    except EOFError:
        break
