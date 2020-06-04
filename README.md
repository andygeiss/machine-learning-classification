# Machine Learning with Golang (Classification)

This repository provides a Golang implementation from scratch to solve a classification problem using the K-Nearest Neighbour algorithm.

It can also be used as a standard project layout for machine learning projects, because it combines the
[Standard Go Project Layout](https://github.com/golang-standards/project-layout) and 
[Cookiecutter's Data Science - Directory structure](https://drivendata.github.io/cookiecutter-data-science/#directory-structure).

## Iris flower data set

First we need to gather the raw data from the external, [official web archive](https://archive.ics.uci.edu/ml/machine-learning-databases/iris).
Next we transform that CSV to an equal Protobuf format. Finally, we organize our final, processed Model format.

    data
    ├── external
    │   └── iris.csv
    ├── interim
    │   └── iris_interim.pb
    └── processed
        └── iris_processed.pb
        
But wait! Why Protobuf? Well, with Protobuf we are able to use the model within Golang, Python or many other languages.
We do not need any additional data type conversions to use the model in different programming languages.

Start the initial setup with the following commands:

    make

This will install the protobuf-compiler, corresponding protobuf-gen-go plugin generates the internal API and compiles all commands.

## Internal API

With the last step we generated our internal API by using the protobuf-compiler.
Our internal directory structure looks as follows:

    internal
    ├── api
    │   ├── iris_interim_pb2.py
    │   ├── iris_interim.pb.go
    │   ├── iris_interim.proto
    │   ├── iris_processed_pb2.py
    │   ├── iris_processed.pb.go
    │   └── iris_processed.proto

## Gather and Organize data

We treat  [data as immutable](https://drivendata.github.io/cookiecutter-data-science/#data-is-immutable). 
Thus, we will use a pipeline to separate each step of data manipulation/transformation.
Finally, we could start building our first pipeline to automatically gather, organize the data and print some common statistics to get the following output:

    ./gather_and_organize_data.bin
    
    Statistics:
       Column           Mean     Median   Mode     Minimum  Maximum  Range    Variance Std Dev 
       Petal length     5.84     5.80     5.00     4.30     7.90     3.60     0.68     0.83    
       Petal width      3.05     3.00     3.00     2.00     4.40     2.40     0.19     0.43    
       Sepal length     3.76     4.35     1.50     1.00     6.90     5.90     3.09     1.76    
       Sepal width      1.20     1.30     0.20     1.00     6.90     5.90     3.09     1.76    

The corresponding source of the command could be found [here](cmd/gather_and_organize_data/main.go).

## Evaluate the Model

The K-Nearest Neighbour is a lazy algorithm. It doesn't learn via training, it "memorizes" the training dataset instead.
Thus, we don't call the following step model training. We evaluate the parameter k and feature-combinations in the Iris flower data set.
In addition to that we use [Standardization](pkg/floats/scale.go) to scale the values down between 0 and 1 and
zero values are replaced by the [Mean](pkg/floats/central_tendency.go). 

    ./evaluate_model.bin
    
    Statistics:
       Column           Mean     Median   Mode     Minimum  Maximum  Range    Variance Std Dev 
       Petal length     0.43     0.42     0.19     0.00     1.00     1.00     0.05     0.23    
       Petal width      0.44     0.42     0.42     0.00     1.00     1.00     0.03     0.18    
       Sepal length     0.47     0.57     0.08     0.00     1.00     1.00     0.09     0.30    
       Sepal width      0.46     0.50     0.04     0.00     1.00     1.00     0.09     0.30    
    K-Nearest Neighbour with k = 3 :
       Petal Length/Width Accuracy: 72.00
       Sepal Length/Width Accuracy: 95.00
    Evaluation time: 11.786393ms

The corresponding source of the command could be found [here](cmd/evaluate_model/main.go).

Finally, we found out that the Sepal length and Sepal has a very high accuracy of 95.00%.

## Predict

The final model will be stored at <code>models/iris_knn.pb</code>.
To predict a single feature combination of Sepal length (x) and Sepal width (y) with K=3 use the following command:

    ./predict.bin -x 3 -y 4 -k 3
    
    K-Nearest Neighbour - K: 3, Given: [3 4], Predicted: Iris-setosa
    Prediction time: 152.338µs

The corresponding source of the command could be found [here](cmd/predict/main.go).

**UPDATE**: Prediction time reduced by ~20% using [Manhattan distance-calculation](pkg/floats/knn.go).

    ./predict.bin -x 3 -y 4 -k 3
   
    K-Nearest Neighbour - K: 3, Given: [3 4], Predicted: Iris-setosa
    Prediction time: 127.856µs
