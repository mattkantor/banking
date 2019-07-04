import sys

from db import *


RESULTS_FILE = '../results.txt'

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)
f_handler = logging.FileHandler(RESULTS_FILE)
f_handler.setLevel(logging.INFO)
logger.addHandler(f_handler)

INPUT_FILE = "../input.txt"
OUTPUT_FILE = "../output.txt"

def process_item(data):
    db = CustomerDatabase()
    success , code = db.deposit(data["customer_id"], data["id"], data["time"], data["load_amount"])
    print(success, code)
    if  code != 403:
        data = write_output(data["customer_id"], data["id"], success)
    return data, code

def write_output(customer_id, txn_id, status):
    data = dict(id=txn_id, customer_id=customer_id, accepted=status)
    #logger.info(json.dumps(data))
    return data

def load_json_lines(filename):
    data_arr  = []
    with open(filename) as f:
        for line in f:
            json_dict = json.loads(line)
            data_arr.append(json_dict)
    return data_arr

def ingest():
    bad_counter = 0
    with open(RESULTS_FILE, 'w'):
        pass
    todos = load_json_lines(INPUT_FILE)
    answers = load_json_lines(OUTPUT_FILE)
    counter = 0
    for to_do in todos:
        result, code = process_item(to_do)

        if  code != 403:
            if result["accepted"] != answers[counter]["accepted"]:
                print(result, answers[counter],  "At", str(counter) , "for", to_do["load_amount"])
                bad_counter +=1
            counter += 1
        else:
            print("403")
    print("BAD COUNTER=" , str(bad_counter))



if __name__ == "__main__":
    # execute only if run as a script
    ingest()