def exchange_money(budget, exchange_rate):
    return budget / exchange_rate

def get_change(budget, exchanging_value):
    return budget - exchanging_value

def get_number_of_bills(budget, denomination):
    return budget // denomination

def get_value_of_bills(denomination, number_of_bills):
    return denomination * number_of_bills

def exchangeable_value(budget, exchange_rate, spread, denomination):
    actual_exchange_rate = exchange_rate * (1 + (spread / 100))
    exchanging_value = exchange_money(budget, actual_exchange_rate)
    number_of_bills = get_number_of_bills(exchanging_value, denomination)
    return int(get_value_of_bills(denomination, number_of_bills))
    # return int(((budget / (exchange_rate + (exchange_rate * spread / 100))) // denomination) * denomination)

def non_exchangeable_value(budget, exchange_rate, spread, denomination):
    actual_exchange_rate = exchange_rate * (1 + (spread / 100))
    exchanging_value = exchange_money(budget, actual_exchange_rate)
    return int(exchanging_value % denomination)
