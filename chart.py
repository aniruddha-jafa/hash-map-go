import matplotlib.pyplot as plt
import pandas as pd 
import seaborn as sns


df = pd.read_csv('./out/tale_linear.csv')

df1 = df[df.columns.difference(['key', 'cumulative_avg'])]
sns.scatterplot(data = df1, x = "s_no", y = "num_equality_compares", s=5, label="number of equality compares")

df2 = df[df.columns.difference(['key', 'num_equality_compares'])]
sns.scatterplot(data = df2, x = "s_no", y = "cumulative_avg", s=1, label="cumulative average of equality compares")

#concatenated = pd.concat([df1.assign(dataset='set1'), df2.assign(dataset='set2')])
plt.legend()
plt.show()
