import datetime

def parse_time_zone(date_of_action):
    return "UTC"

def convert_to_utc(date_of_action):
    return date_of_action

def parse_amount(amount):
    return int(amount)

def get_start_of_week(time):
    return datetime.datetime.now()