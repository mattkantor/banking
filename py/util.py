import datetime
from decimal import  ROUND_HALF_UP, InvalidOperation, Decimal
import decimal

# def parse_time_zone(date_of_action):
#     return "UTC"

def convert_to_utc(datetime_of_action):
    return datetime.datetime.strptime(datetime_of_action, "%Y-%m-%dT%H:%M:%SZ") #unclear on the timezone struct for this

def breakdown_date(today):
    monday =  today + datetime.timedelta(days=-today.weekday(), weeks=0)
    date_index = today.weekday()

    return datetime.datetime.strftime(monday, "%Y-%m-%d"), str(date_index)


def sanitize_currency(money):

    from re import sub
    try:
        return Decimal(Decimal(sub(r'[^\d.]', '', money)).quantize(Decimal('.01'), rounding=ROUND_HALF_UP))
    except InvalidOperation as e:
        return 0
