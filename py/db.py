import sys
from collections import OrderedDict
import json
from singleton_decorator import singleton

from util import *
import logging

from validation import Validator

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

customer = dict()
WEEKLY_THRESHOLD = 20000
DAILY_THRESHOLD = 5000
TOTAL_LOADS_DAILY = 3



@singleton
class CustomerDatabase():

    def __init__(self):
        logger.info('inititlizing db')
        self.db = dict()
        self.validator = Validator()

    def __init_customer_data(self, customer_id):
        if not customer_id in self.db:

            logger.debug("creating new customer")
            default_customer = {"logs": [], "deposit_records": {}}
            self.db[customer_id] = default_customer

        return self.db[customer_id]

    def __does_transaction_exist(self, logs, txn_id):
        in_logs = txn_id in logs
        return in_logs

    def __update_transaction_history(self, customer_id, txn_id):
        logs = self.db[customer_id]["logs"]

        logs.append(txn_id)
        self.db[customer_id]["logs"] = logs

    def __update_db(self, customer_id, deposits):
        try:
            self.db[customer_id]["deposit_records"] = deposits

        except Exception as e:
            logger.exception(e, exc_info=True)

    def deposit(self,  customer_id,txn_id, time,  amount):
        try:

            amount = sanitize_currency(amount)
            customer_data = self.__init_customer_data(customer_id)

            if self.__does_transaction_exist(customer_data["logs"], txn_id):
                return False, 403

            deposits = customer_data["deposit_records"]
            monday, date_index = breakdown_date(convert_to_utc(time))
            valid = self.validator.validate(deposits, monday, date_index, amount)

            self.__update_transaction_history(customer_id, txn_id)

            if valid:
                if monday not in deposits:
                    deposits[monday] = {date_index: []}
                if date_index not in deposits[monday]:
                    deposits[monday][date_index] = []
                deposits[monday][date_index].append(amount)

                self.__update_db(customer_id, deposits)

                return True, 200
            else:
                return False, 406
        except Exception as e:
            logging.exception("Could not process data", exc_info=True)
            sys.exit(0)
    
