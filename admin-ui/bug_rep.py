import PySimpleGUI as sg

def enable_enter_clicks(event, window):
	QT_ENTER_KEY1 = 'special 16777220'
	QT_ENTER_KEY2 = 'special 16777221'

	if event in ('\r', QT_ENTER_KEY1, QT_ENTER_KEY2):         # Check for ENTER key
		# go find element with Focus
		elem = window.find_element_with_focus()
		if elem is not None and elem.Type == sg.ELEM_TYPE_BUTTON:       # if it's a button element, click it
			elem.Click()

buttonRow = [sg.Push(),
  sg.Submit('Apply Changes', key='-APPLY_CHANGES-', size=(11, 1)), 
  sg.Submit('Change Pin', key='-CHANGE_PIN-', size=(11, 1)),
  sg.Submit('Delete user', key='-DEL_USR-', size=(11, 1)),
  sg.Push()
]


layout = [buttonRow]
window = sg.Window("bug?", layout, return_keyboard_events=True) #return_keyboard_events=True

while True:
	event, values = window.read()
	enable_enter_clicks(event, window)

	print('event, ' + str(event))
	#print('values:')
	#print(values)

	if event  in ("Exit", sg.WIN_CLOSED, 'none'):
		break


window.close()
