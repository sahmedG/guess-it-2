import sys

class Prediction:
    def __init__(self, lower_limit, upper_limit):
        self.lower_limit = lower_limit
        self.upper_limit = upper_limit


def calculate_range(numbers):
    if len(numbers) < 2:
        return Prediction(numbers[0] - 1, numbers[0] + 1)
    x = list(range(len(numbers)))
    sum_x = sum(x)
    sum_y = sum(numbers)
    sum_xy = sum(x[i] * numbers[i] for i in range(len(numbers)))
    sum_x2 = sum(x[i] ** 2 for i in range(len(numbers)))

    n = len(numbers)

    b = (n * sum_xy - sum_x * sum_y) / (n * sum_x2 - sum_x ** 2)
    a = (sum_y - b * sum_x) / n

    next_x = len(numbers)
    lower_limit = a + b * (next_x + 0.5)
    upper_limit = a + b * (next_x + 1.5)

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

            print(int(round(prediction.lower_limit)),
                  int(round(prediction.upper_limit)))
        except ValueError:
            print("Invalid input. Please enter a valid number.")


if __name__ == "__main__":
    main()
