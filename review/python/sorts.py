import random


class Solution:
    def randomPartiiton(self, arr: [int], low: int, high: int):
        i = random.randint(low, high)
        arr[i],arr[high]=arr[high],arr[i]
        return self.paratition(arr,low,high)

    def paratition(self,arr:[int],low:int,high:int):
        i=low-1
        pivot = arr[high]

        for j in range(low,high):
            if arr[j]<=pivot
