
from db import *

RESULTS_FILE = '../results_py.txt'
INPUT_FILE = "../input.txt"
OUTPUT_FILE = "../output.txt"

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)
f_handler = logging.FileHandler(RESULTS_FILE)
f_handler.setLevel(logging.DEBUG)
logger.addHandler(f_handler)




def process_item(data):

    db = CustomerDatabase()
    success, code = db.deposit(data["customer_id"], data["id"], data["time"], data["load_amount"])

    if code != 403:
        data = write_output(data["customer_id"], data["id"], success)
    return data, code

def write_output(customer_id, txn_id, status):
    data = dict(id=txn_id, customer_id=customer_id, accepted=status)

    return data

def load_json_lines(filename):
    data_arr  = []
    with open(filename) as f:
        for line in f:
            json_dict = json.loads(line)
            data_arr.append(json_dict)
    return data_arr

def test_all():
    bad_counter = 0
    with open(RESULTS_FILE, 'w') as outfile:

        todos = load_json_lines(INPUT_FILE)
        answers = load_json_lines(OUTPUT_FILE)
        counter = 0

        for to_do in todos:
            result, code = process_item(to_do)
            outfile.write(json.dumps(result))
            if  code != 403:
                assert result["accepted"] == answers[counter]["accepted"]
                counter += 1




if __name__ == "__main__":
    # execute only if run as a script
    test_all()