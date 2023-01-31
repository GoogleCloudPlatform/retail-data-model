
## Python Environment

### Macintosh OS X

Install Brew

Update Directory Permissions
```shell
# You should change the ownership of these directories to your user.
$ chown -R $(whoami) /usr/local/bin /usr/local/etc /usr/local/sbin /usr/local/share /usr/local/share/doc

# And make sure that your user has write permission.
$ chmod u+w /usr/local/bin /usr/local/etc /usr/local/sbin /usr/local/share /usr/local/share/doc

# Install Python 3.9
$ brew install cython@3.9
```

