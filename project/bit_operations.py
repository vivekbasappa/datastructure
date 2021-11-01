class BitFields(object):
 	""" A bit field class

			fieldwidth is a tuple or list containing the
			bit width of each field, from least significant
			to most significant.
	"""

	def __init__(self, totalwidth, fieldwidths, value=0):
		if sum(fieldwidths) != totalwidth:
			raise ValueError, "Field width error"

		self.fieldwidths = fieldwidths
		self.num_fields = len(fieldwidths)

		#Calculate field offsets
		self.offsets = []
		pos = 0
		for w in fieldwidths:
			self.offsets.append(pos)
			pos += w

		#Set up bitfield attribute names
		self.field_names = ['b' + str(i) for i in range(self.num_fields)]

		#Initialize
		self.setvalue(value)

 		#Set all fields to zero
		def clear(self):
			for f in self.field_names:
				setattr(self, f, 0)

		#A generator expression of all the field values
		def _all_fields(self):
			return (getattr(self, f) for f in self.field_names)

		def __str__(self):
		return ', '.join(['%s: 0x%x' % (f, v)
                    for f, v in zip(self.field_names, self._all_fields())])

		#Get the register value as an int
		@property
		def value(self):
			return sum(v << p for v, p in zip(self._all_fields(), self.offsets))

		#Set field values
		def regset(self, **kwargs):
			for f, v in kwargs.items():
				setattr(self, f, v)

		#Set the register from an int
		def setvalue(self, value):
			for f, w in zip(self.field_names, self.fieldwidths):
				#print f, w
				mask = (1 << w) - 1
				v = value & mask
				value >>= w
				setattr(self, f, v)
