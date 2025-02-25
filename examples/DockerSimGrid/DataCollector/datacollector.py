import pymongo
import requests

MONGO_URI = 'mongodb://mongo:27017/'
DATABASE_NAME = 'simulation_db'
COLLECTION_NAME = 'results'
MASTER_NODE_URL = 'http://masternode:8080/end-simulation'

def collect_data():
    # Simulate data collection
    data = {
        'task_id': 1,
        'result': 'Simulation Result'
    }

    client = pymongo.MongoClient(MONGO_URI)
    db = client[DATABASE_NAME]
    collection = db[COLLECTION_NAME]
    collection.insert_one(data)

    # Notify MasterNode
    requests.post(MASTER_NODE_URL, json={'task_id': data['task_id']})

if __name__ == "__main__":
    collect_data()