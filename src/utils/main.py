# main.py

import argparse
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score
from sklearn.model_selection import GridSearchCV
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.pipeline import Pipeline
from sklearn.linear_model import LogisticRegression
from sklearn.naive_bayes import MultinomialNB

def train_model(X_train, y_train):
    pipeline = Pipeline([
        ('vectorizer', TfidfVectorizer()),
        ('classifier', MultinomialNB())
    ])

    parameters = {
        'classifier__alpha': [0.1, 0.5, 1.0],
        'classifier__fit_prior': [True, False]
    }

    grid_search = GridSearchCV(pipeline, parameters, cv=5, n_jobs=-1, scoring='f1_macro')
    grid_search.fit(X_train, y_train)

    return grid_search.best_estimator_

def main():
    parser = argparse.ArgumentParser(description='Train a text classification model')
    parser.add_argument('input_file', help='Path to input file')
    parser.add_argument('output_file', help='Path to output file')
    args = parser.parse_args()

    # Load data
    X, y = load_data(args.input_file)

    # Split data
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # Train model
    model = train_model(X_train, y_train)
    y_pred = model.predict(X_test)

    # Evaluate model
    accuracy = accuracy_score(y_test, y_pred)
    print(f'Model accuracy: {accuracy:.2f}')

    # Save model
    import pickle
    with open(args.output_file, 'wb') as f:
        pickle.dump(model, f)

def load_data(file_path):
    # Load data from file
    import pandas as pd
    return pd.read_csv(file_path, sep='\t', header=0).values[:, 0], pd.read_csv(file_path, sep='\t', header=0).values[:, 1]

if __name__ == '__main__':
    main()