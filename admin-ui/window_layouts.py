import PySimpleGUI as sg

"""
Window for the user tab, includes a list of people and an info panel with details
Expects to be passed a list of users and will fill list based on __str__() method of objects in the list
"""
def usr_window(userList):
	#!todo add refresh btn to reload

	leftUsrCol = sg.Frame('Stormer people', font='Any 16', layout=[
	  [sg.Listbox(values=userList, select_mode=sg.SELECT_MODE_SINGLE, enable_events=True, size=(40, 20), key='-USERLIST-', expand_y=True)],
	  [sg.Text('Search:'), sg.Input(size=(25, 1), focus=True, enable_events=True, key='-FILTER-'),
		sg.Button('Add', size=5, enable_events=True, key='-ADD_USER-')]
	], element_justification='l', expand_y=True)

	#explicitly no middle part for name here for all non dutchies, if data base has names in first, middle, last merge middle and last.
	personalInfo = sg.Col([
	  [sg.Text('Name:', size=8), sg.Input(default_text = '', size=(25, 1), key='-FIRSTNAME-')],
	  [sg.Text('Last Name:', size=8), sg.Input(default_text = '', size=(25, 1), key='-LASTNAME-')],
	  [sg.Text('vunetID:', size=8), sg.Input(default_text = '', size=(25, 1), key='-VUNETID-')],
	])

	updateBtn = sg.Col([[sg.Submit('Apply Changes', key='-APPLY_CHANGES-')]])

	adminInfo = sg.Col([
	  [sg.Text('Balance:', size=8), sg.Text('', size=(25, 1), key='-BALANCE-')],
	])

	transactionPanel = sg.Frame('Transaction',
	[
	  [sg.Text('Amount:', size=8), sg.Input(size=(10, 1), key='-BALANCE_OPERAND-')],
	  sg.vtop([sg.Text('Comment:'), sg.Multiline(size=(35, 5), key='-TRANSACTION_COMMENT-')]),
	  [sg.Submit('Commit Transaction', key='-TRANSACTION-')]
	])

	rightUsrCol = sg.Frame('Personal Info', font='Any 16',
	  layout=[[personalInfo], [updateBtn], [adminInfo], [transactionPanel]],
	  key='-USR_INFO_PANEL-', visible=False, element_justification='l', expand_x=True, expand_y=True)

	userLayout = [
			  sg.vtop([leftUsrCol, rightUsrCol], expand_y=True)
			  ]

	return userLayout

"""
Specefic window used for adding a new user.
"""
def add_usr_window():
	personalInfo = sg.Col([
	  [sg.Text('Name:', size=8), sg.Input(default_text = '', size=(25, 1), key='-FIRSTNAME-')],
	  [sg.Text('Last Name:', size=8), sg.Input(default_text = '', size=(25, 1), key='-LASTNAME-')],
	  [sg.Text('vunetID:', size=8), sg.Input(default_text = '', size=(25, 1), key='-VUNETID-')],
	])

	adminInfo = sg.Col([
	  [sg.Text('Balance:', size=8), sg.Text('0', size=(25, 1), key='-BALANCE-')],
	])

	transactionPanel = sg.Frame('Transaction (Optional)',
	[
	  [sg.Text('Amount:', size=8), sg.Input(size=(10, 1), key='-BALANCE_OPERAND-')],
	  sg.vtop([sg.Text('Comment:'), sg.Multiline(size=(35, 5), key='-TRANSACTION_COMMENT-')]),
	])

	addBtn = sg.Col([[sg.Submit('Add user', key='-ADD_USER-')]])

	rightUsrCol = sg.Frame('', layout=[[personalInfo], [adminInfo], [transactionPanel], [addBtn]])

	return [[rightUsrCol]]
