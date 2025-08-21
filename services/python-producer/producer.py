from kafka import KafkaProducer
import time

producer = KafkaProducer(bootstrap_servers='localhost:9092')

topic = "golang-logs"

for i in range(10):
    message = f"log message {i}"
    producer.send(topic, value=message.encode("utf-8"))
    print(f"Sent: {message}")
    time.sleep(1)

producer.flush()
producer.close()
