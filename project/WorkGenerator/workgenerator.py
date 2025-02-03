from confluent_kafka import Producer
import json
import time
import pymongo

# Configurações
KAFKA_BROKER = 'kafka:9092'
TOPIC_NAME = 'simulation_tasks'
MONGO_URI = 'mongodb://mongo:27017/'
DATABASE_NAME = 'simulation_db'
COLLECTION_NAME = 'results'

# Inicializa o produtor Kafka
producer = Producer({'bootstrap.servers': KAFKA_BROKER})

def generate_tasks(num_tasks):
    """Gera tarefas e as publica no Kafka."""
    for i in range(num_tasks):
        task = {
            'task_id': i,
            'data': f"Task data {i}"
        }
        producer.produce(TOPIC_NAME, key='task_group_1', value=json.dumps(task))
        print(f"Task {i} published to Kafka.")
    producer.flush()

def fetch_results():
    """Busca os resultados no MongoDB até que todas as tarefas sejam concluídas."""
    client = pymongo.MongoClient(MONGO_URI)
    db = client[DATABASE_NAME]
    collection = db[COLLECTION_NAME]
    
    while True:
        results = list(collection.find())
        print(f"Current results: {len(results)}/{num_tasks}")
        if len(results) == num_tasks:
            print("All tasks completed:", results)
            break
        time.sleep(5)

if __name__ == "__main__":
    num_tasks = 100
    generate_tasks(num_tasks)
    fetch_results()