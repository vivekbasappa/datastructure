class ListNode(object):
	def __init__(self, x):
		self.value = x
		self.next = None

class List(object):
	def __init__(self):
		self.head = None

	def add(self, x):
		if self.head is None:
			self.head = ListNode(x)
		else:
			node = self.head
			while node.next is not None:
				node = node.next
			node.next = ListNode(x)

	def print1(self):
		node = self.head
		while node.next is not None:
			print(node.value, " ")

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        sum = ''
        remainder = 0
        current_l1_node = l1.head
        current_l2_node = l2.head
        while current_l1_node is not None and current_l2_node is not None:
            position_sum = current_l1_node.value + current_l2_node.value + remainder
            print("adding %d %d %d" % (current_l1_node.value,  current_l2_node.value, position_sum))
            current_l1_node = current_l1_node.next
            current_l2_node = current_l2_node.next
            coief = 0
            if position_sum  >= 10:
                remainder = 1
                coief = 10 - position_sum 
                print("coief---------------:", coief)
            else:
                coief = position_sum
                remainder = 0
            sum += str(coief)
        sum = sum[::-1]  
        return int(sum)


list1 = List()
list1.add(2)
list1.add(4)
list1.add(3)

list2 = List()
list2.add(5)
list2.add(6)
list2.add(4)

solution = Solution()
print(solution.addTwoNumbers(list1, list2))
