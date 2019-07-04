
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


customer = dict()
WEEKLY_THRESHOLD = 20000
DAILY_THRESHOLD = 5000
TOTAL_LOADS_DAILY = 3


class Validator:

    def __exceeds_total_amount_for_date(self, deposits, monday, date_index, amount):

        if not date_index in deposits[monday]:
            logger.debug("date index not found in array")
            return False
        per_date_txns = deposits[monday][date_index]
        per_date = sum(per_date_txns)

        return per_date + amount > DAILY_THRESHOLD

    def __exceeds_total_amount_for_week(self, deposits, monday, amount):
        weekly_total = 0
        per_week_txns = deposits[monday]

        for d in per_week_txns:
            weekly_total = weekly_total + sum(per_week_txns[d])
        return weekly_total + amount > WEEKLY_THRESHOLD

    def __exceeds_num_deposits_for_day(self, deposits, monday, date_index):
        if  date_index not in deposits[monday]:
            return False

        per_date_txns = deposits[monday][date_index]

        return len(per_date_txns) >= TOTAL_LOADS_DAILY

    def validate(self, deposits, monday, date_index, amount):
        # no entries for this week yet
        if (DAILY_THRESHOLD - amount < 0):
            return False

        if monday not in deposits:
            return True

        if self.__exceeds_total_amount_for_date(deposits, monday, date_index, amount):
            logger.debug("Total with " + str(amount) + " exceeds date")
            return False

        if self.__exceeds_total_amount_for_week(deposits, monday, amount):
            logger.debug("Total with " + str(amount) + " exceeds week")
            return False

        if self.__exceeds_num_deposits_for_day(deposits, monday, date_index):
            logger.debug("Total with exceeds count")
            return False

        #valid transaction
        return True
