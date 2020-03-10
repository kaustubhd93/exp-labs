# Directory structure of admin control/project/site
```
└── sampleapp               --just a container.
    ├── db.sqlite3
    ├── manage.py           --command line utility to interact with Django project.
    └── sampleapp
        ├── __init__.py
        ├── settings.py     --configuration of django project.
        ├── urls.py         --URL declarations for django project.
        └── wsgi.py         
```

# "include" in urls.py of project settings

The include() function allows referencing other URLconfs. Whenever Django encounters include(), it chops off whatever part of the URL matched up to that point and then sends the remaining string to the included URLconf for further processing.

# How to add more APIs ?

Create a view in views.py, reference and map the view in urls.py of the app. The main URLconf of the app has already been included in the project urls.py so no need to change anything there.

# How to map views in urls.py ?

Add new urls in the below format in the list urlpatterns.
path("route", views.func_name, name="func_name", kwargs)
kwargs and name is optional.

# Installed APPS defaults in main setting.py

- django.contrib.admin             --Admin site.  
- django.contrib.auth              --Authentication system
- django.contrib.contenttypes      --A framework for content types
- django.contrib.sessions          --A session framework        
- django.contrib.messages          --A messagin framework
- django.contrib.staticfiles       --A framework for managing static files.

# models.py

- It is the blueprint of the tables in your database.

# migrations in django

- Migrations are how Django stores changes to your models(and thus your database schema)-they're just files on disk.
- The "migrate" command takes all migrations that haven't been applied(Django tracks every migration by using a special table in the database called django_migrations) and runs them against the database essentially, synchronizing the changes you made to your models with the schema in the database.
- Migrations can let us change models over time, as we develop our project in real time without the need to delete database or tables and make new ones -- database can be upgraded live without any loss of data.

# 3 steps to modify models

- Change code in models.py
- Run `python manage.py makemigrations <app_name>` (Creates migrations)
- Run `python manage.py migrate` (Applies migration to database)

# How do views and urls work ?

- When one requests http://mysite.com/4 (4 here being the question id) Django loads the mysite.urls module because it is pointed by the ROOT_URLCONF settings. It finds the variable named urlpatterns and traverses the patterns in order. After finding relevant match, it strips off the matching part (in this case http://mysite.com) and sends the remaining text "/4" to "polls.urls" URLconf for further processing. There it matches <int:question_id>/ and calls detail() view.
- The question_id=4 part comes from <int:question_id>. Using angle brackets "captures" part of the URL and sends it as a keyword argument to the view function. The question_id> part of the string defines the name that will be used to identify the matched pattern, and the <int: part is a converter that determines what patterns should match this part of the URL path. If we do not use the int part any random keyword will be matched like  "http://mysite.com/fdsfkds".
