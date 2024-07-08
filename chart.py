import matplotlib.pyplot as plt
import pandas as pd 
import seaborn as sns

def create_equality_compares_graph(filepath: str):
    df = pd.read_csv(filepath)

    df1 = df[df.columns.difference(['key', 'cumulative_avg'])]
    sns.scatterplot(data = df1, x = "s_no", y = "num_equality_compares", s=5, label="number of equality comparisons")

    df2 = df[df.columns.difference(['key', 'num_equality_comparisons'])]
    sns.scatterplot(data = df2, x = "s_no", y = "cumulative_avg", s=1, label="cumulative average of equality comparisons")

    plt.legend()
    plt.xlabel("Operation No.")
    plt.ylabel("Number of equality comparisons")

    plt.show()


if __name__=='__main__':
    create_equality_compares_graph('./out/tale_linear.csv')
    create_equality_compares_graph('./out/tale_chain.csv')