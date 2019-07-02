
from Collections import defaultdict
from util import *

customer = defaultdict()
WEEKLY_THRESHOLD = 20000
DAILY_THRESHOLD = 5000
TOTAL_LOADS_DAILY = 3

'''
customer{weekof:date, days_of_week:[{date:datetime, deposits:[]}]}

deposit_log:{key: log_id, value:{customer_id,  amount, successful}}

'''


def process_item(data):
    exists = check_for_dup(data["id"])
    if exists:
        return
    customer_id = data["customer_id"]

    customer_data = get_customer_data(customer_id)
    if customer_data == None:
        return

    future_deposit = data["amount"]
    per_week = get_total_deposits_for_week(customer_id)
    if per_week + future_deposit > WEEKLY_THRESHOLD:
        return False
    per_date = get_total_deposits_for_date(customer_id)
    if per_date + future_deposit > DAILY_THRESHOLD:
        return False
    per_amount = get_num_deposits_for_period(customer_id)
    if per_amount  > TOTAL_CENTS_ALLOWED:
        return False

    success = deposit_amount(data)
    return success
    #don't forget tot log the txn


def deposit_amount(data):
    start_of_week = get_start_of_week(data["time"])
    #add to log
    #insert into user detail someplace
    return False

def check_for_dup(txn_id):
    return False

def get_customer_data(customer_id):
    return customer_database[customer_id]

def get_num_deposits_for_date(customer_id):
    start_of_week = get_start_of_week(customer_data["time"])
    

def get_num_deposits_for_week(customer_id):
    return 9

def get_current_total_deposits_for_period(customer_id):
    return 9

class CustomerDatabase():

    def __init__(self):
        self.db = defaultdict()
    
    def get_customer_data(self, customer_id):
        return self.db[customer_id]

    def get_customer_data_for_week(self, customer_id, starting_date):
        data =  get_customer_data(self, customer_id)
        return data["weekly"][starting_date]

    def get_customer_data_for_day(self, customer_id, day):
        data =  get_customer_data(self, customer_id)
        return data["daily"][day]

    def deposit(self, customer_id, day, amount)
        return False
    
