'''

data = datetime(2022, 3, 19, 22, 14, 57)
data_for_fmt = '2022-03-19 22:14:57'
data_fmt = '%Y-%m-%d %H:%M:%S'

data_fmt = datetime.strptime(data_for_fmt, data_fmt)

print(data_fmt)
'''

from  datetime import datetime, timedelta
from pytz import timezone


#data = datetime.now()
#print(data.timestamp())
#data = datetime(2022, 3, 22, 15, 16, 54, tzinfo=timezone("Asia"))

data_inicio = datetime(2022, 3, 19, 22, 14, 57)
