# How to set up

- Switch to root user : `sudo su -`
- Create a python virtual environment : `python3 -m venv ip_tracker_ssh`
- Activate the environment : `source ip_tracker_ssh/bin/activate`
- Install Flask web framework : `pip install flask`
- Start application : `nohup python trackclientips.py &`

# File location

├── captureIP.sh             --> Place this is in /etc/profile.d  
└── trackclientips.py        --> Ensure you run this as root in background  

