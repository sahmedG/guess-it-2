import numpy as np
import sys
from sklearn.linear_model import LinearRegression

class Prediction:
    def __init__(self, lower_limit, upper_limit):
        self.lower_limit = lower_limit
        self.upper_limit = upper_limit


def calculate_range(numbers):
    if len(numbers) < 2:
        return Prediction(numbers[0] - 1, numbers[0] + 1)
    x = np.arange(len(numbers)).reshape(-1, 1)
    y = np.array(numbers)

    model = LinearRegression().fit(x, y)

    next_x = np.array([[len(numbers)]])
    lower_limit = model.predict(next_x - 3)
    upper_limit = model.predict(next_x + 1)
    return Prediction(lower_limit, upper_limit)

def main():
    numbers = []
    old_prediction = None

    for line in sys.stdin:
        try:
            number = float(line.strip())
            if old_prediction is not None and number >= old_prediction.lower_limit and number <= old_prediction.upper_limit:
                prediction = old_prediction
            else:
                numbers.append(number)
                prediction = calculate_range(numbers)
                old_prediction = prediction

            print(int((prediction.lower_limit)),
                  int((prediction.upper_limit)))
        except ValueError:
            print("Invalid input. Please enter a valid number.")


if __name__ == "__main__":
    main()
