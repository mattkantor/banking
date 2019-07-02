import file
import json
import dict

# read in a line
# check the limits for each of the last iterations
# convert to UTC
# - in the last day / check day of week
# if true, add to the hash and deduct the balance
# if false, return false, log the event?
# compare the logs
# {"id":"15887","customer_id":"528","load_amount":"$3318.47","time":"2000-01-01T00:00:00Z"}

INPUT_FILE = "../input.txt"
log = {}







def write_to_log(line):
    id = line["id"]
    date = line["time"]
    logline = dict(id=id, time=time)
    #deddup
    log.append(logline)




def ingest():
    '''ingest the file'''
    d = {}
    with open(INPUT_FILE) as f:
        for line in f:
        json_line = line.split()
        json_data = json.loads(json_line)
        d[int(key)] = val


def main():
    ingest()