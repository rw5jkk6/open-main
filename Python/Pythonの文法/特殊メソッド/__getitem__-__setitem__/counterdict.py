from collections import defaultdict

class CounterDict:
    def __init__(self):
        self._data = {}
        self._get_count = defaultdict(int)
        self._set_count = defaultdict(int)

    def __getitem__(self, key):
        self._get_count[key] += 1
        return self._data[key]

    def __setitem__(self, key, value):
        self._set_count[key] += 1
        self._data[key] = value
    
    @property
    def count(self):
        return {
            'set': list(self._set_count.items()),
            'get': list(self._get_count.items()),
        }

c = CounterDict()
c['x'] = 1  # 1
print(c['x'])

c['x'] = 2
c['y'] = 3
print(c.count)
# {'set': [('x', 2), ('y', 1)], 'get': [('x', 1)]}
