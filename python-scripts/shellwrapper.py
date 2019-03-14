import subprocess


def run_in_bash(cmd):
    bash = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
    stdout, stderr = bash.communicate()
    returnCode = bash.returncode
    if returnCode == 0:
        return str(stdout)
    else:
        return str(stderr)
