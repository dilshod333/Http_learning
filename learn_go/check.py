from quran import Quran

item_to_check = 'alcohol'
quran_result = Quran.check(item_to_check)

print(f"Is '{item_to_check}' mentioned in the Quran? {quran_result}")
