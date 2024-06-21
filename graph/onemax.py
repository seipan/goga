import matplotlib.pyplot as plt

# データセット1
data1 = [
    (4.000, 30), (7.000, 60), (10.000, 90), (11.000, 120), (12.000, 150),
    (12.000, 180), (12.000, 210), (13.000, 240), (15.000, 270), (16.000, 300),
    (17.000, 330), (18.000, 360), (19.000, 390), (20.000, 420), (20.000, 450),
    (20.000, 480), (20.000, 510), (20.000, 540), (20.000, 570), (20.000, 600),
    (20.000, 630), (20.000, 660), (20.000, 690), (20.000, 720), (20.000, 750),
    (20.000, 780), (20.000, 810), (20.000, 840), (20.000, 870), (20.000, 900),
    (20.000, 930), (20.000, 960), (20.000, 990), (20.000, 1020), (20.000, 1050),
    (20.000, 1080), (20.000, 1110), (20.000, 1140), (20.000, 1170), (20.000, 1200)
]

# データセット2
data2 = [
    (8.000, 30), (8.000, 60), (9.000, 90), (10.000, 120), (11.000, 150),
    (13.000, 180), (13.000, 210), (14.000, 240), (15.000, 270), (17.000, 300),
    (17.000, 330), (17.000, 360), (17.000, 390), (17.000, 420), (17.000, 450),
    (17.000, 480), (17.000, 510), (17.000, 540), (17.000, 570), (17.000, 600),
    (17.000, 630), (17.000, 660), (17.000, 690), (17.000, 720), (17.000, 750),
    (17.000, 780), (17.000, 810), (17.000, 840), (17.000, 870), (17.000, 900),
    (17.000, 930), (17.000, 960), (17.000, 990), (17.000, 1020), (17.000, 1050),
    (17.000, 1080), (17.000, 1110), (17.000, 1140), (17.000, 1170), (17.000, 1200)
]

# データセット3
data3 = [
    (8.000, 30), (9.000, 60), (10.000, 90), (11.000, 120), (14.000, 150),
    (15.000, 180), (15.000, 210), (16.000, 240), (16.000, 270), (16.000, 300),
    (16.000, 330), (17.000, 360), (18.000, 390), (19.000, 420), (19.000, 450),
    (19.000, 480), (19.000, 510), (20.000, 540), (20.000, 570), (20.000, 600),
    (20.000, 630), (20.000, 660), (20.000, 690), (20.000, 720), (20.000, 750),
    (20.000, 780), (20.000, 810), (20.000, 840), (20.000, 870), (20.000, 900),
    (20.000, 930), (20.000, 960), (20.000, 990), (20.000, 1020), (20.000, 1050),
    (20.000, 1080), (20.000, 1110), (20.000, 1140), (20.000, 1170), (20.000, 1200)
]

fitness_values1, fitness_counts1 = zip(*data1)
fitness_values2, fitness_counts2 = zip(*data2)
fitness_values3, fitness_counts3 = zip(*data3)

plt.figure(figsize=(10, 6))
plt.plot(fitness_counts1, fitness_values1, marker='o', linestyle='-', color='b', label='Dataset 1')
plt.plot(fitness_counts2, fitness_values2, marker='s', linestyle='--', color='r', label='Dataset 2')
plt.plot(fitness_counts3, fitness_values3, marker='^', linestyle='-.', color='g', label='Dataset 3')
plt.title('Fitness over FitnessCount')
plt.xlabel('FitnessCount')
plt.ylabel('Fitness')
plt.legend()
plt.grid(True)
plt.show()
