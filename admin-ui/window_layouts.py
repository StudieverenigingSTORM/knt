import PySimpleGUI as sg

mline_right_click_menu = ['', ['Copy', 'Paste', 'Select All', 'Cut']]
filter_tooltip = "Enter a (partial) name or vunetID to filter the list!"

"""Main window that ties it all together"""
def main_window_layout():
	sg.theme('SystemDefaultForReal')
	userLayout = usr_window_layout()
	productLayout = product_window_layout()

	mainLayout = [[
		sg.TabGroup([
				[sg.Tab('Stormers', userLayout, key='-USERTAB-'), sg.Tab('Products', productLayout, key='-PRODTAB-')]
			], expand_x=True, expand_y=True, enable_events=True, key='-TAB_SWITCH-')
	]]

	return mainLayout

"""
Window for the user tab, includes a list of people and an info panel with detailsW
"""
def usr_window_layout(): #todo split up in few more functions!
	#!todo add refresh btn to reload, perhaps

	leftUsrCol = sg.Frame('Stormer people', font='Any 16', layout=[
	  [sg.Listbox(values=[], select_mode=sg.SELECT_MODE_SINGLE, enable_events=True,
	    size=(40, 20), key='-USERLIST-', expand_y=True)],
	  [sg.Text('Search:', tooltip=filter_tooltip), sg.Input(size=(25, 1), focus=True, enable_events=True, key='-U_FILTER-', tooltip=filter_tooltip),
		sg.Button('Add', size=5, enable_events=True, key='-ADD_USER-')]
	], element_justification='l', expand_y=True)

	#explicitly no middle part for name here for all non dutchies, if data base has names in first, middle, last merge middle and last.
	personalInfo = sg.Col([
	  [sg.Text('Name:', size=8), sg.Input(default_text = '', size=(25, 1), key='-FIRSTNAME-')],
	  [sg.Text('Last Name:', size=8), sg.Input(default_text = '', size=(25, 1), key='-LASTNAME-')],
	  [sg.Text('vunetID:', size=8), sg.Input(default_text = '', size=(25, 1), key='-VUNETID-')],
	])

	buttonRow = [sg.Push(), sg.Submit('Apply Changes', key='-APPLY_CHANGES-', size=(11, 1)), sg.Submit('Change Pin', key='-CHANGE_PIN-', size=(11, 1)), sg.Submit('Delete user', key='-DEL_USR-', size=(11, 1)), sg.Push()]

	adminInfo = [sg.Text('Balance:', size=8), sg.Text('', size=(25, 1), key='-BALANCE-')]

	transactionPanel = sg.Frame('Transaction',
	[
	  [sg.Text('Amount:', size=8), sg.Input(size=(10, 1), key='-BALANCE_OPERAND-')],
	  sg.vtop([sg.Text('Comment:'), sg.Multiline(size=(35, 5), key='-TRANSACTION_COMMENT-', right_click_menu=mline_right_click_menu)]),
	  [sg.Push(), sg.Submit('Commit Transaction', key='-TRANSACTION-'), sg.Push()]
	])

	rightUsrCol = sg.Frame('Personal Info', font='Any 16',
	  layout=[[personalInfo], buttonRow, adminInfo, [transactionPanel]],
	  key='-USR_INFO_PANEL-', visible=False, element_justification='l', expand_x=True, expand_y=True)

	userLayout = [sg.vtop([leftUsrCol, rightUsrCol], expand_y=True)]

	return userLayout

"""Window for the products tab"""
def product_window_layout():
	leftCol = sg.Frame('Products', font='Any 16', layout=[
	  [sg.Listbox(values=[], select_mode=sg.SELECT_MODE_SINGLE, enable_events=True,
	    size=(40, 20), key='-PRODUCTLIST-', expand_y=True)],
	  [sg.Text('Search:', tooltip=filter_tooltip), sg.Input(size=(25, 1), focus=True, enable_events=True, key='-P_FILTER-', tooltip=filter_tooltip),
		sg.Button('Add', size=5, enable_events=True, key='-ADD_PRODUCT-')]
	], element_justification='l', expand_y=True)
	return [[leftCol]]

"""
Specefic window used for adding a new user.
"""
def add_usr_window_layout(): #!todo add cancel button!
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
	  sg.vtop([sg.Text('Comment:'), sg.Multiline(size=(35, 5), key='-TRANSACTION_COMMENT-', right_click_menu=mline_right_click_menu)]),
	])

	buttonRow = [sg.Push(), sg.Submit('Add user', key='-ADD_USER-'), sg.Submit('Cancel', key='-CANCEL-')]

	rightUsrCol = sg.Frame('', layout=[[personalInfo], [adminInfo], [transactionPanel], buttonRow])

	return [[rightUsrCol]]


"""Window for changing pin"""
def change_pin_window_layout():
	element = sg.Frame("", [
		[sg.Text('New Pin:', size=8), sg.Input(default_text = '', size=(25, 1), key='-PIN-', enable_events=True, password_char='*')],
		[sg.Checkbox('Show pin', enable_events=True, key='-SHOW_PIN-'), sg.Push(), sg.Submit("Apply", key='-APPLY_PIN-'), sg.Submit('Cancel', key='-CANCEL-')]
	])

	return [[element]]


"""Window for confirmation of deleting user"""
def delete_usr_window_layout():
	layout = [
		[sg.Text('', key='-TEXT-')],
		[sg.Push(), sg.Submit('Noo!!!', key='-NO_DEL-'), sg.Submit("Confirm", key='-CONFIRM_DELETE-'), sg.Push()]
	]
	return layout