import math


class MinStack(object):

	def __init__(self):
			self.stack = []
			self.minimum = None

	def push(self, val):
			"""
			:type val: int
			:rtype: None
			"""
			minimum = val
			if self.stack and val > self.stack[-1][1]: 
					minimum = self.stack[-1][1] 
			self.stack.append((val, minimum))

	def pop(self):
			"""
			:rtype: None
			"""
			val = self.stack.pop()
			return val[0] if val else None

	def top(self):
			"""
			:rtype: int
			"""
			return self.stack[-1][0] if len(self.stack) > 0 else None

	def getMin(self):
			"""
			:rtype: int
			"""
			return self.stack[-1][1] if len(self.stack) > 0 else None
        

# Your MinStack object will be instantiated and called as such:
val = 0
obj = MinStack()
obj.push(val)
print(obj.top())
obj.push(1)
print(obj.top())
obj.push(0)
print("stack:%s, getMin:%d" %(str(obj.stack), obj.getMin()))
print("stacl:%s, pop:%d" % (str(obj.stack), obj.pop()))
print("stacl:%s, top:%d" % (str(obj.stack), obj.top()))
print("stacl:%s, getMin:%d" % (str(obj.stack), obj.getMin()))
