# it should parse the file correctly

# it should manage a users account correctly

#unit
from _decimal import Decimal, ROUND_HALF_UP

from db import Validator, CustomerDatabase
from util import breakdown_date, convert_to_utc
from util import sanitize_currency

def test_wednesday_is_monday_offset_2():
    datestring = "2019-07-03T12:23:34Z"

    monday, index = breakdown_date(convert_to_utc(datestring))
    assert monday == "2019-07-01"
    assert index == "2"

def test_sanitize_currency():

    money = '$1,150,593.21'
    curr = sanitize_currency(money)
    assert curr == round(Decimal(1150593.21),2)

def test_sanitize_currency_from_garbage():
    money = '$zx,cmzxc'
    curr = sanitize_currency(money)
    assert curr == 0


monday = '2019-07-01'

def test_should_not_store_duplicates():
    db = CustomerDatabase()
    results, code = db.deposit("111", "222", "2019-07-01T23:23:12Z", "$123.45")
    assert code == 200
    results, code = db.deposit("111","222", "2019-03-01T23:23:12Z", "$143.45")
    assert code == 403


def test_should_be_invalid_too_much_per_day_over_limit_per():
    deposit_records  = {monday:{'1':[]}}
    v = Validator()
    is_valid = v.validate(deposit_records, monday,"1", 5600)
    assert is_valid is False

def test_should_be_invalid_too_much_per_day():
    deposit_records  = {monday:{'1':[1232.45, 2342.56]}}
    v = Validator()
    is_valid = v.validate(deposit_records, monday, "1", 2300)
    assert is_valid is False

def test_should_be_invalid_too_much_per_week():
    deposit_records = {monday: {'1': [4900], '2': [4000], '3': [4900], '4': [4900], '5': [4900]}}
    v = Validator()
    is_valid = v.validate(deposit_records,monday, "1", 2300)
    assert is_valid is False


def test_should_be_invalid_too_many_per_day():
    deposit_records  = {monday: {'1': [122.45, 2342.56, 100.22]}}
    v = Validator()
    is_valid = v.validate(deposit_records, monday, "1", 230.23)
    assert is_valid is False

def test_validation_should_be_valid():
    deposit_records = {monday: {'1': [122.45, 2342.56]}}
    v = Validator()
    is_valid = v.validate(deposit_records,monday, "1", 230.23)
    assert is_valid is True


# def test_user_entry_should_be_allowed():
#     assert False
#
# def test_user_entry_too_much_per_day():
#     assert False
#
# def test_user_entry_too_much_per_week():
#     assert False
#
# def test_user_entry_too_many_per_day():
#     assert False
#
# def test_transaction_based_logger_is_in_sync():
#     assert False

