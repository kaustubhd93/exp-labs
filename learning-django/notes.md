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
