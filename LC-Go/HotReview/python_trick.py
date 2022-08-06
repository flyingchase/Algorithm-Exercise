

class PythonTrick:
    @classmethod
    def merge_dict(cls,d1,d2):
        print(d1.items()|d2.items())
        print('\n')
        d1.update(d2)
        print(f'd1={d1}, d2={d2}\n')

    @classmethod
    def remote_duplicate_key_in_list(cls,lists):
        print(f'input lists={lists}')
        lists = list(set(lists))
        lists.pop(len(lists)-1)
        lists=lists[::-1]
        print(lists)
        return lists

def tmp(n):
    return isinstance(n,int)
    
def tmp_func(n):
    return lambda a:a*n


class Person:
    def __init__(self,name,age):
        self.name=name
        self.age=age

    def print_name(self):
        print(self.age)

class Student(Person):
    def __init__(self, name, age,score):
        super().__init__(name, age)
        self.score=score

    def print_score(self):
        print(self.score)

if __name__ == "__main__":

    my_str="1234567890"
    my_str_ite=iter(my_str)
    for next(my_str_ite):
        print(next(my_str_ite))




