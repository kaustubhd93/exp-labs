class Solution(object):
    def within_range(self, num):
        if num > 2147483647 or num < -2147483648:
            return False
        else:
            return True

    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        store = []
        store_rev = []
        sample = str(x)
        if len(sample) < 2:
            return x

        num_is_neg = False

        for ele in sample:
            store.append(ele)

        for ele in reversed(range(len(store))):
            store_rev.append(store[ele])

        if "-" in store_rev:
            minus_index = store_rev.index("-")
            store_rev.pop(minus_index)
            num_is_neg = True
        if sample.endswith("0"):
            zero_index = store_rev.index("0")
            store_rev.pop(zero_index)

        res = "".join(store_rev)
        if num_is_neg:
            res = "-" + res
        if not self.within_range(int(res)):
            return 0
        return int(res)
