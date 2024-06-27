import matplotlib.pyplot as plt
import pandas as pd
import os

def load_data_from_directory(directory_path):
    data = {}
    for filename in os.listdir(directory_path):
        if filename.endswith(".csv"):
            file_path = os.path.join(directory_path, filename)
            dataset_name = filename[:-4]  
            df = pd.read_csv(file_path, header=None, names=['Fitness', 'FitnessCount'])
            data[dataset_name] = df
    return data

def plot_data(data):
    plt.figure(figsize=(10, 6))
    for dataset_name, df in data.items():
        plt.plot(df['FitnessCount'], df['Fitness'], marker='o', linestyle='-', label=dataset_name)
    plt.title('Fitness over FitnessCount')
    plt.xlabel('FitnessCount')
    plt.ylabel('Fitness')
    plt.legend()
    plt.grid(True)
    plt.show()

directory_path = "path/to/your/csv/files"

data = load_data_from_directory(directory_path)

plot_data(data)
