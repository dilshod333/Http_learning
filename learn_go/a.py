import math
import time

# Define some parameters
a = 1664525
c = 1013904223
m = 2**32

def pseudo_random(seed):
    seed = (a * seed + c) % m
    return seed

# Generating pseudo-random integers between 2 and 9
seed = int(time.time())  # Use current timestamp as seed
for _ in range(10):
    seed = pseudo_random(seed)
    random_integer = 2 + (seed % 8)  # Output integers between 2 and 9
    print(random_integer)
