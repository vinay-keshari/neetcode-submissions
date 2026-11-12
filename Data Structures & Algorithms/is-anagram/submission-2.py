class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        count = dict()
        for char_s in s:
            if char_s in count.keys():
                count[char_s] += 1
            else:
                count[char_s] = 1

        for char_t in t:
            if char_t in count.keys() and count[char_t] >0:
                count[char_t] -= 1
                if count[char_t] == 0:
                    del count[char_t]
            else:
                count[char_t] = 1
        if len(count):
            return False
        else:
            return True