# Build Tools

## requirements.in

Adding Pip dependencies is done via the 'requirements.in' text file.
This file MUST be compiled with pip-compile to create the 'requirements.txt'
pip lock file. This file is then read by the pip_parse command in the
WORKSPACE file.
