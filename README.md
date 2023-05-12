# KNT but in django

Recommend checking out the [django tutorial](https://docs.djangoproject.com/en/4.2/intro/).
The tutorial has cool commands on how to actually work with django, but a short recap here anyway.

## Setup
To create the initial DB just run `python manage.py migrate` in the root directory.
To create a superuser run `python manage.py createsuperuser`, and follow the instructions.
After that you can start the server with `python manage.py runserver`.
Head over to the admin panel at `localhost:8000/admin` and login with the superuser credentials you have entered.
There you will find a handy panel to manually manage the database.
Create users, products, view transactions etc all nicely in one place.
You can also enter a nice handy shell with `python manage.py shell`, where you get an interactive python console with all the models loaded.


## Development
If you find yourself in the need of editing the model schema, you can do so in `knt_backend/models.py`.
After that you need to run `python manage.py makemigrations` to create the migration files.
Then run `python manage.py migrate` to actually apply the changes to the database.

Useful routes are located in `knt_backend/urls.py`. They actually just point to some functions over in `knt_backend/views.py`.
If you wanna add a route, create a function there, and don't forget to register it in `knt_backend/urls.py`.
