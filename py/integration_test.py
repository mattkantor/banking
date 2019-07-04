import sys

from db import *
from util import load_json_lines

RESULTS_FILE = '../results.txt'

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)
f_handler = logging.FileHandler(RESULTS_FILE)
f_handler.setLevel(logging.INFO)
logger.addHandler(f_handler)

INPUT_FILE = "../input.txt"
OUTPUT_FILE = "../output.txt"


def test_all():

    '''
    The integration test runs agains the input file and compares against the output file, while writing its own output results.txt file

    '''

    def process_item(data):
        db = CustomerDatabase()
        success, code = db.deposit(data["customer_id"], data["id"], data["time"], data["load_amount"])

        if code != 403:
            data = write_output(data["customer_id"], data["id"], success)
        return data, code

    def write_output(customer_id, txn_id, status):
        data = dict(id=txn_id, customer_id=customer_id, accepted=status)
        logger.info(json.dumps(data))
        return data

    with open(RESULTS_FILE, 'w'):
        pass
    todos = load_json_lines(INPUT_FILE)
    answers = load_json_lines(OUTPUT_FILE)
    counter = 0
    for to_do in todos:
        result, code = process_item(to_do)
        if  code != 403:
            assert result["accepted"] == answers[counter]["accepted"]
            counter += 1


if __name__ == "__main__":
    # execute only if run as a script
    test_all()