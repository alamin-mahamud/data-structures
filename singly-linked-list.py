class Node(object):

    def __init__(self, data=None, next_node=None):
        self.data = data
        self.next_node = next_node

    def get_data(self):
        return self.data

    def get_next(self):
        return self.next_node

    def set_next(self, new_next):
        self.next_node = new_next

    def __str__(self):
        return self.data


class LinkedList(object):
    def __init__(self, head=None):
        self.head = head

    def traverse_to_end(self, prnt=False):
        current = self.head
        while current:
            if prnt:
                print(current, end=' ')  # print current node if you want
            temp = current.get_next()
            if temp is None:
                return current
            current = temp

    def add_back(self, data):
        new_node = Node(data)
        last_node = self.traverse_to_end()
        if not last_node:
            self.head = new_node
            return
        last_node.set_next(new_node)

    def add_front(self, data):
        new_node = Node(data)
        new_node.set_next(self.head)
        self.head = new_node

    def size(self):
        current = self.head
        count = 0
        while current:
            count += 1
            current = current.get_next()
        return count

    def search(self, data):
        current = self.head
        found = False
        while current and found is False:
            if current.get_data() == data:
                found = True
            else:
                current = current.get_next()
        if current is None:
            raise ValueError("Data not in list")
        return current

    def remove_from_back(self):
        current = self.head

        if not current:
            return

        if current.get_next() is None:
            self.head = None
            return

        while current:
            one = current.get_next()
            if one:
                if one.get_next():
                    current = one
                else:
                    current.set_next(None)
                    return

    def remove_from_front(self):
        if self.head:
            next_node = self.head.get_next()
            if next_node:
                self.head = next_node
            else:
                self.head = None

    def remove_with_data(self, data):
        current = self.head
        previous = None
        found = False
        while current and found is False:
            if current.get_data() == data:
                found = True
            else:
                previous = current
                current = current.get_next()
        if current is None:
            raise ValueError("Data not in list")
        if previous is None:
            self.head = current.get_next()
        else:
            previous.set_next(current.get_next())

    def reverse_in_pairs(self):
        previous = self.head
        # if no item then no need to reverse
        if not previous:
            return

        current = previous.get_next()
        # if single item no need to reverse
        if not current:
            return

        self.head = current

        back = None

        while current:
            # print('p-', previous, 'c-', current, 'b-', back, end=' ')

            one = current
            two = one.get_next()

            one.set_next(previous)
            previous.set_next(two)

            if back:
                back.set_next(one)
            back = previous

            if not two:
                break

            third = two.get_next()

            previous, current = two, third


def run():
    # initializing a SLL
    my_list = LinkedList()

    # adding nodes in back
    my_list.add_back("1")
    my_list.add_back("2")
    my_list.add_back("3")
    my_list.add_back("4")
    my_list.add_back("5")
    my_list.add_back("6")
    my_list.add_back("7")
    # Now the list should like -> 1, 2, 3, 4, 5, 6, 7
    my_list.traverse_to_end(prnt=True)
    print("----------------------")

    # adding 0 in front of the list
    my_list.add_front("0")
    # Now the list should like -> 0, 1, 2, 3, 4, 5, 6, 7
    my_list.traverse_to_end(prnt=True)
    print("----------------------")

    # removing the first node
    my_list.remove_from_front()
    # Now the list should like -> 1, 2, 3, 4, 5, 6, 7
    my_list.traverse_to_end(prnt=True)
    print("----------------------")

    # removing the last node
    my_list.remove_from_back()
    # Now the list should like -> 1, 2, 3, 4, 5, 6
    my_list.traverse_to_end(prnt=True)
    print("----------------------")

    my_list.add_back("7")
    # Now the list should like -> 1, 2, 3, 4, 5, 6, 7
    my_list.traverse_to_end(prnt=True)
    print("----------------------")

    # reversing in pair
    my_list.reverse_in_pairs()
    # Now the list should like -> 2, 1, 4, 3, 6, 5, 7
    my_list.traverse_to_end(prnt=True)
    print("----------------------")


if __name__ == '__main__':
    run()
