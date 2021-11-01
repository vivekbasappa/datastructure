

class Node(object):
	def __init__(self, val):
		self.value = val
		self.left = None
		self.right = None

class Tree(object):
	def __init__(self):
		self.root = None
	
	def __insert__(self, node, val):
		if val < node.value:
			if node.left == None:
				node.left = Node(val)
			else: 
				self.__insert__(node.left, val)
		elif val > node.value:
			if node.right == None:
				node.right = Node(val)
			else: 
				self.__insert__(node.right, val)

	def insert(self, val):
		if self.root == None:
			 self.root = Node(val)
		else:
			self.__insert__(self.root, val)

	def display(self):
		node = self.root
		self.__print__(node)
	
	def __print__(self, node):
		if node == None:
			return
		self.__print__(node.left)
		print(" "+str(node.value)+" ")
		self.__print__(node.right)

	def is_balanced(self):
		return self.__is_balanced__(self.root)

	def __is_balanced__(self, root): 
		if root == None:
			return True

		if root.left != None and root.left.value > root.value:
				return False
		if root.right != None and root.right.value < root.value:
				return False
		if not self.__is_balanced__(root.left) and not self.__is_balanced__(root.right):
			return False

		return True

tree = Tree()
tree.insert(30)
tree.insert(20)
tree.insert(50)
tree.insert(10)
tree.insert(15)
tree.insert(60)

tree.display()
print("is it balanced tree:", tree.is_balanced())