'''
This class might seem kinda useless on it owns but has some rationale behind it.
From the simplegui lib we have "listbox" (and I believe "Combo" as alt. for this) that give the functionality
 needed to easily display a list. From listbox doc: 
:param values:                     list of values to display. Can be any type including mixed types as long as they have __str__ method
:type values:                      List[Any] or Tuple[Any]

it offers one other usefull field namely "metadata of type 'any'". 

So the problem is that I want to display the names of stormers but identify them by vunetID when looking them up in the
 local cached data or pulling from db (cuz duplicates and such, duh!). Since listbox uses py's build in __str__ having
  a list of tuples or something similiar is not functional since it will display the ID aswell. 
I considered adding a sublist to the metadata param and retrieve the id based on 'index' of name in first list. This 
would have been possible but I didn't like it, it also means that in case of duplicate names I have to extract the correct
index from name list to get id, I believe this is possible with python but I was not sure.

So my other idea was to just simply overwrite the __str__ method of my list of tuple objects, I figured that it probably would
be possible but it seemed dirty. Hence this class just so I can have my own nice little str method.

Might start using it to just store all data though in future, this would mean though that all cached user data ends up in the listbox...

'''

class Stormer():
	def __init__(self, firstName, lastName, vunetId, nickName = '', balance = 0):
		self.firstName = firstName
		self.lastName = lastName
		self.vunetId = vunetId
		self.nickName = nickName
		self.balance = balance

	def __str__(self):
		return self.firstName + ' ' + self.lastName


class Product():
	def __init__(self, name, id, price, hidden):
		self.name = name
		self.id = id
		self.price = price
		self.hidden = hidden
		#maybe added a comment string?

	def __str__(self):
		return self.name
