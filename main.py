import numpy as np
from sklearn.linear_model import LinearRegression

# Create an empty list to store the filtered numbers
numbers = []

# Function to filter numbers and perform linear regression
def predict_range(number):
    if 98 <= number <= 210:
        numbers.append(number)

    # Convert the list of numbers to a NumPy array
    X = np.array(numbers).reshape(-1, 1)

    # Fit a linear regression model to the data
    model = LinearRegression().fit(X, np.arange(len(numbers)).reshape(-1, 1))

    # Predict the lower and upper limits using the linear regression model
    lower_limit = model.predict([[0]])[0][0]
    upper_limit = model.predict([[len(numbers) - 1]])[0][0]

    return int(round(lower_limit)), int(round(upper_limit))

# Read numbers from standard input
while True:
    try:
        number = int(input())
        lower_limit, upper_limit = predict_range(number)
        print(lower_limit, upper_limit)
    except ValueError:
        break
