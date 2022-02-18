import PySimpleGUI as sg
import json
import pathlib
from stormer import Stormer

def load_dummy_dat():
    filePath = pathlib.Path(pathlib.Path.cwd(), 'admin-ui/names.json')
    with open(filePath, 'r') as dummyboidatafile:
        dummyData = json.load(dummyboidatafile)

    print('dummy data contains ' + str(len(dummyData)) + ' entries.')
    return dummyData

def get_usr_data(inData): #works on temp dummy data, not sure what api calls I get yet!
    userList = []

    for user in inData: #!todo, if time for fun try a list comprehension instead!
        userList.append(Stormer(user['firstname'], user['lastname'], user['_id'], user['balance']))

    userList.sort(key=lambda x: x.__str__()) # performance note (in python, hilarious!) refactor to use operator instead of lambda might* be better 

    return userList

def usr_window(userList):
    #!todo add refresh btn to reload

    leftUsrCol = sg.Column([
      [sg.Listbox(values=userList, select_mode=sg.SELECT_MODE_SINGLE, enable_events=True, size=(50, 20), key='-USERLIST-')],
      [sg.Text('Search:'), sg.Input(size=(25, 1), focus=True, enable_events=True, key='-FILTER-')]
      ], element_justification='l', expand_x=True, expand_y=True)

    #explicitly no middle part for name here for all non dutchies, if data base has names in first, middle, last merge middle and last.
    personalInfo = sg.Col([
      [sg.Text('Personal info')],
      [sg.Text('Name:'), sg.Input(default_text = '', size=(25, 1), key='-FIRSTNAME-')],
      [sg.Text('Last name:'), sg.Input(default_text = '', size=(25, 1), key='-LASTNAME-')],
      [sg.Text('vunetID:'), sg.Input(default_text = '', size=(25, 1), key='-VUNETID-')],
    ])

    adminInfo = sg.Col([
      [sg.Text('Administrative')],
      [sg.Text('Balance:'), sg.Text('', size=(25, 1), key='-BALANCE-', background_color='white')],
    ])

    updateBtn = sg.Col([[sg.Submit('Apply Changes', key='-APPLY_CHANGES-')]])

    transactionPanel = sg.Col([
      [sg.Text('Transaction')],
      [sg.Text('Change balance by '), sg.Input(size=(8, 1), key='-BALANCE_OPERAND-')],
      [sg.Text('Comment:'), sg.Multiline(size=(25, 1), key='-TRANSACTION_COMMENT-')],
      [sg.Submit('Commit Transaction', key='-TRANSACTION-')]
    ])

    rightUsrCol = sg.Column([[personalInfo], [updateBtn], [adminInfo], [transactionPanel]], key='-USR_INFO_PANEL-', visible=False)

    userLayout = [[sg.Text('Stormer people', font='Any 16')],
              sg.vtop([leftUsrCol, rightUsrCol])
              ]

    return userLayout

def main_window(names):
    sg.theme('SystemDefaultForReal')
    userLayout = usr_window(names)
    productLayout = [[sg.Text('products')]]

    tabs = [[
        sg.TabGroup([[
            sg.Tab('Stormers', userLayout, key='-USERTAB-'), sg.Tab('Products', productLayout, key='-PRODTAB-')
            ]])
    ]]

    #Define Window
    return sg.Window("Tabs", tabs, resizable=True)

def find_user(id, data): #temp function for dummy data, will have to be refactored based on how data is pulled on init and refreshed/loaded on user selection from remote db api
    for user in data:
        if user['_id'] == id:
            return user

    return None

def event_loop(window, data): #!todo functionize these events instead of bunch of stray code, for now hardcoded until using actual api calls
    completeNameList = get_usr_data(data) #todo, the ui lib adds a threading abstraction for slow/long operations (window.perform_long_operation) which could potentionally be used for async refresh, need to look better into async with the ui in general for this purpose.

    while True:
        event, values = window.read() #todo think about using elif for events, can there be more than one event per loop? lets assume no

        print('And now on this was the event, ' + str(event) + ' !')
        #print('values:')
        #print(values)

        if event == sg.WIN_CLOSED or event == 'none':
            break

        if event == '-USERLIST-':
            Stormer = values['-USERLIST-'][0]

            userInfo = find_user(Stormer.vunetId, data)

            window['-FIRSTNAME-'].update(userInfo['firstname'])
            window['-LASTNAME-'].update(userInfo['lastname'])
            window['-VUNETID-'].update(userInfo['_id'])
            window['-BALANCE-'].update(userInfo['balance'])
            window['-USR_INFO_PANEL-'].update(visible=True)

        if event == '-APPLY_CHANGES-':
            data = {}
            data['firstname'] = window['-FIRSTNAME-'].get()
            data['lastname'] = window['-LASTNAME-'].get()
            data['vunetid'] = window['-VUNETID-'].get()
            #data['balance'] = window['-BALANCE-'].get() #if keep, do int check, prolly should not be editable directly though

            print('updating: ' + json.dumps(data))

        if event == '-TRANSACTION-':
            None #todo start using api calls first

        if event == '-FILTER-':
            new_list = [i for i in completeNameList if (values['-FILTER-'].lower() in i.__str__().lower() or values['-FILTER-'].lower() in i.vunetId.lower())]
            window['-USERLIST-'].update(new_list)
        
        completeNameList = get_usr_data(data) #todo async or something prolly

    window.close()

def main():
    data = load_dummy_dat() #will have to figure out a good way of retrieving this, perhaps pull list of names and only get aditional info on selection of name
    usrData = get_usr_data(data)
    window = main_window(usrData)
    event_loop(window, data)
    print('bye')

if __name__ == "__main__": #python!!!!
    main()
